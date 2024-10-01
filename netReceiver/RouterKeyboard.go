package netReceiver

func (self *Reciever) RouterKeyboard() {
	for report := range self.keyboardMain {
		self.KeyboardReport <- report
	}
}
