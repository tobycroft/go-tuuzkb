package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync/atomic"
)

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
