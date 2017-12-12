package main

import (
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
)

const ServerPath string = "http://localhost:8080"



func MakeRequestToRouter(router *gin.Engine, req *http.Request ) *testResponse{
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	var b []byte

	if resp.Body != nil {
		b = resp.Body.Bytes()
	} else {
		b = []byte("API did not respond to the call") //Not the same as an empty response
	}

	return &testResponse{HttpStatusCode: resp.Code, Body:string(b), Resp: resp  }

}


type testResponse struct {
	HttpStatusCode int
	Body           string
	Resp  *httptest.ResponseRecorder
}
