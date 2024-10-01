package main

import "main.go/netTcp"

func main() {
	//10.0.0.90
	//rx := netReceiver.ClientRx{}
	//rx.IP = "10.0.0.91"
	//rx.Port = 6666
	//rx.Run()
	//var run action.Runnable
	//go run.MainRun(&rx, &tx)
	serverudp := netTcp.ServerUDP{}
	serverudp.Rx()
}
