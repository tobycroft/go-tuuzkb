package netSender

type StandardKeyboardReport struct {
	Buttons  uint8
	ReportID uint8
	Data     [6]uint8
}
