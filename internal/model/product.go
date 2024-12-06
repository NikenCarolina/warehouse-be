package model

import "github.com/NikenCarolina/warehouse-be/internal/dto"

type Product struct {
	Id   int
	Name string
}

func (m *Product) ToDto() *dto.Product {
	return &dto.Product{
		Id:   m.Id,
		Name: m.Name,
	}
}
