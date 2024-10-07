package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendMsRelData(mousebyte MouseData) *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_SEND_MS_REL_DATA

	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, mousebyte)
	//buf.WriteString(str)
	//fmt.Println(string(buf.Bytes()))
	//self.data(buf.Bytes()).send()
	//self.data([]byte{}).send()
	self.data(mousebyte).send()

	return self
}

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendMsRelWheel(wheel int8) *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_SEND_MS_REL_DATA

	buf := bytes.Buffer{}
	mousebyte := MouseData{}
	if wheel >= 0 {
		mousebyte.Wheel = byte(wheel)
	} else {
		mousebyte.Wheel = byte(wheel) + 0x80
	}
	binary.Write(&buf, binary.BigEndian, mousebyte)
	fmt.Println((buf.Bytes()))
	self.data(mousebyte).send()

	return self
}
