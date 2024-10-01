package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *Action) keyboard_runnable() {
	for c := range self.ClientRx.KeyboardRxChannel {
		//self.ClientTx.CmdSendKbGeneralData(c)
		//fmt.Println("keybaordrecv", c.Ctrl, c)
		if self.MaskingKeyBoard(&c) {
			self.ClientTx.CmdSendKbGeneralData(c)
			fmt.Println("keybaordrecv", c.Ctrl, c)
		}
	}
	panic("键盘通道意外结束")
}

func (self *Action) MaskingKeyBoard(c *netSender.KeyboardData) bool {
	Btn := []byte{}
	for _, btn := range c.Button {
		if self.masking(btn) {
			Btn = append(Btn, 0)
		} else {
			Btn = append(Btn, btn)
		}
	}
	self.Button = Btn
	return true
	//fmt.Println(self.Button, Btn)
}

func (self *Action) masking(key byte) bool {
	switch key {

	case hid.CmdErrorRollOver:
		return true

	default:
		return false
	}
}
