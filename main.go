package main

import (
	"main.go/action"
	"main.go/netTcp"
)

func main() {
	//10.0.0.90
	go netTcp.ClientTx()
	netTcp.ClientRx()

	var run action.Runnable
	go run.MainRun(&rx, &tx)
}
