package datastruct

import (
	"main.go/define/cmd"
)

func (kb *Kb) CmdGetInfo() *Kb {
	kb.calcHead()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
	kb.calcData([]byte{}).calcSum()
	return kb
}
