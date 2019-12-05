package patientinfoapi

import (
	"fmt"

	"google.golang.org/grpc"
)

func NewGrpcClient(addressPort string) (PatientInfoClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error creating patientinfo-grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewPatientInfoClient(conn), cleanup, nil
}