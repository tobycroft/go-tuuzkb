package main

import (
	"github.com/tarm/serial"
	"log"
	"main.go/datastruct"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyS5", Baud: 9600}
	//c := &serial.Config{Name: "/dev/ttyS5", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	kb := datastruct.Kb{}
	kb.CmdGetUsbString()
	//kb.CmdSetUsbString()
	//kb.CmdSetParaCfg()
	//kb.CmdReadMyHidData()
	//kb.CmdSetDefaultCfg()
	//kb.CmdReset()
	//fmt.Println(kb.Sendbuf.Bytes())

	n, err := s.Write(kb.Sendbuf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%X\n", buf[:n])
	kb.Crx.CmdGetUsbStringRecv(buf)
	//kb = datastruct.Kb{}
	////kb.CmdGetParaCfg()
	//kb.CmdReset()
	//fmt.Println(kb.Sendbuf.Bytes())
	//
	//n, err = s.Write(kb.Sendbuf.Bytes())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//buf = make([]byte, 128)
	//n, err = s.Read(buf)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%X\n", buf[:n])
	//kb.Crx.CmdGetParaCfgRecv(buf)

}
