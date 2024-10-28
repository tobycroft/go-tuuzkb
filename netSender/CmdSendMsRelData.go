package netSender

import (
	"main.go/define/cmd"
)

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendMsRelData(mousebyte MouseData) *ClientTx {
	mousebyte.Resv = 0x01
	//fmt.Println(mousebyte)
	SApi.Head(cmd.CMD_SEND_MS_REL_DATA).Data(mousebyte).Send()
	return self
}

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendMsRelWheel(wheel int8) *ClientTx {
	mousebyte := MouseData{
		Resv:  0x01,
		Wheel: byte(-wheel),
	}
	//fmt.Println(mousebyte)
	SApi.Head(cmd.CMD_SEND_MS_REL_DATA).Data(mousebyte).Send()
	return self
}
