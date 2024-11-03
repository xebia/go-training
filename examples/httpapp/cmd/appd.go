package main

import (
	"fmt"

	"github.com/xebia/go-training/examples/httpapp/internal/server"
)

func main() {
	dbService := interface{}(nil)

	srv := server.NewServer(dbService)
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "8090")
	srv.ListenAndServe(addr)
}
