package main

import (
	"net"

	pb "productinfo/service/ecommerce"

	"productinfo/service/internal/app"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// tcp listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start listener: %s", err)
	}

	// gRPC server
	s := grpc.NewServer()
	// register implemented service
	pb.RegisterProductInfoServer(s, &app.Server{})

	// start listener
	log.Infof("Starting gRPC listener on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
