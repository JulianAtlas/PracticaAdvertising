package main

import (
	"net/http"

	"github.com/PracticaAdvertising/src/api/rest"
	"github.com/PracticaAdvertising/src/api/service"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	StartApp()
}

func StartApp() {
	mainController := service.NewMainController()
	server := rest.NewServer(mainController)

	r := rest.SetupRouter(server)
	r.Run(":8080")
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
