package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdSetDefaultCfg() *ClientTx {
	send.SendApi.Head(cmd.CMD_SET_DEFAULT_CFG).Data([]byte{}).Send()
	return self
}
