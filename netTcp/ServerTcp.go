package netTcp

import (
	"encoding/hex"
	"fmt"
	"net"
)

func ServerRx() {
	buff := make([]byte, 256)
	keyboard_server, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err.Error())
	}
	for {
		Conn, err := keyboard_server.Accept()
		if err != nil {
			panic(err.Error())
		}
		//aa, err := Conn.Read(buff)
		//fmt.Println(Conn, aa, buff)
		go func() {
			for {
				aa, _ := Conn.Read(buff)
				fmt.Println(Conn.RemoteAddr(), aa, hex.EncodeToString(buff))
			}
		}()
	}

}

func ServerTx() {

}
