package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ClientTx struct {
	sendBuf  bytes.Buffer
	sendData sendData

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
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
}

func (self *ClientTx) head() *ClientTx {
	self.sendData = sendData{}
	self.sendData.Head = uint16(start1)<<8 | uint16(start2)
	return self
}

func (self *ClientTx) data(data any) *ClientTx {
	bb := bytes.Buffer{}
	err := binary.Write(&bb, binary.BigEndian, data)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	self.sendData.Len = self.sendData.Len + byte(bb.Len())
	err = binary.Write(&self.sendBuf, binary.BigEndian, self.sendData)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	//fmt.Println(bb.Len())
	//fmt.Println(self.ClientTx.Len)
	self.sendBuf.Write(bb.Bytes())
	return self
}

func (self *ClientTx) sum() *ClientTx {
	sum := byte(0x00)
	for _, b := range self.sendBuf.Bytes() {
		sum = sum + (b)
	}
	err := binary.Write(&self.sendBuf, binary.BigEndian, sum&0xff)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	return self
}

func (self *ClientTx) send() {
	self.sum()
	self.TxChannel <- self.sendBuf.Bytes()
	self.sendBuf.Reset()
}
