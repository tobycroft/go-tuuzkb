package netSender

import (
	"bytes"
	"encoding/binary"
	"main.go/define/cmd"
)

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendMsRelData(mousebyte MouseData) *ClientTx {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, mousebyte)
	//buf.WriteString(str)
	//fmt.Println(string(buf.Bytes()))
	//self.data(buf.Bytes()).Send()
	//self.Data([]byte{}).Send()
	new(SendTx).Head(cmd.CMD_SEND_MS_REL_DATA).Data(mousebyte).Send()

	return self
}

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendMsRelWheel(wheel int8) *ClientTx {
	buf := bytes.Buffer{}
	mousebyte := MouseData{
		Resv:  0x01,
		Wheel: byte(-wheel),
	}
	binary.Write(&buf, binary.BigEndian, mousebyte)
	//fmt.Println(mousebyte)
	new(SendTx).Head(cmd.CMD_SEND_MS_REL_DATA).Data(mousebyte).Send()

	return self
}
