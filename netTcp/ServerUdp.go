package netTcp

import (
	"fmt"
	"net"

	"main.go/netSender"
)

type ServerUDP struct {
	SendServer net.Addr
	conn       net.PacketConn
}
type receiverIndex struct {
	Bytes []byte
	Index int
}

var receiverMap = make(map[string]*receiverIndex)

func (self *ServerUDP) Start() *ServerUDP {
	var err error
	self.conn, err = net.ListenPacket("udp", ":6666")
	if err != nil {
		panic(err.Error())
	}

	go self.udpchannel()
	go self.udpchannel()
	go self.udpchannel()
	go self.udpchannel()

	for {
		blen, addr, err := self.conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}

		//buff[:blen]就已经是究极最少得状态了，不需要额外切分了
		udppack := buff[:blen]
		//fmt.Println("buffudp:", blen, udppack, addr.String())
		switch len(udppack) {
		case 0:
			break // 空包，直接丢弃
		case 1:
			recmap, ok := receiverMap[addr.String()]
			if !ok {
				recmap = &receiverIndex{
					Bytes: make([]byte, 1024),
					Index: 0,
				}
				receiverMap[addr.String()] = recmap
			}
			if udppack[0] == 0x57 {
				copy(recmap.Bytes[recmap.Index:], udppack)
				recmap.Index = recmap.Index + 1
				receiverMap[addr.String()] = recmap
			}
			break

		default:
			if udppack[0] == 0xab {
				recmap, ok := receiverMap[addr.String()]
				if !ok {
					recmap = &receiverIndex{
						Bytes: make([]byte, 1024),
						Index: 0,
					}
					receiverMap[addr.String()] = recmap
				}
				if recmap.Index == 1 {
					if recmap.Bytes[0] == 0x57 {
						recmap.Index = 0
						DataChannel <- udppack[1:]
						go fmt.Println("udp拼接数据:", udppack[1:])
					}
				} else {
					recmap.Index = 0
					receiverMap[addr.String()] = recmap
				}
			} else if udppack[0] == 0x57 && udppack[1] == 0xAB {
				DataChannel <- udppack[2:]
			}
			break
		}
	}
}

func (self *ServerUDP) udpchannel() {
	for keyboard := range netSender.Ctx.UdpChannel {
		// fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
		self.conn.WriteTo(keyboard, self.SendServer)
	}
}
