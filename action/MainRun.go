package action

import (
	"main.go/netReceiver"
	"main.go/netSender"
	"sync"
	"sync/atomic"
	"time"
)

type Action struct {
	//将你需要缓存的数据存在这里
	AutoPressed sync.Map
}

var lastPressSum = &atomic.Value{}

var Mask = &mask{}

type mask struct {
	Button sync.Map
	Ctrl   sync.Map
}

func (self *Action) MainRun() {
	go func() {
		time.Sleep(3 * time.Second)
		netReceiver.SetUsbString()
		netSender.Ctx.CmdGetParaCfg()
	}()

	go Kb_banSomeKeys()
	go self.mouse_runnable()
	self.keyboard_runnable()
	panic("runnable")
}
