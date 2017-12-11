package domain

import (
	"fmt"
	"net/http"

	"github.com/PracticaAdvertising/src/api/cc"
)

type Product struct {
	Id   int
	Name string
}

func NewProduct(name string) (*Product, *cc.MyError) {
	if name == "" {
		return nil, &cc.MyError{Error: fmt.Errorf("El nombre del producto no puede ser vacio"), Status: http.StatusBadRequest}
	}
	return &Product{Name: name}, nil
}

func (p *Product) Copy() *Product {
	return &Product{Name: p.Name, Id: p.Id}
}
