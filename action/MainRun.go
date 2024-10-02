package action

import (
	"main.go/netReceiver"
	"main.go/netSender"
	"sync"
)

type Action struct {
	//将你需要缓存的数据存在这里
	ClientRx *netReceiver.ClientRx
	ClientTx *netSender.ClientTx

	MaskKey sync.Map

	button
	key
}

type button struct {
	Ctrl byte

	Button0 byte
	Button1 byte
	Button2 byte
	Button3 byte
	Button4 byte
	Button5 byte
}

func (self *Action) MainRun(clientrx *netReceiver.ClientRx, clienttx *netSender.ClientTx) {
	self.ClientRx = clientrx
	self.ClientTx = clienttx
	go self.mouse_runnable()
	self.keyboard_runnable()
	panic("runnable")
}
