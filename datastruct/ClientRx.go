package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ClientRx struct {
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
}

func (kb *Kb) Recv(buf []byte) [50]byte {
	bs := bytes.NewReader(buf)
	crx := ClientRx{}
	binary.Read(bs, binary.BigEndian, &crx)
	dats := [50]byte{}
	binary.Read(bs, binary.BigEndian, &dats)
	fmt.Println(crx)
	fmt.Println(dats)
	return dats
}
