package action

import (
	"fmt"
)

type mouse struct {
}

func (self *Action) mouse_runnable() {
	for c := range self.ClientRx.MouseRxChannel {
		fmt.Println(c)
		//go common.PrintRedis("匹配鼠標", c)
	}
	panic("鼠标通道意外结束")
}
