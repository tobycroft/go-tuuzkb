package action

import (
	"main.go/netSender"
)

func (self *Action) kb_gen_output(c netSender.KeyboardData) (out netSender.KeyboardData2) {
	out.Ctrl, out.Button, out.Resv = self.kb_washing(c)
	return out
}

func (self *Action) checkKeyIsPressed(c netSender.KeyboardData, Ctrl byte, Btn ...byte) bool {
	num := 0
	btns := [6]byte{c.Button0, c.Button1, c.Button2, c.Button3, c.Button4, c.Button5}
	if c.Ctrl == Ctrl {
		for i, btn := range Btn {
			if btns[i] == btn {
				num += 1
			}
		}
		//for _, btn := range Btn {
		//	switch btn {
		//	case c.Button0, c.Button1, c.Button2, c.Button3, c.Button4, c.Button5:
		//		num += 1
		//		break
		//
		//	default:
		//		break
		//	}
		//}

	}

	if num == len(Btn) {
		return true
	} else {
		return false
	}
}
func (self *Action) kb_washing(c netSender.KeyboardData) (Ctrl byte, Button [6]byte, sum byte) {
	_, ok := self.Mask.Button.Load(c.Button0)
	if !ok {
		Button[0] = c.Button0
	}
	sum += Button[0]
	_, ok = self.Mask.Button.Load(c.Button1)
	if !ok {
		Button[1] = c.Button1
	}
	sum += Button[1]
	_, ok = self.Mask.Button.Load(c.Button2)
	if !ok {
		Button[2] = c.Button2
	}
	sum += Button[2]
	_, ok = self.Mask.Button.Load(c.Button3)
	if !ok {
		Button[3] = c.Button3
	}
	sum += Button[3]
	_, ok = self.Mask.Button.Load(c.Button4)
	if !ok {
		Button[4] = c.Button4
	}
	sum += Button[4]
	_, ok = self.Mask.Button.Load(c.Button5)
	if !ok {
		Button[5] = c.Button5
	}
	sum += Button[5]
	self.ClientRx.OriginCtrl.Range(func(key, value interface{}) bool {
		_, ok = self.Mask.Ctrl.Load(key.(byte))
		if !ok {
			Ctrl += key.(byte)
		}
		return true
	})
	sum += Ctrl
	return
}
