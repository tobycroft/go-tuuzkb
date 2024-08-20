package datastruct

// 发送数据联合体
type ClientTx struct {
	Head uint16 // 帧头 (2个字节)
	Addr byte   // 地址码 (1个字节)
	Cmd  byte   // 命令码 (1个字节)
	Len  byte   // 后续数据长度 (1个字节)
	Data []byte // 后续数据 (0到64个字节)
	Sum  byte   // 累加和 (1个字节)
}
