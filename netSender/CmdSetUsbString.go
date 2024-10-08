package netSender

import (
	"bytes"
	"encoding/binary"
	"main.go/define/cmd"
	"time"
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
	//fmt.Println(string(buf.Bytes()))
	self.data(buf.Bytes()).send()
	time.Sleep(500 * time.Millisecond)
	return self
}

// 0x00 表示厂商字符串描述符；0x01 表示产品字符串描述符；
// 0x02 表示序列号字符串描述符
const (
	StrTypeManufacturer = 0x00
	StrTypeProduct      = 0x01
	StrTypeSerial       = 0x02
)
