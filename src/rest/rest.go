package rest

import (
	"net/http"
	"strconv"

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
	r.POST("/delete", server.DeleteProduct)
	r.POST("/update", server.UpdateProduct)
	r.GET("/list", server.ListProducts)
	r.GET("/search/:id", server.SearchProduct)
	return r
}

func NewServer(aManager *service.MainController) *Server {
	return &Server{ServiceManager: aManager}
}

func (sv *Server) CreateProduct(c *gin.Context) {
	/*params*/
	var err error

	var productDto crossCutting.ProductDto
	c.Bind(&productDto)
	productDto.Id, err = sv.ServiceManager.CreateProduct(&productDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, crossCutting.ApiErr{err.Error(), http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, productDto)
	return

}

func (sv *Server) ListProducts(c *gin.Context) {

	c.JSON(http.StatusOK, sv.ServiceManager.ListProducts())

}

func (sv *Server) DeleteProduct(c *gin.Context) {
	var productDto crossCutting.ProductDto
	c.Bind(&productDto)
	c.JSON(http.StatusOK, sv.ServiceManager.DeleteProduct(productDto.Id))
}

func (sv *Server) UpdateProduct(c *gin.Context) {
	var err error

	var productDto crossCutting.ProductDto
	c.Bind(&productDto)
	productDto.Id, err = sv.ServiceManager.UpdateProduct(&productDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, crossCutting.ApiErr{err.Error(), http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, productDto)
	return
}

func (sv *Server) SearchProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	productDto, err := sv.ServiceManager.SearchProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, crossCutting.ApiErr{err.Error(), http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, productDto)
	return

}
