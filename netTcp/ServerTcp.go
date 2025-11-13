package netTcp

import (
	"bufio"
	"net"
	"sync"

	"main.go/netSender"
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
	addr := conn.RemoteAddr().String()
	for {
		blen, err := reader.Read(buff)
		if err != nil {
			go addrToConn.Delete(addr)
			go addrToLock.Delete(addr)
			return
		}
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
