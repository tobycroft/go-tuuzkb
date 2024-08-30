package datastruct

import "main.go/define/cmd"

func (kb *Kb) CmdSendKbGeneralData() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_SEND_KB_GENERAL_DATA
	kb.data([]byte{}).send()
	return kb
}
