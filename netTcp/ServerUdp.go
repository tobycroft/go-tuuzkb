package netTcp

import (
	"bytes"
	"net"

	"main.go/netSender"
)

type ServerUDP struct {
	SendServer net.Addr
	conn       net.PacketConn
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

	buffer := bytes.Buffer{}
	for {
		blen, _, err := self.conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}
		buffer.Write(buff[:blen])
		//fmt.Println("buffudp:", buff[:blen])

		for {
			data := buffer.Bytes() // 获取当前缓冲区中的所有数据
			idx := bytes.Index(data, []byte{0x57, 0xab})
			//fmt.Println("idx:", idx)
			if idx == -1 {
				break
			} else if idx == 0 {
				buffer.Next(2)
				//fmt.Println("bufftcp-deal:", buffer.Bytes(), buffer.Len())
				//go netReceiver.Crx.MessageRouter(buffer.Bytes())
				DataChannel <- buffer.Bytes()
				buffer.Next(buffer.Len())
				break
			} else {
				segment := data[:idx]
				if len(segment) > 0 {
					//fmt.Println("Processed:", segment)
					//fmt.Println(conn.RemoteAddr().String(), hex.EncodeToString(buff))
					//if addr.String() == "10.0.0.91:6666" {
					//go netReceiver.Crx.MessageRouter(segment)
					DataChannel <- segment
				}
				buffer.Next(idx + 2) // 跳过 `0x57 0xab` 分隔符
			}
		}
	}
}

func (self *ServerUDP) udpchannel() {
	for keyboard := range netSender.Ctx.UdpChannel {
		//fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
		self.conn.WriteTo(keyboard, self.SendServer)
	}
}
