package driver

import (
	"context"
	"errors"
	"strconv"

	"github.com/howeyc/crc16"
	"github.com/shokHorizon/GoDriver/internals/config"
	"github.com/shokHorizon/GoDriver/internals/driver/command"
	"github.com/shokHorizon/GoDriver/internals/driver/connector"
)

type Device struct {
	Conn connector.Connector
}

func ConnDevice(cfg config.Config) (Device, error) {
	conn, err := newConnector(cfg)
	return Device{conn}, err
}

func newConnector(cfg config.Config) (connector.Connector, error) {
	switch cfg.Device_interface {
	case "usb":
		return connector.NewUSBConn(cfg)
	case "rs232":
		return connector.NewRsConnection(cfg)
	}
	return nil, errors.New("unknown interface")
}

func (d *Device) SetTare(ctx context.Context) (string, error) {
	return "Команда не поддерживается", nil
}

func (d *Device) SetTareValue(ctx context.Context, val int32) (string, error) {
	b := command.SetTareBody{
		Code: command.CMD_SET_TARE,
		Tare: val,
	}
	msg := command.Head{
		Headers: command.HEADERS,
		Len:     int16(5),
	}
	d.Conn.Write(msg.Headers[:])
	d.Conn.Write([]byte(string(msg.Len)))
	d.Conn.Write([]byte(string(b.Code)))
	d.Conn.Write([]byte(string(b.Tare)))
	d.Conn.Write([]byte(string(crc16.Checksum([]byte(string(b.Code)+string(b.Tare)), crc16.CCITTTable))))

	var buf = make([]byte, 256)
	if n, err := d.Conn.Read(buf); err != nil || n == 0 {
		return "Соединение потеряно", nil
	}
	msgCode, _ := strconv.ParseInt(string(buf[5]), 10, 16)
	if msgCode != command.CMD_ACK_SET_TARE {
		return "Не удалось установить параметры", nil
	}
	return "", nil
}
