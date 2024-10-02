package action

import "main.go/netSender"

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
