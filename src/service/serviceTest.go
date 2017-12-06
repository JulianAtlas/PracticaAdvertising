package service_test

import (
	"testing"

	"github.com/PracticaAdvertising/src/domain"
	"github.com/PracticaAdvertising/src/crossCutting"
)


func TestTheNewProductIsAddedToTheManagerListProduct(t *testing.T){
	mc := NewMainController()
	var nombreProducto string = "celular"
	mc.CreateProducto(&ProductoDto{Nombre: nombreProducto})


	if len(mc.Products)  != 1 {
		t.Error("La longitud es diferente de un producto")
	}
	if mc[1].Nombre != nombreProducto {
		t.Error("El nombre de producto no coincide con el que agregue")
	}

}

