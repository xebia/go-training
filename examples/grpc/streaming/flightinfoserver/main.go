package main

import (
	"log"

	pb "github.com/xebia/go-training/examples/grpc/streaming/flightinfoapi"
)

func main() {
	s := NewServer()
	err := s.GRPCListenBlocking(pb.DefaultAddressPort)
	if err != nil {
		log.Fatalf("Error starting flight-info server: %s", err)
	}
}
