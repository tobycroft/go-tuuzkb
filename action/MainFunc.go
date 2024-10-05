package action

import (
	"fmt"
	"main.go/netSender"
)

func (self *Action) KeyUp(c netSender.KeyboardData) (out netSender.KeyboardData) {
	for i, button := range self.tx.CmdKeyboard.Button {
		_, ok := self.CurrentPressed.Load(key)
		if ok {
			break
		}
		if button == 0 {
			self.CurrentPressed.Store(key, int64(i))
			self.tx.CmdKeyboard.Button[i] = key
			break
		}
	}
}
func (self *Action) KeyDown(key byte) {
	for i, button := range self.tx.CmdKeyboard.Button {
		_, ok := self.CurrentPressed.Load(key)
		if ok {
			break
		}
		if button == 0 {
			self.CurrentPressed.Store(key, int64(i))
			self.tx.CmdKeyboard.Button[i] = key
			break
		}
	}
	return
}

func (self *Action) SendKbGeneralDataRaw(c netSender.KeyboardData) (out netSender.KeyboardData2) {
	out.Ctrl, out.Button, out.Resv = self.kb_washing(c)
	if out.Resv != self.lastPress {
		self.lastPress = out.Resv
		out.Resv = 0x00
		self.ClientTx.CmdSendKbGeneralDataRaw(out)
		fmt.Println("keybaordsnd", out)
	}
	return out
}

func (self *Action) checkKeyIsPressed(c netSender.KeyboardData, Ctrl byte, Btn ...byte) bool {
	num := 0
	btns := [6]byte{c.Button0, c.Button1, c.Button2, c.Button3, c.Button4, c.Button5}
	if c.Ctrl == Ctrl {
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

func (self *Action) checkKeyIsPressedByOrder(c netSender.KeyboardData, Ctrl byte, Btn ...byte) bool {
	num := 0
	btns := [6]byte{c.Button0, c.Button1, c.Button2, c.Button3, c.Button4, c.Button5}
	if c.Ctrl == Ctrl {
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

func (self *Action) kb_washing(c netSender.KeyboardData) (Ctrl byte, Button [6]byte, sum byte) {
	_, ok := Mask.Button.Load(c.Button0)
	if !ok {
		Button[0] = c.Button0
	}
	sum += Button[0]
	_, ok = Mask.Button.Load(c.Button1)
	if !ok {
		Button[1] = c.Button1
	}
	sum += Button[1]
	_, ok = Mask.Button.Load(c.Button2)
	if !ok {
		Button[2] = c.Button2
	}
	sum += Button[2]
	_, ok = Mask.Button.Load(c.Button3)
	if !ok {
		Button[3] = c.Button3
	}
	sum += Button[3]
	_, ok = Mask.Button.Load(c.Button4)
	if !ok {
		Button[4] = c.Button4
	}
	sum += Button[4]
	_, ok = Mask.Button.Load(c.Button5)
	if !ok {
		Button[5] = c.Button5
	}
	sum += Button[5]
	self.ClientRx.OriginCtrl.Range(func(key, value interface{}) bool {
		_, ok = Mask.Ctrl.Load(key.(byte))
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
