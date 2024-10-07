package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

type Usbstr struct {
	HidStringType byte
	HidLen        byte
}

// 获取字符串描述符配置
// 通过该命令向芯片获取当前所使
// 用的 USB 字符串描述符配置
func (self *ClientTx) CmdGetUsbString() *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_GET_USB_STRING
	self.data([]byte{0x1}).send()
	return self
}

// 获取字符串描述符配置
func CmdGetUsbStringRecv(buf []byte) string {
	bs := bytes.NewReader(buf)
	crx := sendData{}
	binary.Read(bs, binary.BigEndian, &crx)
	us := Usbstr{}
	binary.Read(bs, binary.BigEndian, &us)
	bt := make([]byte, us.HidLen)
	binary.Read(bs, binary.BigEndian, &bt)
	fmt.Println("")
	fmt.Println(crx)
	fmt.Println(us)
	fmt.Println(string(bt))
	return string(bt)
}
