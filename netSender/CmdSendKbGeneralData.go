package netSender

import "main.go/define/cmd"

// 发送 USB 键盘普通数据
func (kb *ClientTx) CmdSendKbGeneralData() *ClientTx {
	kb.head()
	kb.sendStruct.Cmd = cmd.CMD_SEND_KB_GENERAL_DATA
	kb.data([]byte{}).send()
	return kb
}
