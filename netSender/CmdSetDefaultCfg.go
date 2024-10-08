package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdSetDefaultCfg() *ClientTx {
	self.head(cmd.CMD_SET_DEFAULT_CFG).data([]byte{}).send()
	return self
}
