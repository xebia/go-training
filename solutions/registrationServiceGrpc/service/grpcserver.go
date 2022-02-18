package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/regprotobuf"
)

func StartGrpcServer(port string, service *RegistrationService) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	regprotobuf.RegisterRegistrationServiceServer(grpcServer, service)

	log.Printf("GRPPC server starts listening at port %s...", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
