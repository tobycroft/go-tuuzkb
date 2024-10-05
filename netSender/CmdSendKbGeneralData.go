package netSender

import (
	"bytes"
	"encoding/binary"
	"main.go/define/cmd"
)

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendKbGeneralData(keybytes KeyboardData) *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_SEND_KB_GENERAL_DATA

	buf := bytes.Buffer{}
	keybytes.Resv = 0x00
	binary.Write(&buf, binary.BigEndian, keybytes)
	//buf.WriteString(str)
	//fmt.Println(string(buf.Bytes()))
	//self.data(buf.Bytes()).send()
	//self.data([]byte{}).send()
	self.data(keybytes).send()

	return self
}

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendKbGeneralDataRaw(keybytes KeyboardData2) *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_SEND_KB_GENERAL_DATA

	buf := bytes.Buffer{}
	keybytes.Resv = 0x00
	binary.Write(&buf, binary.BigEndian, keybytes)
	self.data(keybytes).send()

	return self
}
