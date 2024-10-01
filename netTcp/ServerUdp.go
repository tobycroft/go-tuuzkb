package netTcp

import (
	"bytes"
	"main.go/netReceiver"
	"main.go/netSender"
	"net"
)

type ServerUDP struct {
	SendServer net.Addr

	Conn net.PacketConn

	ClientTx *netSender.ClientTx
	ClientRx *netReceiver.ClientRx
}

func (self *ServerUDP) Rx() *ServerUDP {
	buff := make([]byte, 512)
	var err error
	self.Conn, err = net.ListenPacket("udp", ":6666")
	if err != nil {
		panic(err.Error())
	}

	for {
		_, addr, err := self.Conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}
		slice_byte := bytes.Split(buff, []byte{0x57, 0xab})
		for _, ddd := range slice_byte {
			self.ClientRx.MessageRouter(ddd, addr, self.Conn)
		}
	}
}
