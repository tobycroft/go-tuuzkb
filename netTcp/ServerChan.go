package netTcp

import (
	"main.go/netReceiver"
)

var DataChannel = make(chan []byte, 4)

func Start() {
	for c := range DataChannel {
		netReceiver.Crx.MessageRouter(c)
	}
}
