// Package plugindemo a demo plugin.
package traefikreplay

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Config the plugin configuration.
type Config struct {
	Percentage int    `json:"percentage,omitempty"`
	ReplayUrl  string `json:"replayUrl,omitempty"`
	OnlyIfJson bool   `json:"onlyIfJson,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Demo a Demo plugin.
type Demo struct {
	next       http.Handler
	percentage int
	replayUrl  string
	onlyIfJson bool
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.ReplayUrl == "" {
		return nil, fmt.Errorf("ReplayUrl cannot be empty")
	}

	return &Demo{
		next:       next,
		percentage: config.Percentage,
		replayUrl:  config.ReplayUrl,
		onlyIfJson: config.OnlyIfJson,
	}, nil
}

func shouldReplay(percentage int) bool {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	formatMessage := "random number: %d, percentage: %d, should replay: %t \n"
	if randomNumber < percentage {
		os.Stdout.WriteString(fmt.Sprintf(formatMessage, randomNumber, percentage, true))
		return true
	}
	os.Stdout.WriteString(fmt.Sprintf(formatMessage, randomNumber, percentage, false))
	return false
}

func isJsonRequest(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")

	if contentType == "application/json" {
		os.Stdout.WriteString("request is json \n")
		return true
	}
	os.Stdout.WriteString("request is not json \n")
	return false
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	doReplay := true
	if !(a.onlyIfJson && isJsonRequest(req)) {
		doReplay = false
	}
	if !(shouldReplay(a.percentage)) {
		doReplay = false
	}

	if doReplay == true {
		body, err := io.ReadAll(req.Body)
		//defer req.Body.Close()
		req.Body = io.NopCloser(bytes.NewBuffer(body))
		newRequest, err := http.NewRequest(req.Method, a.replayUrl, bytes.NewBuffer(body))
		for key, value := range req.Header {
			newRequest.Header.Set(key, value[0])
		}
		newRequest.Header.Set("X-Original-Url", req.URL.RequestURI())
		client := http.Client{}
		resp, err := client.Do(newRequest)
		if err != nil {

		}
		defer resp.Body.Close()
	}
	a.next.ServeHTTP(rw, req)
}
