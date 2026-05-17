package netReceiver

import (
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *ClientRx) RouterKeyboard() {
	for report := range self.keyboardMain {
		//fmt.Println(report)
		if self.maskingKeyBoard2(report) > 0 {
			self.KeyboardRxChannel <- report
		}
	}
}

func (self *ClientRx) maskingKeyBoard2(c *netSender.KeyboardData2) int {
	num := 0
	if byte(originCtrl.Load()) != c.Ctrl {
		CtrlToMap(c.Ctrl)
		originCtrl.Store(uint32(c.Ctrl))
		num += 1
	}
	for i, button := range c.Button {
		c.Button[i] = self.banKey(button)
		if byte(originButton[i].Load()) != c.Button[i] {
			if c.Button[i] > byte(originButton[i].Load()) {
				OriginalButtonMu.Lock()
				OriginalButtonMap[button] = true
				OriginalButtonMu.Unlock()
			} else if byte(originButton[i].Load()) > c.Button[i] {
				OriginalButtonMu.Lock()
				delete(OriginalButtonMap, byte(originButton[i].Load()))
				OriginalButtonMu.Unlock()
			}
			originButton[i].Store(uint32(c.Button[i]))
			num += 1
		}
	}
	return num
}

func CtrlToMap(ctrl byte) byte {
	OriginCtrlMu.Lock()
	if ctrl&hid.LeftCtrl != 0 {
		OriginCtrlMap[byte(hid.LeftCtrl)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.LeftCtrl))
	}

	if ctrl&hid.RightCtrl != 0 {
		OriginCtrlMap[byte(hid.RightCtrl)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.RightCtrl))
	}

	if ctrl&hid.LeftShift != 0 {
		OriginCtrlMap[byte(hid.LeftShift)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.LeftShift))
	}

	if ctrl&hid.RightShift != 0 {
		OriginCtrlMap[byte(hid.RightShift)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.RightShift))
	}

	if ctrl&hid.LeftAlt != 0 {
		OriginCtrlMap[byte(hid.LeftAlt)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.LeftAlt))
	}

	if ctrl&hid.RightAlt != 0 {
		OriginCtrlMap[byte(hid.RightAlt)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.RightAlt))
	}

	if ctrl&hid.LeftWindows != 0 {
		OriginCtrlMap[byte(hid.LeftWindows)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.LeftWindows))
	}

	if ctrl&hid.RightWindows != 0 {
		OriginCtrlMap[byte(hid.RightWindows)] = true
	} else {
		delete(OriginCtrlMap, byte(hid.RightWindows))
	}

	if ctrl == hid.CmdNone {
		for k := range OriginCtrlMap {
			delete(OriginCtrlMap, k)
		}
	}
	OriginCtrlMu.Unlock()

	return 0
}

func (self *ClientRx) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}

	return key
}