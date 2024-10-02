package action

import "main.go/netSender"

func (self *Action) checkKeyIsPressed(c netSender.KeyboardData, Ctrl byte, Btn ...byte) bool {
	switch Btn[0] {
	case c.Button0:
		return c.Ctrl == Ctrl

	case c.Button1:
		return c.Ctrl == Ctrl

	case c.Button2:
		return c.Ctrl == Ctrl

	case c.Button3:
		return c.Ctrl == Ctrl

	case c.Button4:
		return c.Ctrl == Ctrl

	case c.Button5:
		return c.Ctrl == Ctrl

	default:
		return false
	}
}
