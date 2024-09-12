package ttlSender

import (
	"main.go/define/cmd"
)

func (kb *Kb) CmdGetInfo() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_GET_INFO
	kb.data([]byte{}).send()
	return kb
}
