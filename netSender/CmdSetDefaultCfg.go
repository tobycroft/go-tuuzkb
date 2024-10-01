package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdSetDefaultCfg() *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_SET_DEFAULT_CFG
	self.data([]byte{}).send()
	return self
}
