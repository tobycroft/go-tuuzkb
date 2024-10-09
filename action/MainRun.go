package action

import (
	"main.go/netReceiver"
	"main.go/netSender"
	"sync"
	"sync/atomic"
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

	self.ClientTx.CmdSetUsbString(netSender.StrTypeManufacturer, "2.4G ManualFacture")
	self.ClientTx.CmdSetUsbString(netSender.StrTypeProduct, "2.4G Reciever")
	self.ClientTx.CmdSetUsbString(netSender.StrTypeSerial, "05ac")
	go self.kb_banSomeKeys()

	go self.mouse_runnable()
	self.keyboard_runnable()
	panic("runnable")
}
