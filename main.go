package main

import (
	"main.go/netTcp"
)

func main() {
	//10.0.0.90
	go netTcp.ClientTx()
	netTcp.ClientRx()
}
