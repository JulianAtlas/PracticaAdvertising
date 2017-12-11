package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
	)

const ServerPath string = "http://localhost:8080"

func TestMain(m *testing.M) {
	Setup()

	code := m.Run()

	os.Exit(code)
}

func Setup() {
	//main.StartApp()

	time.Sleep(500 * time.Millisecond)
}

func DoLocalRequest(method string, url string, headers http.Header, body string) *testResponse {
	var reader io.Reader = nil

	if body != "" {
		reader = strings.NewReader(body)
	}

	request, _ := http.NewRequest(method, url, reader)

	request.Header = headers

	response := httptest.NewRecorder()

	//main.ServeHTTP(response, request)

	var b []byte

	if response.Body != nil {
		b = response.Body.Bytes()
	} else {
		b = []byte("API did not respond to the call") //Not the same as an empty response
	}

	return &testResponse{HttpStatusCode: response.Code, Headers: response.Header(), Body: b}
}

type testResponse struct {
	HttpStatusCode int
	Headers        http.Header
	Body           []byte
}
