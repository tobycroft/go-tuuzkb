package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdReset() *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_RESET
	self.data([]byte{}).send()
	return self
}
