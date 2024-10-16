package action

import (
	"fmt"
	"main.go/netReceiver"
)

type mouse struct {
}

func (self *Action) mouse_runnable() {
	for c := range netReceiver.Crx.MouseRxChannel {
		fmt.Println(c)
		//go common.PrintRedis("匹配鼠標", c)
	}
	panic("鼠标通道意外结束")
}
