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

func (mc *MainController) CreateProduct(productDto *crossCutting.ProductDto) error {
	aProduct, err := domain.NewProduct(productDto.Nombre)

	if err != nil {
		return err
	}

	mc.Products[aProduct.Id] = aProduct
	return nil
}

func (mc *MainController) ListProducts() map[int]*domain.Product {
	return mc.Products
}
