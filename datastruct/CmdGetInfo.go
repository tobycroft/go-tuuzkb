package datastruct

import "main.go/define/cmd"

func (kb *Kb) CmdGetInfo() {
	kb.CalcHead()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
}
