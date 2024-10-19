package netSender

var SendApi = SendFrameApi[SendFrame](&SendFrame{})

type SendFrameApi[T SendFrame | SendTx] interface {
	Head(Cmd byte) *T
	Data(data any) *T
	sum() *T
	Send()
}
