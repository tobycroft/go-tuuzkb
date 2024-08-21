package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"main.go/datastruct"
	"main.go/define/cmd"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyS5", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	//txbuf := []byte{0xAA, 0x01, 0x0f, 0x00, 0x00, 0xBA}
	v1 := datastruct.ClientTx{
		Cmd: cmd.CMD_GET_PARA_CFG,
	}
	v1.CalcHead()
	var sendbuf bytes.Buffer
	err = binary.Write(&sendbuf, binary.BigEndian, &v1)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	err = binary.Write(&sendbuf, binary.BigEndian, [0]byte{})
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	sum := byte(0x00)
	for _, b := range sendbuf.Bytes() {
		sum = sum + b
	}
	err = binary.Write(&sendbuf, binary.BigEndian, byte(sum&0xFF))
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	fmt.Println(sendbuf.Bytes())

	n, err := s.Write(sendbuf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%X\n", buf[:n])

}
