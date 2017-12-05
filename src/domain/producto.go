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

	return &Product{Nombre: nombre}, nil
}
