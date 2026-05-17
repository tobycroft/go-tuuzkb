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

var (
	buff = make([]byte, 10240) // 保持向后兼容
	tcpBufferPool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte, 10240)
			return &buf
		},
	}
)

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
	addr := conn.RemoteAddr()
	for {
		localBuff := tcpBufferPool.Get().(*[]byte)
		blen, err := reader.Read(*localBuff)
		if err != nil {
			tcpBufferPool.Put(localBuff)
			addrToConn.Delete(addr)
			addrToLock.Delete(addr)
			return
		}
		//fmt.Println("bufftcp:", blen, (*localBuff)[:blen])

		//(*localBuff)[:blen]就已经是究极最少得状态了，不需要额外切分了
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
		tcpBufferPool.Put(localBuff)
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