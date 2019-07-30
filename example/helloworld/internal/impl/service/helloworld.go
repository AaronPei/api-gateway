package service

import (
	"context"
	echo "learn/api-gateway/example/echo/service"
	pb "learn/api-gateway/example/helloworld/service"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// create connection
	conn, err := grpc.Dial("echo:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	reply, er := echo.NewEchoClient(conn).Ping(ctx, &types.Empty{})
	if er != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: "Hello " + in.Name + "-" + in.Age + "-echoPing" + reply.String()}, nil
}

func (s *Server) SayBye(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name + "-" + in.Age}, nil
}

func (s *Server) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiReply, error) {
	return &pb.HiReply{ReplyContent: "This content" + in.Content}, nil
}
