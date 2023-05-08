package driver

import (
	"context"
	"errors"

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

func (d *Device) SetTareValue(ctx context.Context, val [4]byte) (string, error) {
	h := command.NewSetTareHead(val)
	err := h.Write(d.Conn)
	if err != nil {
		return "", err
	}

	rh, err := command.ReadHead(d.Conn)
	if err != nil {
		return "", err
	}

	if rh.Message[0] == command.CMD_ACK_SET_TARE {
		return string(rh.Message), nil
	}

	if rh.Message[0] == command.CMD_ERROR {
		return command.GetErrorMesage(rh.Message[1]), errors.New("cmd error")
	}

	return "", errors.New("unknown error")
}

func (d *Device) SetZero(ctx context.Context) (string, error) {
	h := command.NewSetZeroHead()
	err := h.Write(d.Conn)
	if err != nil {
		return "", err
	}

	rh, err := command.ReadHead(d.Conn)
	if err != nil {
		return "", err
	}

	if rh.Message[0] == command.CMD_ACK_SET {
		return string(rh.Message), nil
	}

	if rh.Message[0] == command.CMD_ERROR {
		return command.GetErrorMesage(rh.Message[1]), errors.New("cmd error")
	}

	return "", errors.New("unknown error")
}

func (d *Device) GetInstantWeight(ctx context.Context) (string, error) {
	h := command.NewGetInstantWeightHead()
	err := h.Write(d.Conn)
	if err != nil {
		return "", err
	}

	rh, err := command.ReadHead(d.Conn)
	if err != nil {
		return "", err
	}

	if rh.Message[0] == command.CMD_ACK_MASSA {
		return string(rh.Message), nil
	}

	if rh.Message[0] == command.CMD_ERROR {
		return command.GetErrorMesage(rh.Message[1]), errors.New("cmd error")
	}

	return "", errors.New("unknown error")
}

func (d *Device) GetScalePar(ctx context.Context) (string, error) {
	h := command.NewGetScaleHead()
	err := h.Write(d.Conn)
	if err != nil {
		return "", err
	}

	rh, err := command.ReadHead(d.Conn)
	if err != nil {
		return "", err
	}

	if rh.Message[0] == command.CMD_ACK_SCALE_PAR {
		return string(rh.Message), nil
	}

	if rh.Message[0] == command.CMD_ERROR {
		return command.GetErrorMesage(rh.Message[1]), errors.New("cmd error")
	}

	return "", errors.New("unknown error")
}
