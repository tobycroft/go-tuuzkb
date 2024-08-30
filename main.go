package main

import (
	"github.com/tarm/serial"
	"log"
	"main.go/datastruct"
)

func main() {
	//c := &serial.Config{Name: "/dev/ttyS5", Baud: 9600}
	c := &serial.Config{Name: "/dev/ttyS5", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	kb := datastruct.Kb{
		SerialPort: s,
	}
	kb.CmdGetUsbString()
	//kb.CmdSetUsbString(1, "2.4G Wireless Receiver")
	//kb.CmdSetParaCfg()
	//kb.CmdReadMyHidData()
	//kb.CmdSetDefaultCfg()
	//kb.CmdReset()
	//fmt.Println(kb.Sendbuf.Bytes())

	kb.Crx.CmdGetUsbStringRecv(kb.Recv())
	//kb.Recv()

}
