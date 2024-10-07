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

func (self *ClientRx) maskingKeyBoard2(c *netSender.KeyboardData2) int {
	num := 0
	if self.keys.Ctrl != c.Ctrl {
		self.ctrl_define(c.Ctrl)
		self.keys.Ctrl = c.Ctrl
		num += 1
	}
	for i, button := range c.Button {
		if self.keys.Button[i] != self.banKey(button) {
			if button > self.keys.Button[i] {
				self.OriginalButton.Store(button, true)
			} else if self.keys.Button[i] > button {
				self.OriginalButton.Delete(self.keys.Button[i])
			}
			self.keys.Button[i] = self.banKey(button)
			num += 1
		}
	}
	return num
}

func (self *ClientRx) ctrl_define(ctrl byte) byte {
	if ctrl&hid.LeftCtrl != 0 {
		self.OriginCtrl.Store(byte(hid.LeftCtrl), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.LeftCtrl))
	}

	if ctrl&hid.RightCtrl != 0 {
		self.OriginCtrl.Store(byte(hid.RightCtrl), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.RightCtrl))
	}

	if ctrl&hid.LeftShift != 0 {
		self.OriginCtrl.Store(byte(hid.LeftShift), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.LeftShift))
	}

	if ctrl&hid.RightShift != 0 {
		self.OriginCtrl.Store(byte(hid.RightShift), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.RightShift))
	}

	if ctrl&hid.LeftAlt != 0 {
		self.OriginCtrl.Store(byte(hid.LeftAlt), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.LeftAlt))
	}

	if ctrl&hid.RightAlt != 0 {
		self.OriginCtrl.Store(byte(hid.RightAlt), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.RightAlt))
	}

	if ctrl&hid.LeftWindows != 0 {
		self.OriginCtrl.Store(byte(hid.LeftWindows), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.LeftWindows))
	}

	if ctrl&hid.RightWindows != 0 {
		self.OriginCtrl.Store(byte(hid.RightWindows), true)
	} else {
		self.OriginCtrl.Delete(byte(hid.RightWindows))
	}

	if ctrl == hid.CmdNone {
		self.OriginCtrl.Clear()
	}

	return 0
}

func (self *ClientRx) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}

	return key
}
