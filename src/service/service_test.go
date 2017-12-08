package service_test

import (
	"strconv"
	"testing"

	"github.com/PracticaAdvertising/src/cc"
	"github.com/PracticaAdvertising/src/service"
)

var mc *service.MainController

func setupTest(){
	mc = nil
	mc = service.NewMainController()
}

func TestTheNewProductIsAddedToTheManagerListProduct(t *testing.T) {
	setupTest()
	product, myErr := mc.CreateProduct(&cc.ProductDto{Name: "producto1"})

	if len(mc.Products) != 1 {
		t.Error("La longitud es diferente de un producto")
	}
	if mc.Products[1].Name != product.Name {
		t.Error("El nombre de producto no coincide con el que agregue")
	}
	if myErr != nil {
		t.Error(myErr.Error.Error())
	}
}


func TestListItemsReturnsAllItems(t *testing.T) {
	setupTest()
	var nombreProducto1 string = "celular"
	mc.CreateProduct(&cc.ProductDto{Name: nombreProducto1})
	var nombreProducto2 string = "celular2"
	mc.CreateProduct(&cc.ProductDto{Name: nombreProducto2})
	var nombreProducto3 string = "celular3"
	mc.CreateProduct(&cc.ProductDto{Name: nombreProducto3})

	//operacion
	products := mc.ListProducts()

	//test
	if len(products) == 0 {
		t.Error("el mapa no deberia estar vacio")
	}

	if products[2].Name != nombreProducto2 {
		t.Error("el nombre en el mapa es incorrecto")
	}
	if len(products) != 3 {
		t.Error("la cantidad de productos no es correcta")
	}
}

func TestDeleteProductRemovesIt(t *testing.T) {
	setupTest()
	var nombreProducto1 string = "computadora"
	product, _ := mc.CreateProduct(&cc.ProductDto{Name: nombreProducto1})

	//operacion
	myErr := mc.DeleteProduct(product.Id)

	//test
	if myErr != nil {
		t.Error(myErr.Error.Error())
	}

	if len(mc.Products) != 0 {
		t.Error("el producto no fue borrado correctamente")
	}
}

func TestExpectErrorInInvalidIdWhenDelete(t *testing.T) {
	setupTest()
	var nombreProducto1 string = "billetera"
	product, _ := mc.CreateProduct(&cc.ProductDto{Name: nombreProducto1})

	invalidId := product.Id - 10

	//operacion
	myErr := mc.DeleteProduct(invalidId)

	//test
	if myErr == nil {
		t.Error("deberia haber error al usar id invalido")
	}

	if len(mc.Products) != 1 {
		t.Error("No deberia borrarse el elemento")
	}

	if myErr.Error.Error() != "No existe producto con Id : "+strconv.Itoa(invalidId) {
		t.Error("Mensaje de error no era el esperado")
	}
}

func TestSearchProductById(t *testing.T) {
	setupTest()
	var nombreProducto string = "teclado"
	product, _ := mc.CreateProduct(&cc.ProductDto{Name: nombreProducto})

	productDto, myErr := mc.GetProductById(product.Id)

	if myErr != nil {
		t.Error(myErr.Error.Error())
	}
	if productDto.Name != nombreProducto {
		t.Error("No se encuentra el producto pedido")
	}

}

func TestUpdateProduct(t *testing.T) {
	setupTest()
	var nombreProducto string = "gorra"
	product, _ := mc.CreateProduct(&cc.ProductDto{Name: nombreProducto})

	productDto, _ := mc.GetProductById(product.Id)

	if productDto.Name != nombreProducto {
		t.Error("Se agrego mal el producto")
	}

	var otherNameForProduct string = "Pelota"

	mc.UpdateProduct(&cc.ProductDto{Id: product.Id, Name: otherNameForProduct})

	productDto, _ = mc.GetProductById(product.Id)

	if productDto.Name != otherNameForProduct {
		t.Error("El producto no se updeteo")
	}

}
