package netReceiver

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *ClientRx) RouterKeyboard() {
	for report := range self.keyboardMain {
		//fmt.Println(report)
		if self.maskingKeyBoard2(&report) > 0 {
			self.KeyboardRxChannel <- report
			self.CurrentPressed.Range(func(key, value interface{}) bool {
				fmt.Println("originctrl", key, value)
				return true
			})
		}
	}
}

func (self *ClientRx) maskingKeyBoard2(c *netSender.KeyboardData) int {
	num := 0
	if self.keys.Ctrl != c.Ctrl {
		self.ctrl_define(int16(c.Ctrl) - int16(self.keys.Ctrl))
		self.keys.Ctrl = c.Ctrl
		num += 1
	}
	if self.keys.Button0 != self.banKey(c.Button0) {
		if c.Button0 > self.keys.Button0 {
			self.CurrentPressed.Store(c.Button0, true)
		} else if self.keys.Button0 > c.Button0 {
			self.CurrentPressed.Delete(self.keys.Button0)
		}
		self.keys.Button0 = self.banKey(c.Button0)
		num += 1
	}
	if self.keys.Button1 != self.banKey(c.Button1) {
		if c.Button1 > self.keys.Button1 {
			self.CurrentPressed.Store(c.Button1, true)
		} else if self.keys.Button1 > c.Button1 {
			self.CurrentPressed.Delete(self.keys.Button1)
		}
		self.keys.Button1 = self.banKey(c.Button1)
		num += 1
	}
	if self.keys.Button2 != self.banKey(c.Button2) {
		if c.Button2 > self.keys.Button2 {
			self.CurrentPressed.Store(c.Button2, true)
		} else if self.keys.Button2 > c.Button2 {
			self.CurrentPressed.Delete(self.keys.Button2)
		}
		self.keys.Button2 = self.banKey(c.Button2)
		num += 1
	}
	if self.keys.Button3 != self.banKey(c.Button3) {
		if c.Button3 > self.keys.Button3 {
			self.CurrentPressed.Store(c.Button3, true)
		} else if self.keys.Button3 > c.Button3 {
			self.CurrentPressed.Delete(self.keys.Button3)
		}
		self.keys.Button3 = self.banKey(c.Button3)
		num += 1
	}
	if self.keys.Button4 != self.banKey(c.Button4) {
		if c.Button4 > self.keys.Button4 {
			self.CurrentPressed.Store(c.Button4, true)
		} else if self.keys.Button4 > c.Button4 {
			self.CurrentPressed.Delete(self.keys.Button4)
		}
		self.keys.Button4 = self.banKey(c.Button4)
		num += 1
	}
	if self.keys.Button5 != self.banKey(c.Button5) {
		if c.Button5 > self.keys.Button5 {
			self.CurrentPressed.Store(c.Button5, true)
		} else if self.keys.Button5 > c.Button5 {
			self.CurrentPressed.Delete(self.keys.Button5)
		}
		self.keys.Button5 = self.banKey(c.Button5)
		num += 1
	}
	return num
}

func (self *ClientRx) ctrl_define(ctrl int16) byte {
	switch ctrl {
	case hid.LeftCtrl:
		self.OriginCtrl.Store(hid.LeftCtrl, true)

	case -hid.LeftCtrl:
		self.OriginCtrl.Delete(hid.LeftCtrl)

	case hid.RightCtrl:
		self.OriginCtrl.Store(hid.RightCtrl, true)

	case -hid.RightCtrl:
		self.OriginCtrl.Delete(hid.RightCtrl)

	case hid.LeftShift:
		self.OriginCtrl.Store(hid.LeftShift, true)

	case -hid.LeftShift:
		self.OriginCtrl.Delete(hid.LeftShift)

	case hid.RightShift:
		self.OriginCtrl.Store(hid.RightShift, true)

	case -hid.RightShift:
		self.OriginCtrl.Delete(hid.RightShift)

	case hid.LeftAlt:
		self.OriginCtrl.Store(hid.LeftAlt, true)

	case -hid.LeftAlt:
		self.OriginCtrl.Delete(hid.LeftAlt)

	case hid.RightAlt:
		self.OriginCtrl.Store(hid.RightAlt, true)

	case -hid.RightAlt:
		self.OriginCtrl.Delete(hid.RightAlt)

	case hid.LeftWindows:
		self.OriginCtrl.Store(hid.LeftWindows, true)

	case -hid.LeftWindows:
		self.OriginCtrl.Delete(hid.LeftWindows)

	case hid.RightWindows:
		self.OriginCtrl.Store(hid.RightWindows, true)

	case -hid.RightWindows:
		self.OriginCtrl.Delete(hid.RightWindows)

	}

	return 0
}

func (self *ClientRx) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}

	return key
}
