package action

import (
	"main.go/netReceiver"
	"main.go/netSender"
)

func (self *Action) mouse_runnable() {
	for c := range netReceiver.Crx.MouseRxChannel {
		//fmt.Println(*c)
		netSender.Ctx.CmdSendMsRelData(*c)
		//go common.PrintRedis("匹配鼠標", c)
	}
	panic("鼠标通道意外结束")
}
