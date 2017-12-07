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
	r.POST("/products", server.CreateProduct)
	r.DELETE("/products/:id", server.DeleteProduct)
	r.PUT("/products", server.UpdateProduct)
	r.GET("/products", server.ListProducts)
	r.GET("/products/:id", server.GetProduct)
	return r
}

func NewServer(aManager *service.MainController) *Server {
	return &Server{ServiceManager: aManager}
}

func (sv *Server) CreateProduct(c *gin.Context) {
	/*params*/
	var err error
	var myErr *crossCutting.MyError

	var productDto crossCutting.ProductDto
	err = c.Bind(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, crossCutting.ApiErr{err.Error(), http.StatusBadRequest})
		return
	}

	productDto.Id, myErr = sv.ServiceManager.CreateProduct(&productDto)

	if myErr != nil {
		c.JSON(myErr.Status, crossCutting.ApiErr{myErr.Error.Error(), myErr.Status})
		return
	}

	c.JSON(http.StatusCreated, productDto)
	return

}

func (sv *Server) ListProducts(c *gin.Context) {

	c.JSON(http.StatusOK, sv.ServiceManager.ListProducts())

}

func (sv *Server) DeleteProduct(c *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, crossCutting.ApiErr{err.Error(), http.StatusBadRequest})
		return
	}
	var myErr *crossCutting.MyError = sv.ServiceManager.DeleteProduct(id)
	if myErr != nil {
		c.JSON(myErr.Status, crossCutting.ApiErr{myErr.Error.Error(), myErr.Status})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (sv *Server) UpdateProduct(c *gin.Context) {
	var err error
	var myErr *crossCutting.MyError

	var productDto crossCutting.ProductDto

	err = c.Bind(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, crossCutting.ApiErr{err.Error(), http.StatusBadRequest})
		return
	}
	productDto.Id, myErr = sv.ServiceManager.UpdateProduct(&productDto)
	if myErr != nil {
		c.JSON(myErr.Status, crossCutting.ApiErr{myErr.Error.Error(), myErr.Status})
		return
	}
	c.JSON(http.StatusOK, productDto)
	return
}

func (sv *Server) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, crossCutting.ApiErr{err.Error(), http.StatusBadRequest})
		return
	}

	productDto, myErr := sv.ServiceManager.GetProductById(id)

	if myErr != nil {
		c.JSON(myErr.Status, crossCutting.ApiErr{myErr.Error.Error(), myErr.Status})
		return
	}
	c.JSON(http.StatusOK, productDto)
	return

}
