package netSender

type MouseData2 struct {
	Ctrl   byte
	Resv   byte
	Button [6]byte
}

type MouseData struct {
	Left   byte
	Right  byte
	Middle byte
	X      byte
	Y      byte
	Wheel  byte
}
