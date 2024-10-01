package netTcp

import (
	"bytes"
	"main.go/netReceiver"
	"main.go/netSender"
	"net"
)

type ServerUDP struct {
	//IP   string
	//Port int

	ClientTx *netSender.ClientTx
	ClientRx *netReceiver.ClientRx
}

func (self *ServerUDP) Rx() *ServerUDP {
	buff := make([]byte, 512)
	Conn, err := net.ListenPacket("udp", ":6666")
	if err != nil {
		panic(err.Error())
	}

	for {
		_, addr, err := Conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}
		slice_byte := bytes.Split(buff, []byte{0x57, 0xab})
		for _, ddd := range slice_byte {
			self.ClientRx.MessageRouter(ddd, addr)
		}
	}
}

func (self *ServerUDP) Tx() *ServerUDP {
	//buff := make([]byte, 512)
	return self
}
