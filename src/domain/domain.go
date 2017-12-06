package domain

import (
	"fmt"
)

var currentId int = 0

type Product struct {
	Id     int
	Nombre string
}

func NewProduct(nombre string) (*Product, error) {
	if nombre == "" {
		return nil, fmt.Errorf("El nombre del producto no puede ser vacio")
	}
	currentId++
	return &Product{Nombre: nombre, Id: currentId}, nil
}

func (p *Product) Copy() *Product {
	return &Product{Nombre: p.Nombre, Id: p.Id}
}
