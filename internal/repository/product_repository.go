package repository

import (
	"context"

	"github.com/NikenCarolina/warehouse-be/internal/apperror"
	"github.com/NikenCarolina/warehouse-be/internal/model"
)

type ProductRepository interface {
	GetAll(ctx context.Context, name string) ([]model.Product, error)
}

type productRepository struct {
	db database
}

func NewProductRepository(db database) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll(ctx context.Context, name string) ([]model.Product, error) {
	query := `
		SELECT
			"ProductPK",
			"ProductName"
		FROM
			"MasterProduct"
		WHERE
			"ProductName" ILIKE '%' || $1 || '%'
	`
	rows, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()
	products := []model.Product{}
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}
	return products, nil

}
