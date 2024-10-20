package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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
func (self *SendFrame) sum() *SendFrame {
	self.Checksum = self.Header[0] + self.Header[1] + self.AddressCode + self.CommandCode + self.DataLength
	for _, b := range self.DataSection {
		self.Checksum += b
	}
	return self
}

func (self *SendFrame) Send() {
	self.sum()
	// 创建一个缓冲区，大小为固定部分 + 数据长度 + 校验和
	buf := make([]byte, 5+len(self.DataSection)+1) // 4 = 2 (header) + 1 (address) + 1 (command)

	// 写入固定部分
	copy(buf[0:2], self.Header[:])
	buf[2] = self.AddressCode
	buf[3] = self.CommandCode
	buf[4] = self.DataLength

	// 写入可变长度数据
	copy(buf[5:], self.DataSection)

	// 写入校验和
	buf[5+len(self.DataSection)] = self.Checksum
	//fmt.Println(buf)

	Ctx.TxChannel <- buf
}
