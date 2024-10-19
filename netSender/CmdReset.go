package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdReset() *ClientTx {
	send.SendApi.Head(cmd.CMD_RESET).Data([]byte{}).Send()
	return self
}
