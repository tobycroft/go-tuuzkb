package datastruct

import (
	"encoding/binary"
	"fmt"
)

func (kb *Kb) CmdGetUsb() *Kb {
	err := binary.Write(&kb.Sendbuf, binary.BigEndian, &kb.Ctx)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	return kb
}
