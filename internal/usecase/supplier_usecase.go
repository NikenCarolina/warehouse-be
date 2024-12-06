package usecase

import (
	"context"

	"github.com/NikenCarolina/warehouse-be/internal/dto"
	"github.com/NikenCarolina/warehouse-be/internal/repository"
)

type SupplierUseCase interface {
	GetSuppliers(ctx context.Context, name string) ([]dto.Supplier, error)
}

type supplierUseCase struct {
	store repository.Store
}

func NewSupplierUseCase(store repository.Store) SupplierUseCase {
	return &supplierUseCase{store: store}
}

func (u *supplierUseCase) GetSuppliers(ctx context.Context, name string) ([]dto.Supplier, error) {
	supplierRepo := u.store.Supplier()

	suppliers, err := supplierRepo.GetAll(ctx, name)
	if err != nil {
		return nil, err
	}

	res := []dto.Supplier{}
	for _, supplier := range suppliers {
		res = append(res, *supplier.ToDto())
	}

	return res, nil
}
