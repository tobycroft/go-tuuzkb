package cmd

const (
	//获取芯片版本等信息
	CMD_GET_INFO = 0x01

	//发送USB键盘普通数据
	CMD_SEND_KB_GENERAL_DATA = 0x02

	//发送USB键盘多媒体数据
	CMD_SEND_KB_MEDIA_DATA = 0x03

	//发送USB绝对鼠标数据
	CMD_SEND_MS_ABS_DATA = 0x04

	//发送USB相对鼠标数据
	CMD_SEND_MS_REL_DATA = 0x05

	//发送USB自定义HID设备数据
	CMD_SEND_MY_HID_DATA = 0x06

	//读取USB自定义HID设备数据
	CMD_READ_MY_HID_DATA = 0x87

	//获取参数配置
	CMD_GET_PARA_CFG = 0x08

	//设置参数配置
	CMD_SET_PARA_CFG = 0x09

	//获取字符串描述符配置
	CMD_GET_USB_STRING_DESC = 0x0A

	//设置字符串描述符配置
	CMD_SET_USB_STRING = 0x0B

	//恢复出厂默认配置
	CMD_SET_DEFAULT_CFG = 0x0C

	//复位芯片
	CMD_RESET = 0x0F
)
