package rest

import (
	"net/http"

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
	r.GET("/list", server.ListProducts)
	return r
}

func NewServer(aManager *service.MainController) *Server {
	return &Server{ServiceManager: aManager}
}

func (sv *Server) CreateProduct(c *gin.Context) {
	/*params*/
	var productDto crossCutting.ProductDto
	c.Bind(&productDto)
	err := sv.ServiceManager.CreateProduct(&productDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, crossCutting.ApiErr{err.Error(), http.StatusInternalServerError})
		return
	}
}

func (sv *Server) ListProducts(c *gin.Context) {

	c.JSON(http.StatusOK, sv.ServiceManager.ListProducts())

}
