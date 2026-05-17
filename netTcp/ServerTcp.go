package netTcp

import (
	"bufio"
	"fmt"
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
	addr := conn.RemoteAddr()
	for {
		blen, err := reader.Read(buff)
		if err != nil {
			go addrToConn.Delete(addr)
			go addrToLock.Delete(addr)
			return
		}
		//fmt.Println("bufftcp:", blen, buff[:blen])

		//buff[:blen]就已经是究极最少得状态了，不需要额外切分了
		udppack := buff[:blen]
		//fmt.Println("buffudp:", blen, udppack, addr.String())
		switch len(udppack) {
		case 0:
			break // 空包，直接丢弃
		case 1:
			recmap, ok := receiverMap[addr.String()]
			if !ok {
				recmap = receiverIndex{
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
					recmap = receiverIndex{
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
