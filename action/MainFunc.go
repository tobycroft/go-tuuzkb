package action

import (
	"main.go/define/hid"
	"main.go/netReceiver"
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
	netSender.Ctx.CmdSendKbGeneralDataRaw(out)
	//go fmt.Println("keyboardAutoDN", out)
}

func (self *Action) KeyUp(key byte) {
	self.AutoPressed.Delete(key)
	out := netSender.KeyboardData2{}
	out.Ctrl, out.Button, out.Resv = self.kb_washing()
	netSender.Ctx.CmdSendKbGeneralDataRaw(out)
	//go fmt.Println("keyboardAutoUP", out)
}

func (self *Action) SendKbGeneralDataRaw() {
	out := netSender.KeyboardData2{}
	out.Ctrl, out.Button, out.Resv = self.kb_washing()
	if out.Resv != lastPressSum.Load() {
		//go fmt.Println("keybaordsnd", out)
		lastPressSum.Store(out.Resv)
		out.Resv = 0x00
		netSender.Ctx.CmdSendKbGeneralDataRaw(out)
	}
}

func (self *Action) checkKeyIsPressedAny(Ctrl byte, Btn ...byte) bool {
	if byte(CurrentPress.Ctrl.Load()) == Ctrl || Ctrl == hid.CmdNone {
		for _, btn := range Btn {
			for _, b := range OnchangePress.Button {
				if byte(b.Load()) == btn {
					return true
				}
			}
		}
	}
	return false
}

func (self *Action) checkKeyIsPressed(Ctrl byte, Btn ...byte) bool {
	num := 0
	if byte(CurrentPress.Ctrl.Load()) == Ctrl {
		for _, btn := range Btn {
			for _, b := range CurrentPress.Button {
				if byte(b.Load()) == btn {
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
	if byte(CurrentPress.Ctrl.Load()) == Ctrl {
		for i, btn := range Btn {
			if byte(CurrentPress.Button[i].Load()) == btn {
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
	Mask.ButtonMu.RLock()
	for i, button := range CurrentPress.Button {
		_, ok := Mask.Button[byte(button.Load())]
		if !ok {
			Button[i] = byte(button.Load())
		} else {
			Button[i] = 0
		}
		sum += Button[i]
	}
	Mask.ButtonMu.RUnlock()

	Mask.CtrlMu.RLock()
	netReceiver.OriginCtrlMu.RLock()
	for key := range netReceiver.OriginCtrlMap {
		_, ok := Mask.Ctrl[key]
		if !ok {
			Ctrl += key
		}
	}
	netReceiver.OriginCtrlMu.RUnlock()
	Mask.CtrlMu.RUnlock()
	sum += Ctrl
	return
}

func kb_add_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		Mask.CtrlMu.Lock()
		Mask.Ctrl[key] = true
		Mask.CtrlMu.Unlock()
	} else {
		Mask.ButtonMu.Lock()
		Mask.Button[key] = true
		Mask.ButtonMu.Unlock()
	}
}

func kb_remove_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		Mask.CtrlMu.Lock()
		delete(Mask.Ctrl, key)
		Mask.CtrlMu.Unlock()
	} else {
		Mask.ButtonMu.Lock()
		delete(Mask.Button, key)
		Mask.ButtonMu.Unlock()
	}
}

func kb_chec_mask(key byte, is_ctrl bool) bool {
	if is_ctrl {
		Mask.CtrlMu.RLock()
		_, ok := Mask.Ctrl[key]
		Mask.CtrlMu.RUnlock()
		return ok
	} else {
		Mask.ButtonMu.RLock()
		_, ok := Mask.Button[key]
		Mask.ButtonMu.RUnlock()
		return ok
	}
}