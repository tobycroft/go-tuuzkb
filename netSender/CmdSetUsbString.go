package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

// 设置字符串描述符配置
func (self *ClientTx) CmdSetUsbString(HidStingType byte, str string) *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_SET_USB_STRING
	//self.sendData.Len = 2
	usbstr := Usbstr{
		HidStringType: HidStingType,
		HidLen:        0,
	}
	usbstr.HidLen = byte(len(str))
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, usbstr)
	buf.WriteString(str)
	fmt.Println(string(buf.Bytes()))
	self.data(buf.Bytes()).send()
	return self
}
