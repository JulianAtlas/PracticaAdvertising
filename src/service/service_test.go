package service_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/PracticaAdvertising/src/crossCutting"
	"github.com/PracticaAdvertising/src/service"
)

func TestTheNewProductIsAddedToTheManagerListProduct(t *testing.T) {
	mc := service.NewMainController()
	var nombreProducto string = "botella"
	_, myErr := mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto})

	if len(mc.Products) != 1 {
		t.Error("La longitud es diferente de un producto")
	}
	if mc.Products[1].Nombre != nombreProducto {
		t.Error("El nombre de producto no coincide con el que agregue")
	}
	if myErr != nil {
		t.Error(myErr.Error.Error())
	}
}

//por que este test pasa cuando lo corremos individualmente pero no cuando corremos todo el package?
func TestListItemsReturnsAllItems(t *testing.T) {
	mc := service.NewMainController()
	var nombreProducto1 string = "celular"
	mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto1})
	var nombreProducto2 string = "celular2"
	mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto2})
	var nombreProducto3 string = "celular3"
	mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto3})

	//operacion
	products := mc.ListProducts()

	//test
	if len(products) == 0 {
		t.Error("el mapa no deberia estar vacio")
	}
	fmt.Println(products[2].Nombre)
	if products[2].Nombre != nombreProducto2 {
		t.Error("el nombre en el mapa es incorrecto")
	}
	if len(products) != 3 {
		t.Error("la cantidad de productos no es correcta")
	}
}

func TestDeleteProductRemovesIt(t *testing.T) {
	mc := service.NewMainController()
	var nombreProducto1 string = "computadora"
	id, _ := mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto1})

	//operacion
	myErr := mc.DeleteProduct(id)

	//test
	if myErr != nil {
		t.Error(myErr.Error.Error())
	}

	if len(mc.Products) != 0 {
		t.Error("el producto no fue borrado correctamente")
	}
}

func TestExpectErrorInInvalidIdWhenDelete(t *testing.T) {
	mc := service.NewMainController()
	var nombreProducto1 string = "billetera"
	validId, _ := mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto1})

	invalidId := validId - 10

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
	mc := service.NewMainController()
	var nombreProducto string = "teclado"
	idProducto, _ := mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto})

	productDto, myErr := mc.GetProductById(idProducto)

	if myErr != nil {
		t.Error(myErr.Error.Error())
	}
	if productDto.Nombre != nombreProducto {
		t.Error("No se encuentra el producto pedido")
	}

}

func TestUpdateProduct(t *testing.T) {
	mc := service.NewMainController()
	var nombreProducto string = "gorra"
	idProducto, _ := mc.CreateProduct(&crossCutting.ProductDto{Nombre: nombreProducto})

	productDto, _ := mc.GetProductById(idProducto)

	if productDto.Nombre != nombreProducto {
		t.Error("Se agrego mal el producto")
	}

	var otherNameForProduct string = "Pelota"

	mc.UpdateProduct(&crossCutting.ProductDto{Id: idProducto, Nombre: otherNameForProduct})

	productDto, _ = mc.GetProductById(idProducto)

	if productDto.Nombre != otherNameForProduct {
		t.Error("El producto no se updeteo")
	}

}
