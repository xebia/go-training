package server

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
)

func (ss *ServerSuite) TestGetDeviceByID() {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", "/device/1234", nil)
	ss.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	w := httptest.NewRecorder()
	ss.server.getDeviceByID(w, req)

	ss.Equal(http.StatusOK, w.Code)

	b, err := io.ReadAll(w.Body)
	ss.NoError(err)

	ss.Equal("Your device ID is: 1234", string(b))

}
