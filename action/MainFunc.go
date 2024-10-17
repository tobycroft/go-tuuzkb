package action

import (
	"fmt"
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
	fmt.Println("keybaordsnd", out)

	if out.Resv != lastPressSum.Load() {
		lastPressSum.Store(out.Resv)
		out.Resv = 0x00
		netSender.Ctx.CmdSendKbGeneralDataRaw(out)
	}
}

func (self *Action) checkKeyIsPressedAny(Ctrl byte, Btn ...byte) bool {
	if CurrentPress.Ctrl.Load() == Ctrl || Ctrl == hid.CmdNone {
		for _, btn := range Btn {
			for _, b := range OnchangePress.Button {
				if b.Load() == btn {
					return true
				}
			}
		}
	}
	return false
}

func (self *Action) checkKeyIsPressed(Ctrl byte, Btn ...byte) bool {
	num := 0
	if CurrentPress.Ctrl.Load() == Ctrl {
		for _, btn := range Btn {
			for _, b := range CurrentPress.Button {
				if b.Load() == btn {
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
	if CurrentPress.Ctrl.Load() == Ctrl {
		for i, btn := range Btn {
			if CurrentPress.Button[i].Load() == btn {
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
	for i, button := range CurrentPress.Button {
		_, ok := Mask.Button.Load(button.Load().(byte))
		if !ok {
			Button[i] = button.Load().(byte)
		} else {
			Button[i] = 0
		}
		sum += button.Load().(byte)
	}

	netReceiver.OriginCtrl.Range(func(key, value interface{}) bool {
		_, ok := Mask.Ctrl.Load(key.(byte))
		if !ok {
			Ctrl += key.(byte)
		}
		return true
	})
	sum += Ctrl
	return
}

func kb_add_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		Mask.Ctrl.Store(key, true)
	} else {
		Mask.Button.Store(key, true)
	}
}

func kb_remove_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		Mask.Ctrl.Delete(key)
	} else {
		Mask.Button.Delete(key)
	}
}

func kb_chec_mask(key byte, is_ctrl bool) bool {
	if is_ctrl {
		_, ok := Mask.Ctrl.Load(key)
		return ok
	} else {
		_, ok := Mask.Button.Load(key)
		return ok
	}
}
