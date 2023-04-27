package config

import "github.com/Netflix/go-env"

type Config struct {
	Grpc_port        string `env:"gRPC_PORT"`
	Device_interface string `env:"INTERFACE"`
	ComProtocol      string `env:"COM_PROTOCOL"`
	ComName          string `env:"COM_NAME"`
	Ip               string `env:"DEVICE_IP"`
	Mask             string `env:"DEVICE_MASK"`
	Port             int    `env:"DEVICE_PORT"`
	DefaultGateway   string `env:"DEVICE_GW"`
	Password         string `env:"DEVICE_PSWD"`
	Vendor           uint16 `env:"VENDOR"`
	ProductId        uint16 `env:"PRODUCT_ID"`
}

var instance *Config

func GetConfig() (Config, error) {
	if instance == nil {
		cfg := Config{}
		_, err := env.UnmarshalFromEnviron(&cfg)
		if err != nil {
			return cfg, err
		}
		instance = &cfg
	}
	return *instance, nil
}
