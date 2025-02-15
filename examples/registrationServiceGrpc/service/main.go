package main

import (
	"log"

	"github.com/xebia/go-training/examples/registrationServiceGrpc/lib/impl/datastoring"
	"github.com/xebia/go-training/examples/registrationServiceGrpc/lib/impl/emailsending"
	"github.com/xebia/go-training/examples/registrationServiceGrpc/lib/impl/pincoding"
	"github.com/xebia/go-training/examples/registrationServiceGrpc/lib/impl/uuiding"
	"github.com/xebia/go-training/examples/registrationServiceGrpc/regprotobuf"
)

func main() {
	uuidGenerator := uuiding.New()
	patientStore := datastoring.New()
	pincoder := pincoding.New()
	emailSender := emailsending.New()

	service := NewRegistrationService(uuidGenerator, patientStore, pincoder, emailSender)
	err := regprotobuf.StartGrpcServer(regprotobuf.DefaultPort, service)
	if err != nil {
		log.Fatalf("Error starting registration server: %s", err)
	}
}
