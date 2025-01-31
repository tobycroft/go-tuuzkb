package netTcp

import (
	"bufio"
	"bytes"
	"main.go/netSender"
	"net"
	"sync"
)

type ServerTcp struct {
	SendServer net.Addr
}

var addrToConn sync.Map
var addrToLock sync.Map

var buff = make([]byte, 10240)

func (self *ServerTcp) Start() *ServerTcp {
	Conn, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err.Error())
	}
	go self.tcpchannel()
	go self.tcpchannel()
	go self.tcpchannel()
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
	buffer := bytes.Buffer{}
	addr := conn.RemoteAddr().String()
	for {
		blen, err := reader.Read(buff)
		if err != nil {
			go addrToConn.Delete(addr)
			go addrToLock.Delete(addr)
			return
		}
		buffer.Write(buff[:blen])
		//fmt.Println("bufftcp:", blen, buff[:blen])
		for {
			data := buffer.Bytes() // 获取当前缓冲区中的所有数据
			idx := bytes.Index(data, []byte{0x57, 0xab})
			//fmt.Println("idx:", idx)
			if idx == -1 {
				break
			} else if idx == 0 {
				buffer.Next(2)
				//fmt.Println("bufftcp-deal:", buffer.Bytes(), buffer.Len())
				DataChannel <- buffer.Bytes()
				//go netReceiver.Crx.MessageRouter(buffer.Bytes())
				buffer.Next(buffer.Len())
				break
			} else {
				segment := data[:idx]
				if len(segment) > 0 {
					//fmt.Println("Processed:", segment)
					//fmt.Println(conn.RemoteAddr().String(), hex.EncodeToString(buff))
					//if addr.String() == "10.0.0.91:6666" {
					DataChannel <- segment
					//go netReceiver.Crx.MessageRouter(segment)
				}
				buffer.Next(idx) // 跳过 `0x57 0xab` 分隔符
			}
		}
	}
}

func (self *ServerTcp) tcpchannel() {
	for keyboard := range netSender.Ctx.TcpChannel {
		addrToConn.Range(func(key, value interface{}) bool {
			if self.SendServer.String() == key.(string) {
				locker, ok := addrToLock.Load(key)
				if ok {
					//fmt.Println("rss", keyboard, hex.EncodeToString(keyboard))
					locker.(*sync.Mutex).Lock()
					_, err := value.(net.Conn).Write(keyboard)
					//fmt.Println("sendwithn:", n, keyboard)
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
