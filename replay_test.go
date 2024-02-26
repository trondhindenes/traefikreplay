package traefikreplay_test

import (
	"context"
	"github.com/trondhindenes/traefikreplay"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDemo(t *testing.T) {
	cfg := traefikreplay.CreateConfig()
	cfg.ReplayUrl = "http://localhost:8080"
	cfg.Percentage = 0

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefikreplay.New(ctx, next, cfg, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	bodyString := "Hello, this is the request body."

	// Create an io.Reader from the body string
	bodyReader := strings.NewReader(bodyString)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

}
