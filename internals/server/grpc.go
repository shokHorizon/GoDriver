package server

import (
	"context"
	"strconv"

	"github.com/shokHorizon/GoDriver/internals/driver"
	"github.com/shokHorizon/GoDriver/stream"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	stream.ApiCallerScaleServer
	Device driver.Device
}

func NewGRPCServer(d driver.Device) *grpc.Server {
	gsrv := grpc.NewServer()
	gServer := &GRPCServer{Device: d}
	stream.RegisterApiCallerScaleServer(gsrv, gServer)
	return gsrv
}

func (s *GRPCServer) SetTare(ctx context.Context, _ *stream.Empty) (*stream.ResponseSetScale, error) {
	msg, err := s.Device.SetTare(ctx)
	if err != nil {
		return nil, err
	}
	return &stream.ResponseSetScale{Error: msg}, err
}
func (s *GRPCServer) SetTareValue(ctx context.Context, tVal *stream.RequestTareValue) (*stream.ResponseSetScale, error) {
	i, err := strconv.ParseInt(tVal.String(), 10, 32)
	if err != nil {
		return nil, err
	}
	res, err := s.Device.SetTareValue(ctx, int32(i))
	return &stream.ResponseSetScale{Error: res}, err
}
func (s *GRPCServer) SetZero(context.Context, *stream.Empty) (*stream.ResponseSetScale, error) {
	return nil, nil
}
func (s *GRPCServer) GetInstantWeight(context.Context, *stream.Empty) (*stream.ResponseInstantWeight, error) {
	return nil, nil
}
func (s *GRPCServer) GetState(context.Context, *stream.Empty) (*stream.ResponseScale, error) {
	return nil, nil
}
