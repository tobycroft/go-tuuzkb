package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdSetDefaultCfg() *ClientTx {
	new(SendTx).Head(cmd.CMD_SET_DEFAULT_CFG).Data([]byte{}).Send()
	return self
}
