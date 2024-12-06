package usecase

import (
	"context"

	"github.com/NikenCarolina/warehouse-be/internal/dto"
	"github.com/NikenCarolina/warehouse-be/internal/repository"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context, name string) ([]dto.Product, error)
}

type productUseCase struct {
	store repository.Store
}

func NewProductUseCase(store repository.Store) ProductUseCase {
	return &productUseCase{store: store}
}

func (u *productUseCase) GetProducts(ctx context.Context, name string) ([]dto.Product, error) {
	productRepo := u.store.Product()

	products, err := productRepo.GetAll(ctx, name)
	if err != nil {
		return nil, err
	}

	res := []dto.Product{}
	for _, product := range products {
		res = append(res, *product.ToDto())
	}

	return res, nil
}
