package model

import "github.com/NikenCarolina/warehouse-be/internal/dto"

type Warehouse struct {
	Id   int
	Name string
}

func (m *Warehouse) ToDto() *dto.Warehouse {
	return &dto.Warehouse{
		Id:   m.Id,
		Name: m.Name,
	}
}
