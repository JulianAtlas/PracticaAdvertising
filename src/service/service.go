package service

import (
	"fmt"
	"strconv"

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

func (mc *MainController) CreateProduct(productDto *crossCutting.ProductDto) (int, error) {
	aProduct, err := domain.NewProduct(productDto.Nombre)

	if err != nil {
		return 0, err
	}

	mc.Products[aProduct.Id] = aProduct
	return aProduct.Id, nil
}

//Es mejor devolver una lista de copias a todos los productos? O esta bien devolver los punteros originales?
func (mc *MainController) ListProducts() map[int]*domain.Product {
	return *mc.copyMapOfProducts()
}

//Es mejor devolver una lista de copias a todos los productos? O esta bien devolver los punteros originales?
func (mc *MainController) DeleteProduct(id int) error {
	_, err := searchProductById(mc, id)

	if err != nil {
		return err
	}

	delete(mc.Products, id)
	return nil
}

func (mc *MainController) UpdateProduct(productDto *crossCutting.ProductDto) (int, error) {
	aProduct, err := searchProductById(mc, productDto.Id)

	if err != nil {
		return 0, err
	}

	aProduct.Nombre = productDto.Nombre
	return aProduct.Id, nil
}

func (mc *MainController) SearchProduct(id int) (*crossCutting.ProductDto, error) {
	aProduct, err := searchProductById(mc, id)

	if err != nil {
		return nil, err
	}

	return &crossCutting.ProductDto{Id: aProduct.Id, Nombre: aProduct.Nombre}, nil
}

func (mc *MainController) copyMapOfProducts() *map[int]*domain.Product {
	var res map[int]*domain.Product = map[int]*domain.Product{}

	for k, v := range mc.Products {
		res[k] = v.Copy()
	}
	return &res
}

func searchProductById(mc *MainController, id int) (*domain.Product, error) {
	product, existeProducto := mc.Products[id]
	if !existeProducto {
		return nil, fmt.Errorf("No existe producto con Id : " + strconv.Itoa(id))
	}
	return product, nil
}
