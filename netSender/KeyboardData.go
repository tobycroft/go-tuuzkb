package netSender

type KeyboardData2 struct {
	Ctrl   byte
	Resv   byte
	Button [6]byte
}

type KeyboardData struct {
	Ctrl byte

	Resv byte

	Button0 byte
	Button1 byte
	Button2 byte
	Button3 byte
	Button4 byte
	Button5 byte
}

type KbMediaData struct {
	Mediabyte1 byte
	Mediabyte2 byte
}
