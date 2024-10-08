package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdGetInfo() *ClientTx {
	self.head(cmd.CMD_GET_INFO).data([]byte{}).send()
	return self
}
