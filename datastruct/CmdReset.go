package datastruct

import (
	"main.go/define/cmd"
)

func (kb *Kb) CmdReset() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_RESET
	kb.data([]byte{}).send()
	return kb
}
