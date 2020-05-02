package main

import (
	"log"
	"net"

	pb "github.com/iandanarko/grpc-sample/proto"
	"google.golang.org/grpc"
)

const (
	port = ":3000"
)

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
