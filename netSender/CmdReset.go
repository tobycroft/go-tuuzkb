package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdReset() *ClientTx {
	self.head(cmd.CMD_RESET).data([]byte{}).send()
	return self
}
