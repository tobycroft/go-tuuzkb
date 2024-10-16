package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync/atomic"
)

var Ctx = &ClientTx{}

type ClientTx struct {
	TxChannel      chan []byte
	MouseTxChannel chan any
}

func (self *ClientTx) Ready() {
	self.MouseTxChannel = make(chan any)
	self.TxChannel = make(chan []byte)
}

const start1 = 0x57
const start2 = 0xab

// 发送数据联合体
type sendData struct {
	Head [2]byte // 帧头 (2个字节)
	Addr byte    // 地址码 (1个字节)
	Cmd  byte    // 命令码 (1个字节)
	Len  byte    // 后续数据长度 (1个字节)
}

type SendTx struct {
	sendBuf  *bytes.Buffer
	sendData *sendData
	sumhex   *atomic.Uint32
}

func (self *SendTx) Head(Cmd byte) *SendTx {
	return &SendTx{
		sendData: &sendData{
			Head: [2]byte{
				start1,
				start2,
			},
			Addr: 0x00,
			Cmd:  Cmd,
			Len:  0x00,
		},
		sendBuf: &bytes.Buffer{},
		sumhex:  &atomic.Uint32{},
	}
}

func (self *SendTx) Data(data any) *SendTx {
	bb := &bytes.Buffer{}
	err := binary.Write(bb, binary.BigEndian, data)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	self.sendData.Len = self.sendData.Len + byte(bb.Len())
	err = binary.Write(self.sendBuf, binary.BigEndian, self.sendData)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	self.sendBuf.Write(bb.Bytes())
	return self
}

func (self *SendTx) sum() *SendTx {
	self.sumhex.Store(0x00)
	for _, b := range self.sendBuf.Bytes() {
		self.sumhex.Add(uint32(b))
	}
	err := binary.Write(self.sendBuf, binary.BigEndian, byte(self.sumhex.Load()&0xff))
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	//fmt.Println("sum", self.sendBuf.Bytes(), byte(self.sum.Load()&0xff))
	return self
}

func (self *SendTx) Send() {
	self.sum()
	Ctx.TxChannel <- self.sendBuf.Bytes()
}

// 定义帧结构
type SendFrame struct {
	Header      [2]byte // 帧头，固定2个字节
	AddressCode byte    // 地址码，固定1个字节
	CommandCode byte    // 命令码，固定1个字节
	DataLength  byte    // 后续数据长度，固定1个字节
	DataSection []byte  // 后续数据，变长
	Checksum    byte    // 校验和，固定1个字节
}

func (self *SendFrame) Head(Cmd byte) *SendFrame {
	return &SendFrame{
		Header:      [2]byte{0x57, 0xAB},
		AddressCode: 0x00,
		CommandCode: Cmd,
	}
}

func (self *SendFrame) Data(data any) *SendFrame {
	bb := &bytes.Buffer{}
	err := binary.Write(bb, binary.BigEndian, data)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	self.DataSection = bb.Bytes()
	self.DataLength = byte(bb.Len())
	return self
}

// 计算校验和
func (f *SendFrame) sum() {
	sum := f.Header[0] + f.Header[1] + f.AddressCode + f.CommandCode + f.DataLength
	for _, b := range f.DataSection {
		sum += b
	}
	f.Checksum = sum
}

// 编码为[]byte
func (f *SendFrame) encode() ([]byte, error) {
	f.sum()
	// 创建一个缓冲区，大小为固定部分 + 数据长度 + 校验和
	buf := make([]byte, 5+len(f.DataSection)+1) // 4 = 2 (header) + 1 (address) + 1 (command)

	// 写入固定部分
	copy(buf[0:2], f.Header[:])
	buf[2] = f.AddressCode
	buf[3] = f.CommandCode
	buf[4] = f.DataLength

	// 写入可变长度数据
	copy(buf[5:], f.DataSection)

	// 写入校验和
	buf[5+len(f.DataSection)] = f.Checksum
	//fmt.Println(buf)

	return buf, nil
}

func (self *SendFrame) Send() {
	buf, err := self.encode()
	if err != nil {
		panic(err)
	}
	Ctx.TxChannel <- buf
}
