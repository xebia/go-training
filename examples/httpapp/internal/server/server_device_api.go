package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrRequestInvalid = errors.New("request invalid")

func (s *Server) getDeviceByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		writeError(w, fmt.Sprintf("invalid Path: %s", ErrRequestInvalid), http.StatusBadRequest)
		return
	}

	id := parts[2]
	if id == "" {
		writeError(w, fmt.Sprintf("invalid device ID: %s", ErrRequestInvalid), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(w, "Your device ID is: %v", id)
	if err != nil {
		writeError(w, fmt.Sprintf("could not write response header: %s", err), http.StatusInternalServerError)
	}
}
