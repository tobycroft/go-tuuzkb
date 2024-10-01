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
		self.ClientTx.CmdSendKbGeneralData(c)
		if self.MaskingKeyBoard2(&c) > 0 {
			fmt.Println("keybaordrecv", c.Ctrl, c)
		}
	}
	panic("键盘通道意外结束")
}

func (self *Action) masking(key byte) byte {
	switch key {

	case hid.CmdErrorRollOver:
		return 0x00

	default:
		return key
	}
}

func (self *Action) MaskingKeyBoard2(c *netSender.KeyboardData) int {
	num := 0
	if self.Ctrl != c.Ctrl {
		self.Ctrl = c.Ctrl
		num += 1
	}
	if self.Button0 != self.masking(c.Button0) {
		self.Button0 = self.masking(c.Button0)
		num += 1
	}
	if self.Button1 != self.masking(c.Button1) {
		self.Button1 = self.masking(c.Button1)
		num += 1
	}
	if self.Button2 != self.masking(c.Button2) {
		self.Button2 = self.masking(c.Button2)
		num += 1
	}
	if self.Button3 != self.masking(c.Button3) {
		self.Button3 = self.masking(c.Button3)
		num += 1
	}
	if self.Button4 != self.masking(c.Button4) {
		self.Button4 = self.masking(c.Button4)
		num += 1
	}
	if self.Button5 != self.masking(c.Button5) {
		self.Button5 = self.masking(c.Button5)
		num += 1
	}

	return num
	//fmt.Println(self.Button, Btn)
}

//func (self *Action) MaskingKeyBoard(c *netSender.KeyboardData) bool {
//	Btn := []byte{}
//	for _, btn := range c.Button {
//		if self.masking(btn) {
//			Btn = append(Btn, 0)
//		} else {
//			Btn = append(Btn, btn)
//		}
//	}
//	self.Button = Btn
//	return true
//	//fmt.Println(self.Button, Btn)
//}
