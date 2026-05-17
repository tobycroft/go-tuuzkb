package netTcp

import (
	"net"
	"sync"

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

var receiverMap sync.Map // 改为 sync.Map 保证并发安全

var udpBufferPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 10240)
		return &buf
	},
}

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
	go self.udpchannel()
	go self.udpchannel()
	go self.udpchannel()
	go self.udpchannel()

	for {
		localBuff := udpBufferPool.Get().(*[]byte)
		blen, addr, err := self.conn.ReadFrom(*localBuff)
		if err != nil {
			udpBufferPool.Put(localBuff)
			panic(err.Error())
		}

		//localBuff[:blen]就已经是究极最少得状态了，不需要额外切分了
		udppack := (*localBuff)[:blen]
		//fmt.Println("buffudp:", blen, udppack, addr.String())
		switch len(udppack) {
		case 0:
			break // 空包，直接丢弃
		case 1:
			var recmap receiverIndex
			val, ok := receiverMap.Load(addr.String())
			if !ok {
				recmap = receiverIndex{
					Bytes: make([]byte, 1024),
					Index: 0,
				}
			} else {
				recmap = val.(receiverIndex)
			}
			if udppack[0] == 0x57 {
				copy(recmap.Bytes[recmap.Index:], udppack)
				recmap.Index = recmap.Index + 1
				receiverMap.Store(addr.String(), recmap)
			}
			break

		default:
			if udppack[0] == 0xab {
				var recmap receiverIndex
				val, ok := receiverMap.Load(addr.String())
				if !ok {
					recmap = receiverIndex{
						Bytes: make([]byte, 1024),
						Index: 0,
					}
				} else {
					recmap = val.(receiverIndex)
				}
				if recmap.Index == 1 {
					if recmap.Bytes[0] == 0x57 {
						recmap.Index = 0
						data := make([]byte, blen-1)
						copy(data, udppack[1:])
						DataChannel <- data
						//fmt.Println("udp拼接数据:", data)
					}
				} else {
					recmap.Index = 0
					receiverMap.Store(addr.String(), recmap)
				}
			} else if udppack[0] == 0x57 && udppack[1] == 0xAB {
				data := make([]byte, blen-2)
				copy(data, udppack[2:])
				DataChannel <- data
			}
			break
		}
		udpBufferPool.Put(localBuff)
	}
}

func (self *ServerUDP) udpchannel() {
	for keyboard := range netSender.Ctx.UdpChannel {
		// fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
		self.conn.WriteTo(keyboard, self.SendServer)
	}
}