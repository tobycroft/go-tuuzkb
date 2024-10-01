package netReceiver

func (self *ClientRx) RouterKeyboard() {
	for report := range self.keyboardMain {
		self.KeyboardReport <- report
	}
}
