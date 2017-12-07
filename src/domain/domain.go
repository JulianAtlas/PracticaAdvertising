package domain

import (
	"fmt"
	"sync"

	"net/http"

	"github.com/PracticaAdvertising/src/cc"
)

var currentId int = 0
var mutex = &sync.Mutex{}

type Product struct {
	Id   int
	Name string
}

func NewProduct(nombre string) (*Product, *cc.MyError) {
	if nombre == "" {
		return nil, &cc.MyError{Error: fmt.Errorf("El nombre del producto no puede ser vacio"), Status: http.StatusBadRequest}
	}
	mutex.Lock()
	currentId++
	mutex.Unlock()
	return &Product{Name: nombre, Id: currentId}, nil
}

func (p *Product) Copy() *Product {
	return &Product{Name: p.Name, Id: p.Id}
}
