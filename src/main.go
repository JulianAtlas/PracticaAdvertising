package main

import (
	"github.com/PracticaAdvertising/src/rest"
	"github.com/PracticaAdvertising/src/service"
)

func main() {
	mainController := service.NewMainController()
	server := rest.NewServer(mainController)

	r := rest.SetupRouter(server)
	r.Run(":8080")
}
