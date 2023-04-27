package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/shokHorizon/GoDriver/internals/config"
	"github.com/shokHorizon/GoDriver/internals/driver"
	"github.com/shokHorizon/GoDriver/internals/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env file: ", err)
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("Cannot parse config: ", err)
	}

	device, err := driver.ConnDevice(cfg)
	if err != nil {
		log.Fatal("Cannot connect to device: ", err)
	}
	defer device.Conn.Close()

	lis, err := net.Listen("tcp", cfg.Grpc_port)
	if err != nil {
		log.Fatal("Cannot run on port ", cfg.Grpc_port)
	}

	srv := server.NewGRPCServer(device)
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}

}
