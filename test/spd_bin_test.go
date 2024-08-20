package test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/datastruct"
	"math/rand"
	"testing"
)

func BenchmarkBin(b *testing.B) {
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tx := datastruct.ClientTx{
			Head: struct {
				Mac      uint32
				Rand     uint32
				Indexpts uint32
				Cmd      uint32
			}{Mac: rand.Uint32(), Rand: rand.Uint32(), Indexpts: uint32(i), Cmd: rand.Uint32()},
			CmdKeyboard: struct {
				Ctrl   byte
				Resv   byte
				Button [10]byte
			}{Ctrl: 1, Resv: 0, Button: [10]byte{1, 2, 0, 0, 0, 0, 0, 0, 0, 0}},
		}

		var buf bytes.Buffer
		err := binary.Write(&buf, binary.NativeEndian, &tx.Head)
		if err != nil {
			panic(fmt.Sprintln("binary编译失败", err))
		}
		err = binary.Write(&buf, binary.NativeEndian, &tx.CmdKeyboard)
		if err != nil {
			panic(fmt.Sprintln("binary编译失败", err))
		}
		buf.Bytes()
	}

	b.StopTimer()
}
