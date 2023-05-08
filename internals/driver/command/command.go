package command

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Head struct {
	Headers [3]byte
	Len     int16
	Message []byte
	Crc16   int16
}

func NewHead() *Head {
	return &Head{
		Headers: HEADERS,
	}
}

func ReadHead(buf io.Reader) (*Head, error) {
	h := &Head{}
	buf.Read(h.Headers[:])

	err := binary.Read(buf, binary.BigEndian, &h.Headers)
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	err = binary.Read(buf, binary.BigEndian, &h.Len)
	if err != nil {
		return nil, fmt.Errorf("failed to read len: %w", err)
	}
	h.Message = make([]byte, h.Len)
	err = binary.Read(buf, binary.BigEndian, &h.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to read message: %w", err)
	}

	return h, nil
}

func (h *Head) Write(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, h.Headers)
	if err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, h.Len)
	if err != nil {
		return fmt.Errorf("failed to write len: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, h.Message)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, h.Crc16)
	if err != nil {
		return fmt.Errorf("failed to write crc16: %w", err)
	}
	return nil
}

func (h *Head) putMessage(m []byte) {
	h.Message = []byte(m)
	h.Len = int16(len(m))
	h.setCrc()
}

func (h *Head) setCrc() {

}

func NewSetTareHead(tVal [4]byte) *Head {
	h := NewHead()
	msg := make([]byte, 5)
	msg[0] = CMD_SET_TARE
	copy(msg[1:], tVal[:])
	h.putMessage(msg)
	return h
}

func NewSetZeroHead() *Head {
	h := NewHead()
	msg := make([]byte, 1)
	msg[0] = CMD_SET_ZERO
	h.putMessage(msg)
	return h
}

func NewGetInstantWeightHead() *Head {
	h := NewHead()
	msg := make([]byte, 1)
	msg[0] = CMD_GET_MASSA
	h.putMessage(msg)
	return h
}

func NewGetScaleHead() *Head {
	h := NewHead()
	msg := make([]byte, 1)
	msg[0] = CMD_GET_SCALE_PAR
	h.putMessage(msg)
	return h
}

func GetErrorMesage(code byte) string {
	switch code {
	case 0x07:
		return "Команда не поддерживается"
	case 0x08:
		return "Нагрузка на весовом устройстве превышает НПВ"
	case 0x09:
		return "Весовое устройство не в режиме взвешивания"
	case 0x0A:
		return "Ошибка входных данных"
	case 0x0B:
		return "Ошибка сохранения данных"
	case 0x10:
		return "Интерфейс WiFi не поддерживается"
	case 0x11:
		return "Интерфейс Ethernet не поддерживается"
	case 0x15:
		return "Установка >0< невозможна"
	case 0x17:
		return "Нет связи c модулем взвешивающим"
	case 0x18:
		return "Установлена нагрузка на платформу при включении весового устройства"
	case 0x19:
		return "Весовое устройство неисправно"
	case 0xF0:
		return "Неизвестная ошибка"
	}
	return "Такой ошибки даже в доке нет..."
}

var HEADERS = [3]byte{0xF8, 0x55, 0xCE}

const (
	CMD_GET_SCALE_PAR byte = 0x75
	CMD_ACK_SCALE_PAR byte = 0x76
	CMD_ERROR         byte = 0x28
	CMD_GET_MASSA     byte = 0x23
	CMD_ACK_MASSA     byte = 0x24
	CMD_SET_TARE      byte = 0xA3
	CMD_ACK_SET_TARE  byte = 0x12
	CMD_NACK_TARE     byte = 0x15
	CMD_SET_ZERO      byte = 0x72
	CMD_ACK_SET       byte = 0x27
	CMD_GET_NAME      byte = 0x20
	CMD_ACK_NAME      byte = 0x21
	CMD_SET_NAME      byte = 0x22
	CMD_GET_ETHERNET  byte = 0x2D
	CMD_ACK_ETHERNET  byte = 0x2E
	CMD_SET_ETHERNET  byte = 0x39
	CMD_GET_WIFI_IP   byte = 0x33
	CMD_ACK_WIFI_IP   byte = 0x34
	CMD_SET_WIFI_IP   byte = 0x31
	CMD_GET_WIFI_SSID byte = 0x3A
	CMD_SET_WIFI_SSID byte = 0x3C
	CMD_NACK          byte = 0x3F
)
