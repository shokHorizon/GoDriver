package server

import (
	"context"
	"io"

	"github.com/shokHorizon/GoDriver/internals/driver"
	pb "github.com/shokHorizon/GoDriver/stream"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.ApiCallerScaleServer
	Device driver.Device
}

func NewGRPCServer(d driver.Device) *grpc.Server {
	gsrv := grpc.NewServer()
	gServer := &GRPCServer{Device: d}
	pb.RegisterApiCallerScaleServer(gsrv, gServer)
	return gsrv
}

func (s *GRPCServer) ScalesMessageOutChannel(stream pb.ApiCallerScale_ScalesMessageOutChannelServer) error {
	for {
		_, err := stream.Recv()
		if err != io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		res, err := s.Device.GetScalePar(stream.Context())
		if err != nil {
			return err
		}
		resp := &pb.ResponseScale{Message: res}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func (s *GRPCServer) SetTare(ctx context.Context, _ *pb.Empty) (*pb.ResponseSetScale, error) {
	msg, err := s.Device.SetTare(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ResponseSetScale{Error: msg}, err
}
func (s *GRPCServer) SetTareValue(ctx context.Context, tVal *pb.RequestTareValue) (*pb.ResponseSetScale, error) {
	var val [4]byte
	copy(val[:], tVal.Message[0:3])
	res, err := s.Device.SetTareValue(ctx, val)
	return &pb.ResponseSetScale{Error: res}, err
}
func (s *GRPCServer) SetZero(ctx context.Context, _ *pb.Empty) (*pb.ResponseSetScale, error) {
	res, err := s.Device.SetZero(ctx)
	return &pb.ResponseSetScale{Error: res}, err
}
func (s *GRPCServer) GetInstantWeight(ctx context.Context, _ *pb.Empty) (*pb.ResponseInstantWeight, error) {
	res, err := s.Device.GetInstantWeight(ctx)
	return &pb.ResponseInstantWeight{Error: res}, err
}
func (s *GRPCServer) GetState(ctx context.Context, _ *pb.Empty) (*pb.ResponseScale, error) {
	res, err := s.Device.GetScalePar(ctx)
	if err != nil {
		return &pb.ResponseScale{Error: res}, err
	}
	return &pb.ResponseScale{Message: res}, err
}
