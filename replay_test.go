package plugindemo_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/trondhindenes/traefik-replay"
)

func TestDemo(t *testing.T) {
	cfg := plugindemo.CreateConfig()
	cfg.Headers["X-Host"] = "[[.Host]]"
	cfg.Headers["X-Method"] = "[[.Method]]"
	cfg.Headers["X-URL"] = "[[.URL]]"
	cfg.Headers["X-URL"] = "[[.URL]]"
	cfg.Headers["X-Demo"] = "test"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := plugindemo.New(ctx, next, cfg, "demo-plugin")
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

	assertHeader(t, req, "X-Host", "localhost")
	assertHeader(t, req, "X-URL", "http://localhost")
	assertHeader(t, req, "X-Method", "GET")
	assertHeader(t, req, "X-Demo", "test")
	assertBody(t, req, bodyString)
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}

func assertBody(t *testing.T, req *http.Request, bodyString string) {
	t.Helper()
	body, _ := io.ReadAll(req.Body)
	if string(body) != bodyString {
		t.Errorf("unexpected body")
	}
}
