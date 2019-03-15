package service

import (
	"context"
	pb "learn/api-gateway/example/echo/service"

	"github.com/gogo/protobuf/types"
)

type SimpleEchoServer struct {
}

func (s *SimpleEchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Text: req.Text,
	}, nil
}

func (s *SimpleEchoServer) Ping(ctx context.Context, req *types.Empty) (*types.Timestamp, error) {
	return types.TimestampNow(), nil

}
