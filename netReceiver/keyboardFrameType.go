package netReceiver

type IDByte struct {
	Bits763 uint8 // Bit7&6&3
	Type    uint8 // Bit5&4: 01键盘 2鼠标 3多媒体 00其他
	Bits21  uint8 // Bit2&1 01：HID 2：BIOS 00：未知 3：保留
	Port    uint8 // Bit0: 0端口1 1端口2
}

func ParseIDByte(b byte) IDByte {
	bits763 := ((b>>7)&0x01)<<2 | ((b>>6)&0x01)<<1 | ((b >> 3) & 0x01)
	typ := (b >> 4) & 0x03
	bits21 := (b >> 1) & 0x03
	port := b & 0x01

	return IDByte{
		Bits763: bits763,
		Type:    typ,
		Bits21:  bits21,
		Port:    port,
	}
}
