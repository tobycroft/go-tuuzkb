package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdGetInfo() *ClientTx {
	SApi.Head(cmd.CMD_GET_INFO).Data([]byte{}).Send()
	return self
}
