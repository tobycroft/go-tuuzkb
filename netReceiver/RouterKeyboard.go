package netReceiver

func (self *ClientRx) RouterKeyboard() {
	for report := range self.keyboardMain {
		//fmt.Println(report)
		self.KeyboardRxChannel <- report
	}
}
