package repository

import (
	"context"
	"log"

	"github.com/NikenCarolina/warehouse-be/internal/apperror"
	"github.com/NikenCarolina/warehouse-be/internal/model"
)

type WarehouseRepository interface {
	GetAll(ctx context.Context, name string) ([]model.Warehouse, error)
	ItemInHeader(ctx context.Context, whsId, suppId int, note string) (*int, error)
	ItemInDetail(ctx context.Context, headerId, productId, dus, pcs int) error
	ItemOutHeader(ctx context.Context, whsId, suppId int, note string) (*int, error)
	ItemOutDetail(ctx context.Context, headerId, productId, dus, pcs int) error
	GetStockReport(ctx context.Context) ([]model.Stock, error)
	CheckStock(ctx context.Context, warehouseId, productId int) (*int, *int, error)
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

func (r *warehouseRepository) ItemOutHeader(ctx context.Context, whsId, suppId int, note string) (*int, error) {
	query := `
		INSERT INTO "TransaksiPengeluaranBarangHeader" (
			"TrxOutNo", "WhsIdf", "TrxOutDate", "TrxOutSuppIdf", "TrxOutNotes"
		)
		VALUES (gen_random_uuid(), $1, NOW(), $2, $3)
		RETURNING "TrxOutPK"
	`
	var headerId int
	if err := r.db.QueryRowContext(ctx, query, whsId, suppId, note).Scan(&headerId); err != nil {
		log.Println(headerId)
		return nil, err
	}

	return &headerId, nil

}

func (r *warehouseRepository) ItemOutDetail(ctx context.Context, headerId, productId, dus, pcs int) error {
	query := `
		INSERT INTO "TransaksiPengeluaranBarangDetail" (
			"TrxOutIDF", "TrxOutDProductIdf", "TrxOutDQtyDus", "TrxOutDQtyPcs"
		)
		VALUES
			($1, $2, $3, $4)
	`
	if _, err := r.db.ExecContext(ctx, query, headerId, productId, dus, pcs); err != nil {
		return err
	}

	return nil

}

func (r *warehouseRepository) CheckStock(ctx context.Context, warehouseId, productId int) (*int, *int, error) {
	query := `
    WITH ReceivedStock AS (
    SELECT
        h."WhsIdf" AS "WarehouseId",
        d."TrxInDProductIdf" AS "ProductId",
        SUM(d."TrxInDQtyDus") AS "TotalDusReceived",
        SUM(d."TrxInDQtyPcs") AS "TotalPcsReceived"
    FROM "TransaksiPenerimaanBarangDetail" d
    JOIN "TransaksiPenerimaanBarangHeader" h ON d."TrxInIDF" = h."TrxInPK"
    GROUP BY h."WhsIdf", d."TrxInDProductIdf"
	),
	IssuedStock AS (
		SELECT
			h."WhsIdf" AS "WarehouseId",
			d."TrxOutDProductIdf" AS "ProductId",
			SUM(d."TrxOutDQtyDus") AS "TotalDusIssued",
			SUM(d."TrxOutDQtyPcs") AS "TotalPcsIssued"
		FROM "TransaksiPengeluaranBarangDetail" d
		JOIN "TransaksiPengeluaranBarangHeader" h ON d."TrxOutIDF" = h."TrxOutPK"
		GROUP BY h."WhsIdf", d."TrxOutDProductIdf"
	)
	SELECT
			COALESCE(r."TotalDusReceived", 0) - COALESCE(i."TotalDusIssued", 0) AS "AvailableDus",
			COALESCE(r."TotalPcsReceived", 0) - COALESCE(i."TotalPcsIssued", 0) AS "AvailablePcs"
		FROM "MasterProduct" p
		LEFT JOIN ReceivedStock r ON p."ProductPK" = r."ProductId" AND r."WarehouseId" = $1 
		LEFT JOIN IssuedStock i ON p."ProductPK" = i."ProductId" AND i."WarehouseId" = $1 
		WHERE p."ProductPK" = $2;
	`
	var dusStock int
	var pcsStock int
	if err := r.db.QueryRowContext(ctx, query, warehouseId, productId).Scan(&dusStock, &pcsStock); err != nil {
		return nil, nil, err
	}
	return &dusStock, &pcsStock, nil
}

func (r *warehouseRepository) GetStockReport(ctx context.Context) ([]model.Stock, error) {
	query := ` 
		WITH ReceivedStock AS (
			SELECT
				h."WhsIdf" AS "WarehouseId",
				d."TrxInDProductIdf" AS "ProductId",
				SUM(d."TrxInDQtyDus") AS "TotalDusReceived",
				SUM(d."TrxInDQtyPcs") AS "TotalPcsReceived"
			FROM "TransaksiPenerimaanBarangDetail" d
			JOIN "TransaksiPenerimaanBarangHeader" h ON d."TrxInIDF" = h."TrxInPK"
			GROUP BY h."WhsIdf", d."TrxInDProductIdf"
		),
		IssuedStock AS (
			SELECT
				h."WhsIdf" AS "WarehouseId",
				d."TrxOutDProductIdf" AS "ProductId",
				SUM(d."TrxOutDQtyDus") AS "TotalDusIssued",
				SUM(d."TrxOutDQtyPcs") AS "TotalPcsIssued"
			FROM "TransaksiPengeluaranBarangDetail" d
			JOIN "TransaksiPengeluaranBarangHeader" h ON d."TrxOutIDF" = h."TrxOutPK"
			GROUP BY h."WhsIdf", d."TrxOutDProductIdf"
		)
		SELECT
			w."WhsPK" AS "WarehouseId",
			w."WhsName",
			p."ProductPK" AS "ProductId",
			p."ProductName",
			COALESCE(r."TotalDusReceived", 0) - COALESCE(i."TotalDusIssued", 0) AS "StockDus",
			COALESCE(r."TotalPcsReceived", 0) - COALESCE(i."TotalPcsIssued", 0) AS "StockPcs"
		FROM "MasterWarehouse" w
		CROSS JOIN "MasterProduct" p
		LEFT JOIN ReceivedStock r ON w."WhsPK" = r."WarehouseId" AND p."ProductPK" = r."ProductId"
		LEFT JOIN IssuedStock i ON w."WhsPK" = i."WarehouseId" AND p."ProductPK" = i."ProductId"
		ORDER BY w."WhsName", p."ProductName"`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()
	stocks := []model.Stock{}
	for rows.Next() {
		var stock model.Stock
		if err := rows.Scan(
			&stock.WarehouseId,
			&stock.WarehouseName,
			&stock.ProductId,
			&stock.ProductName,
			&stock.DusStock,
			&stock.PcsStock,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		stocks = append(stocks, stock)
	}
	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}
	return stocks, nil
}
