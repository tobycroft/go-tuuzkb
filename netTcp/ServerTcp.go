package netTcp

import (
	"bufio"
	"bytes"
	"main.go/netReceiver"
	"main.go/netSender"
	"net"
	"sync"
)

type ServerTcp struct {
	SendServer net.Addr

	Kb netReceiver.KeyBoard
}

var addrToConn sync.Map
var addrToLock sync.Map

func (self *ServerTcp) Start() *ServerTcp {
	Conn, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err.Error())
	}
	go self.tcpchannel()
	for {
		conn, err := Conn.Accept()
		if err != nil {
			panic(err.Error())
		}
		reader := bufio.NewReader(conn)
		addrToConn.Store(conn.RemoteAddr().String(), conn)
		addrToLock.Store(conn.RemoteAddr().String(), &sync.Mutex{})
		go self.handler(conn, reader)
	}
}

func (self *ServerTcp) handler(conn net.Conn, reader *bufio.Reader) {
	buff := make([]byte, 128)
	for {
		_, err := reader.Read(buff)
		if err != nil {
			addrToConn.Delete(conn.RemoteAddr().String())
			addrToLock.Delete(conn.RemoteAddr().String())
			return
		}
		//fmt.Println(conn.RemoteAddr().String(), hex.EncodeToString(buff))
		//if addr.String() == "10.0.0.91:6666" {
		slice_byte := bytes.Split(buff, []byte{0x57, 0xab})
		for _, ddd := range slice_byte {
			netReceiver.Crx.MessageRouter(ddd, conn.RemoteAddr())
		}
	}

}

func (self *ServerTcp) tcpchannel() {
	for keyboard := range netSender.Ctx.TxChannel {
		//fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
		addrToConn.Range(func(key, value interface{}) bool {
			if self.SendServer.String() == key.(string) {
				locker, ok := addrToLock.Load(key)
				if ok {
					locker.(*sync.Mutex).Lock()
					_, err := value.(net.Conn).Write(keyboard)
					locker.(*sync.Mutex).Unlock()
					if err != nil {
						addrToConn.Delete(key)
						addrToLock.Delete(key)
					}
				}
			}
			return true
		})

	}
}
