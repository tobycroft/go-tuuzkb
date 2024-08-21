package datastruct

const start1 = 0x57
const start2 = 0xab

// 发送数据联合体
type ClientTx struct {
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
}

type Clen struct {
	Data [0]byte // 后续数据 (0到64个字节)
}

type Csum struct {
	Sum byte // 校验和 (1个字节)
}

func (self *ClientTx) CalcHead() {
	self.Head = uint16(start1)<<8 | uint16(start2)
}

func (self *ClientTx) CalcSum() {

}
