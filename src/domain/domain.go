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

func NewProduct(name string) (*Product, *cc.MyError) {
	if name == "" {
		return nil, &cc.MyError{Error: fmt.Errorf("El nombre del producto no puede ser vacio"), Status: http.StatusBadRequest}
	}
	mutex.Lock()
	currentId++
	mutex.Unlock()
	return &Product{Name: name, Id: currentId}, nil
}

func (p *Product) Copy() *Product {
	return &Product{Name: p.Name, Id: p.Id}
}

