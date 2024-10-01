package netReceiver

func (self *Rx) RouterKeyboard() {
	for report := range self.keyboardMain {
		self.KeyboardReport <- report
	}
}
