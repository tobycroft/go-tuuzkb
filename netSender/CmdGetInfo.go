package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdGetInfo() *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_GET_INFO
	self.data([]byte{}).send()
	return self
}
