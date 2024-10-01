package netSender

import "main.go/define/cmd"

func (self *ClientTx) CmdReadMyHidData() *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_READ_MY_HID_DATA
	self.data([]byte{}).send()
	return self
}
