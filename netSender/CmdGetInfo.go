package netSender

import (
	"main.go/define/cmd"
	"main.go/netSender/send"
)

func (self *ClientTx) CmdGetInfo() *ClientTx {
	send.SendApi.Head(cmd.CMD_GET_INFO).Data([]byte{}).Send()
	return self
}
