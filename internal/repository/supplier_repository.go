package repository

import (
	"context"

	"github.com/NikenCarolina/warehouse-be/internal/apperror"
	"github.com/NikenCarolina/warehouse-be/internal/model"
)

type SupplierRepository interface {
	GetAll(ctx context.Context, name string) ([]model.Supplier, error)
}

type supplierRepository struct {
	db database
}

func NewSupplierRepository(db database) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) GetAll(ctx context.Context, name string) ([]model.Supplier, error) {
	query := `
		SELECT
			"SupplierPK",
			"SupplierName"
		FROM
			"MasterSupplier"
		WHERE
			"SupplierName" ILIKE '%' || $1 || '%'
	`
	rows, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()
	suppliers := []model.Supplier{}
	for rows.Next() {
		var supplier model.Supplier
		if err := rows.Scan(
			&supplier.Id,
			&supplier.Name,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		suppliers = append(suppliers, supplier)
	}
	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}
	return suppliers, nil

}
