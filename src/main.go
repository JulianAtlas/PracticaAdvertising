ackage main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/search", services.CreateProduct)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}