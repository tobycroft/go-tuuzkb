package netReceiver

type KeyBoard struct {
	KeyCurrent    uint8
	KeyDown       uint8
	KeyUp         uint8
	KeyPressDebug bool
	// 字母键
	A bool
	B bool
	C bool
	D bool
	E bool
	F bool
	G bool
	H bool
	I bool
	J bool
	K bool
	L bool
	M bool
	N bool
	O bool
	P bool
	Q bool
	R bool
	S bool
	T bool
	U bool
	V bool
	W bool
	X bool
	Y bool
	Z bool

	// 数字键盘
	Num0 bool
	Num1 bool
	Num2 bool
	Num3 bool
	Num4 bool
	Num5 bool
	Num6 bool
	Num7 bool
	Num8 bool
	Num9 bool

	// 特殊字符键
	Esc           bool // 退出键
	Space         bool // 空格键
	Enter         bool // 回车键
	Tab           bool // 制表键
	Backspace     bool // 退格键
	Plus          bool // 加号键
	Minus         bool // 减号键
	Equals        bool // 等号键
	Semicolon     bool // 分号键
	Quote         bool // 引号键
	Backslash     bool // 反斜杠键
	Comma         bool // 逗号键
	Dot           bool // 句号键
	QuestionSlash bool // 斜杠键
	Backquote     bool // 反引号键
	BraceL        bool //左花括号
	BraceR        bool //右花括号

	// 功能键
	F1  bool
	F2  bool
	F3  bool
	F4  bool
	F5  bool
	F6  bool
	F7  bool
	F8  bool
	F9  bool
	F10 bool
	F11 bool
	F12 bool

	// 控制键
	LeftCtrl     bool
	LeftShift    bool
	LeftAlt      bool
	LeftWindows  bool
	RightWindows bool
	RightCtrl    bool
	RightShift   bool
	RightAlt     bool
	CapsLock     bool // 大写锁定键

	// 其他按键...
	ScrollLock  bool
	NumLock     bool
	PrintScreen bool
	PauseBreak  bool
	Insert      bool
	Delete      bool
	Home        bool
	End         bool
	PageUp      bool
	PageDown    bool
	Application bool

	// 小数字键盘区
	NumPad0        bool
	NumPad1        bool
	NumPad2        bool
	NumPad3        bool
	NumPad4        bool
	NumPad5        bool
	NumPad6        bool
	NumPad7        bool
	NumPad8        bool
	NumPad9        bool
	NumPadPlus     bool // 小数字键盘上的加号键
	NumPadMinus    bool // 小数字键盘上的减号键
	NumPadMultiply bool // 小数字键盘上的乘号键
	NumPadDivide   bool // 小数字键盘上的除号键
	NumPadEnter    bool // 小数字键盘上的回车键
	NumPadDot      bool // 小数字键盘上的小数点键

	//方向键
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
}
