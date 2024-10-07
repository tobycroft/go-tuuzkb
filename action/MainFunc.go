package action

import (
	"fmt"
	"main.go/netSender"
)

func (self *Action) KeyDown(key byte) {
	out := netSender.KeyboardData2{}
	out.Ctrl, out.Button, out.Resv = self.kb_washing()
	for i, button := range out.Button {
		if button == 0 {
			self.AutoPressed.Store(key, int64(i))
			out.Button[i] = key
			break
		}
	}
	self.ClientTx.CmdSendKbGeneralDataRaw(out)
	fmt.Println("keyboardAutoDN", out)
}

func (self *Action) KeyUp(key byte) {
	self.AutoPressed.Delete(key)
	out := netSender.KeyboardData2{}
	out.Ctrl, out.Button, out.Resv = self.kb_washing()
	self.ClientTx.CmdSendKbGeneralDataRaw(out)
	fmt.Println("keyboardAutoUP", out)
}

func (self *Action) SendKbGeneralDataRaw() {
	out := netSender.KeyboardData2{}
	out.Ctrl, out.Button, out.Resv = self.kb_washing()
	go fmt.Println("keybaordsnd", out)
	if out.Resv != self.lastPressSum.Load() {
		self.lastPressSum.Store(out.Resv)
		out.Resv = 0x00
		self.ClientTx.CmdSendKbGeneralDataRaw(out)
	}
}

func (self *Action) checkKeyIsPressedAny(Ctrl byte, Btn ...byte) bool {
	btns := self.c.Button
	if self.c.Ctrl == Ctrl {
		for _, btn := range Btn {
			for _, b := range btns {
				if b == btn {
					return true
				}
			}
		}
	}
	return false
}

func (self *Action) checkKeyIsPressed(Ctrl byte, Btn ...byte) bool {
	num := 0
	btns := self.c.Button
	if self.c.Ctrl == Ctrl {
		for _, btn := range Btn {
			for _, b := range btns {
				if b == btn {
					num += 1
				}
			}
		}
	}
	if num == len(Btn) {
		return true
	} else {
		return false
	}
}

func (self *Action) checkKeyIsPressedByOrder(Ctrl byte, Btn ...byte) bool {
	num := 0
	btns := self.c.Button
	if self.c.Ctrl == Ctrl {
		for i, btn := range Btn {
			if btns[i] == btn {
				num += 1
			}
		}
	}
	if num == len(Btn) {
		return true
	} else {
		return false
	}
}

func (self *Action) kb_washing() (Ctrl byte, Button [6]byte, sum byte) {
	for i, button := range self.c.Button {
		_, ok := Mask.Button.Load(button)
		if !ok {
			Button[i] = button
		} else {
			Button[i] = 0
		}
		sum += button
	}

	self.ClientRx.OriginCtrl.Range(func(key, value interface{}) bool {
		_, ok := Mask.Ctrl.Load(key.(byte))
		if !ok {
			Ctrl += key.(byte)
		}
		return true
	})
	sum += Ctrl
	return
}

func (self *Action) kb_add_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		Mask.Ctrl.Store(key, true)
	} else {
		Mask.Button.Store(key, true)
	}
}

func (self *Action) kb_remove_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		Mask.Ctrl.Delete(key)
	} else {
		Mask.Button.Delete(key)
	}
}

func (self *Action) kb_chec_mask(key byte, is_ctrl bool) bool {
	if is_ctrl {
		_, ok := Mask.Ctrl.Load(key)
		return ok
	} else {
		_, ok := Mask.Button.Load(key)
		return ok
	}
}
