package netTcp

import "main.go/netReceiver"

var DataChannel = make(chan []byte, 1)

func Start() {
	for c := range DataChannel {
		netReceiver.Crx.MessageRouter(c)
	}
}
