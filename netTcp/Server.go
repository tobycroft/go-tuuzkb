package netTcp

import (
	"fmt"
	"net"
)

func ServerRx() {
	buff := make([]byte, 512)
	keyboard_server, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err.Error())
	}
	go func() {
		for {
			Conn, err := keyboard_server.Accept()
			if err != nil {
				panic(err.Error())
			}
			_, err = Conn.Read(buff)
			fmt.Println(buff)

		}
	}()
}

func ServerTx() {

}
