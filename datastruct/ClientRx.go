package datastruct

import "log"

type ClientRx struct {
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
}

func (kb *Kb) Recv() []byte {
	buf := make([]byte, 128)
	n, err := kb.SerialPort.Read(buf)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%X\n", buf[:n])
	return buf
}
