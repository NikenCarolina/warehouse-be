package usecase

import (
	"context"

	"github.com/NikenCarolina/warehouse-be/internal/apperror"
	"github.com/NikenCarolina/warehouse-be/internal/dto"
	"github.com/NikenCarolina/warehouse-be/internal/repository"
)

type WarehouseUseCase interface {
	GetWarehouses(ctx context.Context, name string) ([]dto.Warehouse, error)
	ItemsIn(ctx context.Context, req dto.WarehouseProductInReq) error
	ItemsOut(ctx context.Context, req dto.WarehouseProductInReq) error
	GetReport(ctx context.Context) ([]dto.Stock, error)
}

type warehouseUseCase struct {
	store repository.Store
}

func NewWarehouseUseCase(store repository.Store) WarehouseUseCase {
	return &warehouseUseCase{store: store}
}

func (u *warehouseUseCase) GetWarehouses(ctx context.Context, name string) ([]dto.Warehouse, error) {
	warehouseRepo := u.store.Warehouse()

	warehouses, err := warehouseRepo.GetAll(ctx, name)
	if err != nil {
		return nil, err
	}

	res := []dto.Warehouse{}
	for _, warehouse := range warehouses {
		res = append(res, *warehouse.ToDto())
	}

	return res, nil
}

func (u *warehouseUseCase) ItemsIn(ctx context.Context, req dto.WarehouseProductInReq) error {
	_, err := u.store.Atomic(ctx, func(s repository.Store) (any, error) {
		warehouseRepo := s.Warehouse()
		headerId, err := warehouseRepo.ItemInHeader(ctx, req.Warehouse.Id, req.Supplier.Id, req.Note)
		if err != nil {
			return nil, err
		}
		for _, value := range req.Items {
			err := warehouseRepo.ItemInDetail(ctx, *headerId, value.Id, value.Dus, value.Pcs)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *warehouseUseCase) GetReport(ctx context.Context) ([]dto.Stock, error) {
	warehouseRepo := u.store.Warehouse()
	stocks, err := warehouseRepo.GetStockReport(ctx)
	if err != nil {
		return nil, err
	}

	res := []dto.Stock{}
	for _, stock := range stocks {
		res = append(res, *stock.ToDto())
	}

	return res, nil
}

func (u *warehouseUseCase) ItemsOut(ctx context.Context, req dto.WarehouseProductInReq) error {
	_, err := u.store.Atomic(ctx, func(s repository.Store) (any, error) {
		warehouseRepo := s.Warehouse()
		headerId, err := warehouseRepo.ItemOutHeader(ctx, req.Warehouse.Id, req.Supplier.Id, req.Note)
		if err != nil {
			return nil, err
		}
		for _, value := range req.Items {
			dus, pcs, err := warehouseRepo.CheckStock(ctx, req.Warehouse.Id, value.Id)
			if err != nil {
				return nil, err
			}
			if value.Dus > *dus || value.Pcs > *pcs {
				return nil, apperror.ErrInvalidAmount
			}
			err = warehouseRepo.ItemOutDetail(ctx, *headerId, value.Id, value.Dus, value.Pcs)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
