package netReceiver

import (
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *ClientRx) RouterKeyboard() {
	for report := range self.keyboardMain {
		//fmt.Println(report)
		if self.maskingKeyBoard2(&report) > 0 {
			self.KeyboardRxChannel <- report
		}
	}
}

func (self *ClientRx) maskingKeyBoard2(c *netSender.KeyboardData) int {
	num := 0
	if self.keys.Ctrl != c.Ctrl {
		self.keys.Ctrl = c.Ctrl
		num += 1
	}
	if self.keys.Button0 != self.banKey(c.Button0) {
		self.keys.Button0 = self.banKey(c.Button0)
		num += 1
	}
	if self.keys.Button1 != self.banKey(c.Button1) {
		self.keys.Button1 = self.banKey(c.Button1)
		num += 1
	}
	if self.keys.Button2 != self.banKey(c.Button2) {
		self.keys.Button2 = self.banKey(c.Button2)
		num += 1
	}
	if self.keys.Button3 != self.banKey(c.Button3) {
		self.keys.Button3 = self.banKey(c.Button3)
		num += 1
	}
	if self.keys.Button4 != self.banKey(c.Button4) {
		self.keys.Button4 = self.banKey(c.Button4)
		num += 1
	}
	if self.keys.Button5 != self.banKey(c.Button5) {
		self.keys.Button5 = self.banKey(c.Button5)
		num += 1
	}

	return num
}
func (self *ClientRx) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}

	return key
}
