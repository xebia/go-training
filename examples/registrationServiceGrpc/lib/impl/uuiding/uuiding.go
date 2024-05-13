package uuiding

import (
	"github.com/google/uuid"
	"github.com/xebia/go-training/examples/registrationServiceGrpc/lib/api/uuider"
)

type basicuuider struct{}

func New() uuider.UuidGenerator {
	return &basicuuider{}
}

func (u basicuuider) GenerateUuid() string {
	return uuid.NewString()
}
