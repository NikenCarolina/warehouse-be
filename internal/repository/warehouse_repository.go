package repository

import (
	"context"

	"github.com/NikenCarolina/warehouse-be/internal/apperror"
	"github.com/NikenCarolina/warehouse-be/internal/model"
)

type WarehouseRepository interface {
	GetAll(ctx context.Context, name string) ([]model.Warehouse, error)
	ItemInHeader(ctx context.Context, whsId, suppId int, note string) (*int, error)
	ItemInDetail(ctx context.Context, headerId, productId, dus, pcs int) error
}

type warehouseRepository struct {
	db database
}

func NewWarehouseRepository(db database) WarehouseRepository {
	return &warehouseRepository{db}
}

func (r *warehouseRepository) GetAll(ctx context.Context, name string) ([]model.Warehouse, error) {
	query := `
		SELECT
			"WhsPK",
			"WhsName"
		FROM
			"MasterWarehouse"
		WHERE
			"WhsName" ILIKE '%' || $1 || '%'
	`
	rows, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()
	warehouses := []model.Warehouse{}
	for rows.Next() {
		var warehouse model.Warehouse
		if err := rows.Scan(
			&warehouse.Id,
			&warehouse.Name,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		warehouses = append(warehouses, warehouse)
	}
	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}
	return warehouses, nil

}

func (r *warehouseRepository) ItemInHeader(ctx context.Context, whsId, suppId int, note string) (*int, error) {
	query := `
		INSERT INTO "TransaksiPenerimaanBarangHeader" (
			"TrxInNo", "WhsIdf", "TrxInDate", "TrxInSuppIdf", "TrxInNotes"
		)
		VALUES (gen_random_uuid(), $1, NOW(), $2, $3)
		RETURNING "TrxInPK"
	`
	var headerId int
	if err := r.db.QueryRowContext(ctx, query, whsId, suppId, note).Scan(&headerId); err != nil {
		return nil, err
	}

	return &headerId, nil

}

func (r *warehouseRepository) ItemInDetail(ctx context.Context, headerId, productId, dus, pcs int) error {
	query := `
		INSERT INTO "TransaksiPenerimaanBarangDetail" (
			"TrxInIDF", "TrxInDProductIdf", "TrxInDQtyDus", "TrxInDQtyPcs"
		)
		VALUES
			($1, $2, $3, $4)
	`
	if _, err := r.db.ExecContext(ctx, query, headerId, productId, dus, pcs); err != nil {
		return err
	}

	return nil

}
