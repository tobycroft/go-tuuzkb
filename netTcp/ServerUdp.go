package netTcp

import (
	"encoding/hex"
	"fmt"
	"main.go/netReceiver"
	"main.go/netSender"
	"net"
)

type ServerUDP struct {
	ClientTx netSender.ClientTx
	ClientRx netReceiver.ClientRx
}

func (self *ServerUDP) Rx() *ServerUDP {
	buff := make([]byte, 512)
	Conn, err := net.ListenPacket("udp", ":6666")
	if err != nil {
		panic(err.Error())
	}
	for {
		aa, addr, _ := Conn.ReadFrom(buff)
		fmt.Println(addr, aa, hex.EncodeToString(buff))
	}
}
