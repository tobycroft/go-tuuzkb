package main

import (
	"bytes"
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

	kb := datastruct.Kb{Sendbuf: bytes.Buffer{}}
	kb.CmdGetInfo()
	fmt.Println(kb.Sendbuf.Bytes())

	n, err := s.Write(kb.Sendbuf.Bytes())
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
