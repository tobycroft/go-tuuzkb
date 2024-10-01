package main

import (
	"main.go/action"
	"main.go/netReceiver"
	"main.go/netSender"
	"main.go/netTcp"
	"net"
)

func main() {

	ClientTx := &netSender.ClientTx{}
	ClientRx := &netReceiver.ClientRx{}
	ClientRx.Ready()
	ClientTx.Ready()

	action := &action.Action{}
	go action.MainRun(ClientRx, ClientTx)
	serverudp := netTcp.ServerUDP{
		SendServer: &net.UDPAddr{
			IP:   net.ParseIP("10.0.0.90"),
			Port: 6666,
		},
		ClientTx: ClientTx,
		ClientRx: ClientRx,
	}
	serverudp.Start()

}
