package model

import "github.com/NikenCarolina/warehouse-be/internal/dto"

type Supplier struct {
	Id   int
	Name string
}

func (m *Supplier) ToDto() *dto.Supplier {
	return &dto.Supplier{
		Id:   m.Id,
		Name: m.Name,
	}
}
