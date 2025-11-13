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

var partsPool = sync.Pool{
	New: func() any { return make([][]byte, 0, 8) },
}
var leftbyte = make([]byte, 4096)
var sep = []byte{0x57, 0xab} // ch9329固定分隔符

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
		blen, _, err := self.conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println("buffudp:", buff[:blen])

		//先拼接一下避免原来那种缓冲方法造成拷贝
		data := append(leftbyte, buff[:blen]...)

		parts := partsPool.Get().([][]byte)[:0]
		start := 0

		for i := 0; i < len(data)-1; i++ {
			if data[i] == sep[0] && data[i+1] == sep[1] {
				if i > start {
					// 这里是零拷贝切片
					part := data[start:i]
					parts = append(parts, part)
				}
				start = i + 2
				i++
			}
		}
		// 处理最后的残留
		if start < len(data) {
			leftbyte = data[start:]
		} else {
			leftbyte = leftbyte[:0]
		}

		// 把每个分片直接丢到 DataChannel（零拷贝）
		for _, p := range parts {
			DataChannel <- p
		}

		// 放回 pool
		partsPool.Put(parts[:0])
	}
}

func (self *ServerUDP) udpchannel() {
	for keyboard := range netSender.Ctx.UdpChannel {
		//fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
		self.conn.WriteTo(keyboard, self.SendServer)
	}
}
