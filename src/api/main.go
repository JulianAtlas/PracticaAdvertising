package main

import (
	"net/http"

	"github.com/PracticaAdvertising/src/api/rest"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	StartApp()
}

func StartApp() {
	r := rest.SetupRouter()
	r.Run(":8080")
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
