package handler

import (
	"database/sql"

	"github.com/NikenCarolina/warehouse-be/internal/config"
	"github.com/NikenCarolina/warehouse-be/internal/repository"
	"github.com/NikenCarolina/warehouse-be/internal/usecase"
)

type HandlerOpts struct {
	*WarehouseHandler
	*SupplierHandler
	*ProductHandler
}

func Init(db *sql.DB, config *config.Config) *HandlerOpts {
	store := repository.NewStore(db)
	warehouseUseCase := usecase.NewWarehouseUseCase(store)
	supplierUseCase := usecase.NewSupplierUseCase(store)
	productUseCase := usecase.NewProductUseCase(store)
	return &HandlerOpts{
		WarehouseHandler: NewWarehouseHandler(warehouseUseCase),
		SupplierHandler:  NewSupplierHandler(supplierUseCase),
		ProductHandler:   NewProductHandler(productUseCase),
	}
}
