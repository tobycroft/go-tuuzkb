package netTcp

import (
	"bytes"
	"main.go/netReceiver"
	"main.go/netSender"
	"net"
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

	go func() {
		for keyboard := range netSender.Ctx.UdpChannel {
			//fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
			self.conn.WriteTo(keyboard, self.SendServer)
		}
	}()

	buff := make([]byte, 1024)
	buffer := bytes.Buffer{}
	for {
		blen, addr, err := self.conn.ReadFrom(buff)
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
				netReceiver.Crx.MessageRouter(buffer.Bytes(), addr)
				buffer.Next(buffer.Len())
				break
			} else {
				segment := data[:idx]
				if len(segment) > 0 {
					//fmt.Println("Processed:", segment)
					//fmt.Println(conn.RemoteAddr().String(), hex.EncodeToString(buff))
					//if addr.String() == "10.0.0.91:6666" {
					netReceiver.Crx.MessageRouter(segment, addr)
				}
				buffer.Next(idx + 2) // 跳过 `0x57 0xab` 分隔符
			}
		}
	}
}
