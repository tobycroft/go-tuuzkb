package netSender

var Ctx = &ClientTx{}

type ClientTx struct {
	TxChannel      chan []byte
	TcpChannel     chan []byte
	UdpChannel     chan []byte
	MouseTxChannel chan any
}

func (self *ClientTx) Ready() {
	self.MouseTxChannel = make(chan any)
	self.TxChannel = make(chan []byte)
	self.TcpChannel = make(chan []byte, 1)
	self.UdpChannel = make(chan []byte, 1)
	go func() {
		for c := range self.TxChannel {
			self.UdpChannel <- c
			self.TcpChannel <- c
		}
	}()
}
