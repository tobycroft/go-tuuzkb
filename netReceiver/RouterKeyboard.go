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
	if ctrl.Load() != c.Ctrl {
		CtrlToMap(c.Ctrl)
		ctrl.Store(c.Ctrl)
		num += 1
	}
	for i, button := range c.Button {
		if buttons[i].Load() != self.banKey(button) {
			if button > buttons[i].Load().(byte) {
				OriginalButton.Store(button, true)
			} else if buttons[i].Load().(byte) > button {
				OriginalButton.Delete(buttons[i])
			}
			buttons[i].Store(self.banKey(button))
			num += 1
		}
	}
	return num
}

func CtrlToMap(ctrl byte) byte {
	if ctrl&hid.LeftCtrl != 0 {
		OriginCtrl.Store(byte(hid.LeftCtrl), true)
	} else {
		OriginCtrl.Delete(byte(hid.LeftCtrl))
	}

	if ctrl&hid.RightCtrl != 0 {
		OriginCtrl.Store(byte(hid.RightCtrl), true)
	} else {
		OriginCtrl.Delete(byte(hid.RightCtrl))
	}

	if ctrl&hid.LeftShift != 0 {
		OriginCtrl.Store(byte(hid.LeftShift), true)
	} else {
		OriginCtrl.Delete(byte(hid.LeftShift))
	}

	if ctrl&hid.RightShift != 0 {
		OriginCtrl.Store(byte(hid.RightShift), true)
	} else {
		OriginCtrl.Delete(byte(hid.RightShift))
	}

	if ctrl&hid.LeftAlt != 0 {
		OriginCtrl.Store(byte(hid.LeftAlt), true)
	} else {
		OriginCtrl.Delete(byte(hid.LeftAlt))
	}

	if ctrl&hid.RightAlt != 0 {
		OriginCtrl.Store(byte(hid.RightAlt), true)
	} else {
		OriginCtrl.Delete(byte(hid.RightAlt))
	}

	if ctrl&hid.LeftWindows != 0 {
		OriginCtrl.Store(byte(hid.LeftWindows), true)
	} else {
		OriginCtrl.Delete(byte(hid.LeftWindows))
	}

	if ctrl&hid.RightWindows != 0 {
		OriginCtrl.Store(byte(hid.RightWindows), true)
	} else {
		OriginCtrl.Delete(byte(hid.RightWindows))
	}

	if ctrl == hid.CmdNone {
		OriginCtrl.Clear()
	}

	return 0
}

func (self *ClientRx) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}

	return key
}
