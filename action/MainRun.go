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
	ClientRx *netReceiver.ClientRx
	ClientTx *netSender.ClientTx

	AutoPressed  sync.Map
	lastPressSum atomic.Value
	c            netSender.KeyboardData2
}

var Mask = mask{}

type mask struct {
	Button sync.Map
	Ctrl   sync.Map
}

func (self *Action) MainRun(clientrx *netReceiver.ClientRx, clienttx *netSender.ClientTx) {
	self.ClientRx = clientrx
	self.ClientTx = clienttx
	go func() {
		time.Sleep(3 * time.Second)
		SetUsbString()
	}()

	go self.kb_banSomeKeys()
	go self.mouse_runnable()
	self.keyboard_runnable()
	panic("runnable")
}

func SetUsbString() {
	netSender.Ctx.CmdSetUsbString(netSender.StrTypeManufacturer, "2.4G MonkaKeyboard")
	netSender.Ctx.CmdSetUsbString(netSender.StrTypeProduct, "2.4G MonkaReciever")
	netSender.Ctx.CmdSetUsbString(netSender.StrTypeSerial, "A87")
}
