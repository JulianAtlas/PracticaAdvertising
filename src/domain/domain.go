package domain

import (
	"fmt"
	"sync"

	"net/http"

	"github.com/PracticaAdvertising/src/crossCutting"
)

var currentId int = 0
var mutex = &sync.Mutex{}

type Product struct {
	Id     int
	Nombre string
}

func NewProduct(nombre string) (*Product, *crossCutting.MyError) {
	if nombre == "" {
		return nil, &crossCutting.MyError{Error: fmt.Errorf("El nombre del producto no puede ser vacio"), Status: http.StatusBadRequest}
	}
	mutex.Lock()
	currentId++
	mutex.Unlock()
	return &Product{Nombre: nombre, Id: currentId}, nil
}

func (p *Product) Copy() *Product {
	return &Product{Nombre: p.Nombre, Id: p.Id}
}
