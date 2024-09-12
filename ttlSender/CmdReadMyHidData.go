package ttlSender

import "main.go/define/cmd"

func (kb *Kb) CmdReadMyHidData() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_READ_MY_HID_DATA
	kb.data([]byte{}).send()
	return kb
}
