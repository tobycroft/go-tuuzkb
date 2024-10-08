package netSender

import "main.go/define/cmd"

func (self *ClientTx) CmdReadMyHidData() *ClientTx {
	self.head(cmd.CMD_READ_MY_HID_DATA).data([]byte{}).send()
	return self
}
