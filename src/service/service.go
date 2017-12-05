package service

import (
	"github.com/PracticaAdvertising/src/domain"

	"github.com/PracticaAdvertising/src/crossCutting"
)

type MainController struct {
	Products map[int]*domain.Product
}

func NewMainController() *MainController {
	return &MainController{Products: map[int]*domain.Product{}}
}

// func SearchProduct(c *gin.Contex) {

// }

func (mc *MainController) CreateProduct(productDto *crossCutting.ProductDto) {
	aProduct, _ := domain.NewProduct(productDto.Nombre)
	mc.Products[aProduct.Id] = aProduct
}
