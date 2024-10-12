package netSender

import (
	"main.go/define/cmd"
)

func (self *ClientTx) CmdPing() *ClientTx {
	new(SendTx).Head(cmd.CMD_READ_MY_HID_DATA).Data([]byte{}).Send()
	return self
}
