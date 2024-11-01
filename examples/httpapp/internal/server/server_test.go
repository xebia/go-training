package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
	server *Server
}

func (ss *ServerSuite) SetupSuite() {
	addr := "localhost:10000"
	ss.server = NewServer(nil)

	go func() {
		ss.server.ListenAndServe(addr)
	}()
	time.Sleep(10 * time.Millisecond)
}

func (ss *ServerSuite) TearDownSuite() {
	ss.server.Shutdown()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
