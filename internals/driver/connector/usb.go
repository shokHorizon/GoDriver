package connector

import (
	"errors"

	"github.com/karalabe/usb"
	"github.com/shokHorizon/GoDriver/internals/config"
)

func NewUSBConn(cfg config.Config) (Connector, error) {
	if cfg.Vendor == 0 && cfg.ProductId == 0 {
		return nil, errors.New("Device or vendor id is not specified")
	}
	devicesInfo, err := usb.Enumerate(cfg.Vendor, cfg.ProductId)
	if err != nil {
		return nil, err
	}
	if len(devicesInfo) == 0 {
		return nil, errors.New("No devices found")
	}
	return devicesInfo[0].Open()
}
