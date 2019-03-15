package main

import (
	"log"
	"net"

	impl "learn/api-gateway/example/echo/internal/impl/service"
	pb "learn/api-gateway/example/echo/service"
	_ "learn/api-gateway/types"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &impl.SimpleEchoServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
