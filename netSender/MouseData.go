package netSender

type MouseData struct {
	ButtonBits byte
	X          byte
	Y          byte
	Wheel      byte
}

type MouseData2 struct {
	Resv       byte
	ButtonBits byte
	X          byte
	Y          byte
	Wheel      byte
}
