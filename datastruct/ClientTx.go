package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

type Kb struct {
	Sendbuf bytes.Buffer
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

func (kb *Kb) CalcHead() {
	kb.Ctx.Head = uint16(start1)<<8 | uint16(start2)
}

func (kb *Kb) CalcData(data []byte) {
	err := binary.Write(&kb.Sendbuf, binary.BigEndian, data)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
}

func (kb *Kb) CalcSum() {
	sum := byte(0x00)
	for _, b := range kb.Sendbuf.Bytes() {
		sum = sum + b
	}
	err := binary.Write(&kb.Sendbuf, binary.BigEndian, sum&0xFF)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
}

func (kb *Kb) CmdGetInfo() {
	kb.CalcHead()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
}

func (kb *Kb) CmdGetUsb() {
	err := binary.Write(&kb.Sendbuf, binary.BigEndian, &kb.Ctx)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
}
