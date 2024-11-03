package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	ctx := context.Background()

	s := Server{router: http.NewServeMux()}

	cpr := CreatePatientRequest{Patient{
		FullName:    "Patrick",
		AddressLine: "a",
		Allergies:   []string{"cilantro"},
	}}

	b, err := json.Marshal(cpr)

	assert.NoError(t, err)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/patient", bytes.NewBuffer(b))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	s.createPatientHandler().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	resp := CreatePatientResponse{}
	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	err = json.Unmarshal(body, &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Patrick", resp.FullName)
	assert.Equal(t, "a", resp.AddressLine)
	assert.Len(t, resp.Allergies, 1)
	assert.Equal(t, "cilantro", resp.Allergies[0])

}
