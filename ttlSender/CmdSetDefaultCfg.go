package ttlSender

import (
	"main.go/define/cmd"
)

func (kb *Kb) CmdSetDefaultCfg() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_SET_DEFAULT_CFG
	kb.data([]byte{}).send()
	return kb
}
