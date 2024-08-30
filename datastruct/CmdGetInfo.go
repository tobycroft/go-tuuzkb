package datastruct

import (
	"main.go/define/cmd"
)

func (kb *Kb) CmdGetInfo() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
	kb.data([]byte{}).sum()
	return kb
}
