package slogx

import (
	"log/slog"
	"os"
)

func Fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
