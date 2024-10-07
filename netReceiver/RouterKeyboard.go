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
		self.ctrl_define(int16(c.Ctrl) - int16(self.keys.Ctrl))
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

func (self *ClientRx) ctrl_define(ctrl int16) byte {
	switch ctrl {
	case hid.LeftCtrl:
		self.OriginCtrl.Store(byte(hid.LeftCtrl), true)
		break

	case -hid.LeftCtrl:
		self.OriginCtrl.Delete(byte(hid.LeftCtrl))
		break

	case hid.RightCtrl:
		self.OriginCtrl.Store(byte(hid.RightCtrl), true)
		break

	case -hid.RightCtrl:
		self.OriginCtrl.Delete(byte(hid.RightCtrl))
		break

	case hid.LeftShift:
		self.OriginCtrl.Store(byte(hid.LeftShift), true)
		break

	case -hid.LeftShift:
		self.OriginCtrl.Delete(byte(hid.LeftShift))
		break

	case hid.RightShift:
		self.OriginCtrl.Store(byte(hid.RightShift), true)
		break

	case -hid.RightShift:
		self.OriginCtrl.Delete(byte(hid.RightShift))
		break

	case hid.LeftAlt:
		self.OriginCtrl.Store(byte(hid.LeftAlt), true)
		break

	case -hid.LeftAlt:
		self.OriginCtrl.Delete(byte(hid.LeftAlt))
		break

	case hid.RightAlt:
		self.OriginCtrl.Store(byte(hid.RightAlt), true)
		break

	case -hid.RightAlt:
		self.OriginCtrl.Delete(byte(hid.RightAlt))
		break

	case hid.LeftWindows:
		self.OriginCtrl.Store(byte(hid.LeftWindows), true)
		break

	case -hid.LeftWindows:
		self.OriginCtrl.Delete(byte(hid.LeftWindows))
		break

	case hid.RightWindows:
		self.OriginCtrl.Store(byte(hid.RightWindows), true)
		break

	case -hid.RightWindows:
		self.OriginCtrl.Delete(byte(hid.RightWindows))
		break

	case 0:
		self.OriginCtrl.Clear()
		break

	}

	return 0
}

func (self *ClientRx) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}

	return key
}
