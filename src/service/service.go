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


//Es mejor devolver una lista de copias a todos los productos? O esta bien devolver los punteros originales?
func (mc *MainController) ListProducts() map[int]*domain.Product {
	return *mc.copiyMapOfProducts()
}


func (mc *MainController) copiyMapOfProducts() *map[int]*domain.Product{
	var res map[int]*domain.Product = map[int]*domain.Product{}
	
	for k,v := range mc.Products {	
		res[k] = v.Copy()
	}
	return &res
}

