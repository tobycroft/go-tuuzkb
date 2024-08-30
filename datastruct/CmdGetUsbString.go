package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

type Usbstr struct {
	HidStingType byte
	HidLen       byte
}

// 获取字符串描述符配置
// 通过该命令向芯片获取当前所使
// 用的 USB 字符串描述符配置
func (kb *Kb) CmdGetUsbString() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
	kb.data([]byte{0x1}).send()
	return kb
}

// 获取字符串描述符配置
func (rx *ClientRx) CmdGetUsbStringRecv(buf []byte) string {
	bs := bytes.NewReader(buf)
	crx := ClientRx{}
	binary.Read(bs, binary.BigEndian, &crx)
	us := Usbstr{}
	binary.Read(bs, binary.BigEndian, &us)
	bt := make([]byte, us.HidLen)
	binary.Read(bs, binary.BigEndian, &bt)
	fmt.Println(crx)
	fmt.Println(us)
	fmt.Println(string(bt))
	return string(bt)
}
