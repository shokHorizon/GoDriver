package command

type Head struct {
	Headers [3]byte
	Len     int16
	Crc16   int16
}

type SetTareBody struct {
	Code int16
	Tare int32
}

type AckSetTareBody struct {
	Code int16
}

var HEADERS = [3]byte{0xF8, 0x55, 0xCE}

const (
	CMD_GET_SCALE_PAR = 0x75
	CMD_ACK_SCALE_PAR = 0x76
	CMD_ERROR         = 0x28
	CMD_GET_MASSA     = 0x23
	CMD_ACK_MASSA     = 0x24
	CMD_SET_TARE      = 0xA3
	CMD_ACK_SET_TARE  = 0x12
	CMD_NACK_TARE     = 0x15
	CMD_SET_ZERO      = 0x72
	CMD_ACK_SET       = 0x27
	CMD_GET_NAME      = 0x20
	CMD_ACK_NAME      = 0x21
	CMD_SET_NAME      = 0x22
	CMD_GET_ETHERNET  = 0x2D
	CMD_ACK_ETHERNET  = 0x2E
	CMD_SET_ETHERNET  = 0x39
	CMD_GET_WIFI_IP   = 0x33
	CMD_ACK_WIFI_IP   = 0x34
	CMD_SET_WIFI_IP   = 0x31
	CMD_GET_WIFI_SSID = 0x3A
	CMD_SET_WIFI_SSID = 0x3C
	CMD_NACK          = 0x3F
)
