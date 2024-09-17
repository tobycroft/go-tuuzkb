package netReceiver

type StandardKeyboardReport struct {
	ReportID uint8
	Buttons  uint8
	Data     [10]uint8
}
