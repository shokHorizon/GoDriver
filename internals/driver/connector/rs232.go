package connector

import (
	"errors"

	"github.com/shokHorizon/GoDriver/internals/config"
	"github.com/tarm/serial"
)

func NewRsConnection(cfg config.Config) (Connector, error) {
	sC := serial.Config{StopBits: 1, Name: cfg.ComName}

	if cfg.ComName == "" {
		return nil, errors.New("Com name cannot be empty")
	}

	switch cfg.ComProtocol {
	case "1c":
		sC.Baud = 57600
		sC.Parity = serial.ParityNone
	case "2":
		sC.Baud = 4800
		sC.Parity = serial.ParityEven
	case "Stndr":
		sC.Baud = 19200
		sC.Parity = serial.ParitySpace
	default:
		return nil, errors.New("Incorrect COM_PROTOCOL parameter")
	}

	return serial.OpenPort(&sC)
}
