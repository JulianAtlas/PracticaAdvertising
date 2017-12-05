package rest

import (
	"github.com/PracticaAdvertising/src/crossCutting"
	"github.com/PracticaAdvertising/src/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ServiceManager *service.MainController
}

func SetupRouter(server *Server) *gin.Engine {
	r := gin.Default()
	r.POST("/create", server.CreateProduct)
	return r
}

func NewServer(aManager *service.MainController) *Server {
	return &Server{ServiceManager: aManager}
}

func (sv *Server) CreateProduct(c *gin.Context) {
	/*params*/
	var productDto crossCutting.ProductDto
	c.Bind(&productDto)
	sv.ServiceManager.CreateProduct(&productDto)

}
