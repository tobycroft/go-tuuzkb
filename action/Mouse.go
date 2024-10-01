package action

import "main.go/common"

type mouse struct {
}

func (self *Action) mouse_runnable() {
	for c := range self.ClientRx.MouseRxChannel {
		go common.PrintRedis("匹配鼠標", c)
	}
	panic("鼠标通道意外结束")
}
