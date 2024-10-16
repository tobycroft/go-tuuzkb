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

	Kb netReceiver.KeyBoard
}

func (self *ServerUDP) Start() *ServerUDP {
	var err error
	self.Conn, err = net.ListenPacket("udp", ":6666")
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for keyboard := range netSender.Ctx.TxChannel {
			//fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
			self.Conn.WriteTo(keyboard, self.SendServer)
		}
	}()

	buff := make([]byte, 1024)
	buffer := bytes.Buffer{}
	for {
		blen, addr, err := self.Conn.ReadFrom(buff)
		if err != nil {
			panic(err.Error())
		}
		buffer.Write(buff[:blen])

		for {
			data := buffer.Bytes() // 获取当前缓冲区中的所有数据
			idx := bytes.Index(data, []byte{0x57, 0xab})
			if idx == -1 {
				// 没有找到分隔符，等待更多数据
				break
			}
			segment := data[:idx]
			if len(segment) > 0 {
				//if addr.String() == "10.0.0.91:6666" {
				netReceiver.Crx.MessageRouter(segment, addr)
				//if addr.String() == "10.0.0.90:6666" {
				//	fmt.Println(addr.String(), hex.EncodeToString(buff))
				//}
				//} else {
				//	fmt.Println(addr.String(), hex.EncodeToString(buff))
				//}
			}
			buffer.Next(idx + 2) // 跳过 `0x57 0xab` 分隔符
		}
	}
}
