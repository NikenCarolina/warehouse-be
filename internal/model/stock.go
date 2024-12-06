package model

import "github.com/NikenCarolina/warehouse-be/internal/dto"

type Stock struct {
	WarehouseId   int
	WarehouseName string
	ProductId     int
	ProductName   string
	DusStock      int
	PcsStock      int
}

func (m *Stock) ToDto() *dto.Stock {
	return &dto.Stock{
		Warehouse: dto.Warehouse{
			Id:   m.WarehouseId,
			Name: m.WarehouseName,
		},
		Product: dto.Product{
			Id:   m.ProductId,
			Name: m.ProductName,
		},
		DusStock: m.DusStock,
		PcsStock: m.PcsStock,
	}
}
