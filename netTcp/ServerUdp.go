package netTcp

import (
	"bytes"
	"encoding/hex"
	"fmt"
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

func (self *ServerUDP) Start() *ServerUDP {
	buff := make([]byte, 256)
	var err error
	self.Conn, err = net.ListenPacket("udp", ":6666")
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for keyboard := range self.ClientTx.TxChannel {
			fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
			self.Conn.WriteTo(keyboard, self.SendServer)
		}
	}()

	for {
		_, addr, err := self.Conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}
		if addr.String() == "10.0.0.91:6666" {
			slice_byte := bytes.Split(buff, []byte{0x57, 0xab})
			for _, ddd := range slice_byte {
				self.ClientRx.MessageRouter(ddd, addr, self.Conn)
			}
		} else {
			fmt.Println(addr.String(), hex.EncodeToString(buff))
		}

	}
}
