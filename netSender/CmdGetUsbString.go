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
	self.head()
	self.sendData.Cmd = cmd.CMD_GET_USB_STRING
	self.data([]byte{HidStringType}).send()
	time.Sleep(1 * time.Second)
	return self
}

// 获取字符串描述符配置
func CmdGetUsbStringRecv(buf []byte) string {
	bs := bytes.NewReader(buf)
	us := Usbstr{}
	binary.Read(bs, binary.LittleEndian, &us)
	bt := make([]byte, us.HidLen)
	binary.Read(bs, binary.LittleEndian, &bt)
	return string(bt[2:us.HidLen])
}
