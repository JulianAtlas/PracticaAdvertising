package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PracticaAdvertising/src/domain"

	"github.com/PracticaAdvertising/src/cc"
)

type MainController struct {
	Products map[int]*domain.Product
}

func NewMainController() *MainController {
	return &MainController{Products: map[int]*domain.Product{}}
}

// func SearchProduct(c *gin.Contex) {

// }

func (mc *MainController) CreateProduct(productDto *cc.ProductDto) (*cc.ProductDto, *cc.MyError) {
	aProduct, err := domain.NewProduct(productDto.Name)

	if err != nil {
		return nil, err
	}

	mc.Products[aProduct.Id] = aProduct

	return toDto(aProduct), nil
}

//Es mejor devolver una lista de copias a todos los productos? O esta bien devolver los punteros originales?
func (mc *MainController) ListProducts() map[int]*domain.Product {
	return *mc.copyMapOfProducts()
}

//Es mejor devolver una lista de copias a todos los productos? O esta bien devolver los punteros originales?
func (mc *MainController) DeleteProduct(id int) *cc.MyError {
	_, err := getProductById(mc, id)

	if err != nil {
		return err
	}

	delete(mc.Products, id)
	return nil
}

func (mc *MainController) UpdateProduct(productDto *cc.ProductDto) (int, *cc.MyError) {
	aProduct, myError := getProductById(mc, productDto.Id)

	if myError != nil {
		return 0, myError
	}

	aProduct.Name = productDto.Name
	return aProduct.Id, nil
}

func (mc *MainController) GetProductById(id int) (*cc.ProductDto, *cc.MyError) {
	aProduct, myErr := getProductById(mc, id)

	if myErr != nil {
		return nil, myErr
	}

	return &cc.ProductDto{Id: aProduct.Id, Name: aProduct.Name}, nil
}

func (mc *MainController) copyMapOfProducts() *map[int]*domain.Product {
	var res map[int]*domain.Product = map[int]*domain.Product{}

	for k, v := range mc.Products {
		res[k] = v.Copy()
	}
	return &res
}

func (mc *MainController) mapOfProductsToDtoProducts() *map[int]*cc.ProductDto {
	var res map[int]*cc.ProductDto = map[int]*cc.ProductDto{}

	for k, v := range mc.Products {
		res[k] = toDto(v)
	}
	return &res
}

func getProductById(mc *MainController, id int) (*domain.Product, *cc.MyError) {
	product, existeProducto := mc.Products[id]
	if !existeProducto {
		return nil, &cc.MyError{fmt.Errorf("No existe producto con Id : " + strconv.Itoa(id)), http.StatusNotFound}
	}
	return product, nil
}

func toDto(p *domain.Product) *cc.ProductDto {
	return &cc.ProductDto{Id: p.Id, Name: p.Name}
}
