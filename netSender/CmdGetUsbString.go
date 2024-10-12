package netSender

import (
	"bytes"
	"encoding/binary"
	"main.go/define/cmd"
	"time"
)

type Usbstr struct {
	HidStringType byte
	HidLen        byte
}

// 获取字符串描述符配置
// 通过该命令向芯片获取当前所使
// 用的 USB 字符串描述符配置
func (self *ClientTx) CmdGetUsbString(HidStringType byte) *ClientTx {
	new(SendTx).Head(cmd.CMD_GET_USB_STRING).Data([]byte{HidStringType}).Send()
	time.Sleep(500 * time.Millisecond)
	return self
}

// 获取字符串描述符配置
func CmdGetUsbStringRecv(buf []byte) string {
	bs := bytes.NewReader(buf)
	us := Usbstr{}
	binary.Read(bs, binary.BigEndian, &us)
	bt := make([]byte, us.HidLen)
	binary.Read(bs, binary.BigEndian, &bt)
	return string(bt[2:us.HidLen])
}
