package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *Action) keyboard_runnable() {
	for c := range self.ClientRx.KeyboardRxChannel {
		self.MaskingKeyBoard(&c)
		self.ClientTx.CmdSendKbGeneralData(c)
		fmt.Println("keybaordrecv", c.Ctrl, c.Button)
	}
	panic("键盘通道意外结束")
}

func (self *Action) MaskingKeyBoard(c *netSender.KeyboardData) {
	RearrangedButton := [6]byte{}
	for i, button := range c.Button {
		switch button {
		case hid.CmdErrorRollOver:
			break

		default:
			RearrangedButton[i] = button
			break
		}
	}
	c.Button = RearrangedButton
}
