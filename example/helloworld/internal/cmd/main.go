package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	impl "learn/api-gateway/example/helloworld/internal/impl/service"
	pb "learn/api-gateway/example/helloworld/service"
	_ "learn/api-gateway/types"
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
	pb.RegisterGreeterServer(s, &impl.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
