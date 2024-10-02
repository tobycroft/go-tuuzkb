package netReceiver

import "main.go/netSender"

func (self *ClientRx) RouterKeyboard() {
	for report := range self.keyboardMain {
		//fmt.Println(report)
		self.KeyboardRxChannel <- report
	}
}

func (self *ClientRx) maskingKeyBoard2(c *netSender.KeyboardData) {
	num := 0
	if self.Ctrl != c.Ctrl {
		self.Ctrl = c.Ctrl
		num += 1
	}
	if self.Button0 != self.banKey(c.Button0) {
		self.Button0 = self.banKey(c.Button0)
		num += 1
	}
	if self.Button1 != self.banKey(c.Button1) {
		self.Button1 = self.banKey(c.Button1)
		num += 1
	}
	if self.Button2 != self.banKey(c.Button2) {
		self.Button2 = self.banKey(c.Button2)
		num += 1
	}
	if self.Button3 != self.banKey(c.Button3) {
		self.Button3 = self.banKey(c.Button3)
		num += 1
	}
	if self.Button4 != self.banKey(c.Button4) {
		self.Button4 = self.banKey(c.Button4)
		num += 1
	}
	if self.Button5 != self.banKey(c.Button5) {
		self.Button5 = self.banKey(c.Button5)
		num += 1
	}

	return num
}
