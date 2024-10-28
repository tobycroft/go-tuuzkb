package netReceiver

func (self *ClientRx) RouterMouse() {
	for c := range self.mouseMain {
		self.MouseRxChannel <- c
	}
}
