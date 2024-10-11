package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
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
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
}

type snd struct {
	sendBuf  *bytes.Buffer
	ctx      *ClientTx
	sendData *sendData
}

func (self *ClientTx) head(Cmd byte) *snd {
	return &snd{
		sendData: &sendData{
			Head: uint16(start1)<<8 | uint16(start2),
			Addr: 0x00,
			Cmd:  Cmd,
			Len:  0x00,
		},
		ctx:     self,
		sendBuf: &bytes.Buffer{},
	}
}

func (self *snd) data(data any) *snd {
	bb := bytes.Buffer{}
	err := binary.Write(&bb, binary.BigEndian, data)
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

func (self *snd) sum() *snd {
	sum := byte(0x00)
	for _, b := range self.sendBuf.Bytes() {
		sum = sum + (b)
	}
	err := binary.Write(self.sendBuf, binary.BigEndian, sum&0xff)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	return self
}

func (self *snd) send() {
	self.sum()
	self.ctx.TxChannel <- self.sendBuf.Bytes()
	time.Sleep(1 * time.Millisecond)
	self.sendBuf.Reset()
}
