package netSender

import (
	"main.go/define/cmd"
)

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendKbGeneralData(keybytes KeyboardData) {
	keybytes.Resv = 0x00
	SApi.Head(cmd.CMD_SEND_KB_GENERAL_DATA).Data(keybytes).Send()
}

// 发送 USB 键盘普通数据
func (self *ClientTx) CmdSendKbGeneralDataRaw(keybytes KeyboardData2) {
	keybytes.Resv = 0x00
	SApi.Head(cmd.CMD_SEND_KB_GENERAL_DATA).Data(keybytes).Send()
}
