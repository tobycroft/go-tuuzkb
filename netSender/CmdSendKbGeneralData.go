package netSender

import "main.go/define/cmd"

// 发送 USB 键盘普通数据
func (kb *NetSender) CmdSendKbGeneralData() *NetSender {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_SEND_KB_GENERAL_DATA
	kb.data([]byte{}).send()
	return kb
}
