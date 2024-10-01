package netTcp

import (
	"bytes"
	"fmt"
	"main.go/action"
	"main.go/netReceiver"
	"net"
	"time"
)

func ClientRx() {

	buff := make([]byte, 512)

	network := net.TCPAddr{
		IP:   net.ParseIP("10.0.0.91"),
		Port: 6666,
	}

	ntt, err := net.DialTCP("tcp", nil, &network)
	if err != nil {
		panic(err.Error())
	}

	tm := 0
	go func() {
		time.Sleep(1 * time.Second)
		for {
			//fmt.Println(tm)
			tm = 0
			time.Sleep(1 * time.Second)
			_, err = ntt.Write([]byte{0x57, 0xab, 0x87})
			if err != nil {
				panic(err.Error())
			}
		}
	}()
	var ns netReceiver.ClientRx
	ns.Ready()
	var run action.Runnable
	go run.MainRun(&ns)
	for {
		_, err = ntt.Read(buff)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(buff)
		tm = tm + 1
		//fmt.Println(hex.EncodeToString(buff))
		slice_byte := bytes.Split(buff, []byte{0x57, 0xab})
		for _, ddd := range slice_byte {
			ns.TtlRouter(ddd)
		}
	}
}

func ClientTx() {
	buff := make([]byte, 512)

	network := net.TCPAddr{
		IP:   net.ParseIP("10.0.0.90"),
		Port: 6666,
	}

	ntt, err := net.DialTCP("tcp", nil, &network)
	if err != nil {
		panic(err.Error())
	}

	tm := 0
	go func() {
		time.Sleep(1 * time.Second)
		for {
			//fmt.Println(tm)
			tm = 0
			time.Sleep(1 * time.Second)
			_, err = ntt.Write([]byte{0x57, 0xab, 0x87})
			if err != nil {
				panic(err.Error())
			}
		}
	}()
	var ns netReceiver.ClientRx
	ns.Ready()
	var run action.Runnable
	go run.MainRun(&ns)
	for {
		_, err = ntt.Read(buff)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(buff)
		tm = tm + 1
		//fmt.Println(hex.EncodeToString(buff))
		slice_byte := bytes.Split(buff, []byte{0x57, 0xab})
		for _, ddd := range slice_byte {
			fmt.Println("rcv:", ddd)
		}
	}
}
