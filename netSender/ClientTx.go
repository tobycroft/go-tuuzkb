package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type NetSender struct {
	sendBuf bytes.Buffer
	Ctx     ClientTx
}

const start1 = 0x57
const start2 = 0xab

// 发送数据联合体
type ClientTx struct {
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
}

func (kb *NetSender) head() *NetSender {
	kb.Ctx.Head = uint16(start1)<<8 | uint16(start2)
	return kb
}

func (kb *NetSender) data(data any) *NetSender {
	bb := bytes.Buffer{}
	err := binary.Write(&bb, binary.BigEndian, data)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	kb.Ctx.Len = kb.Ctx.Len + byte(bb.Len())
	err = binary.Write(&kb.sendBuf, binary.BigEndian, kb.Ctx)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	//fmt.Println(bb.Len())
	//fmt.Println(kb.Ctx.Len)
	kb.sendBuf.Write(bb.Bytes())
	return kb
}

func (kb *NetSender) sum() *NetSender {
	sum := byte(0x00)
	for _, b := range kb.sendBuf.Bytes() {
		sum = sum + (b)
	}
	err := binary.Write(&kb.sendBuf, binary.BigEndian, sum&0xff)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	return kb
}

func (kb *NetSender) send() {
}
