package netSender

var Ctx = &ClientTx{}

type ClientTx struct {
	TxChannel      chan []byte
	TcpChannel     chan []byte
	UdpChannel     chan []byte
	MouseTxChannel chan any
}

func (self *ClientTx) Ready() {
	self.MouseTxChannel = make(chan any, 4)
	self.TxChannel = make(chan []byte, 4)
	self.TcpChannel = make(chan []byte, 4)
	self.UdpChannel = make(chan []byte, 4)
	go func() {
		for c := range self.TxChannel {
			self.UdpChannel <- c
			self.TcpChannel <- c
		}
	}()
}
