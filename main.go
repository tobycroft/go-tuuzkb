package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"main.go/datastruct"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyS5", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	//txbuf := []byte{0xAA, 0x01, 0x0f, 0x00, 0x00, 0xBA}
	ctx := datastruct.ClientTx{}
	data := [0]byte{}
	ctx.Len = byte(len(data))

	fmt.Println(sendbuf.Bytes())

	n, err := s.Write(sendbuf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 256)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%X\n", buf[:n])

}
