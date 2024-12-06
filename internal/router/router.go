package router

import (
	"net/http"

	"github.com/NikenCarolina/warehouse-be/internal/config"
	"github.com/NikenCarolina/warehouse-be/internal/handler"
	"github.com/NikenCarolina/warehouse-be/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(opts *handler.HandlerOpts, config *config.Config) http.Handler {
	r := gin.Default()
	r.ContextWithFallback = true

	middlewares := []gin.HandlerFunc{
		cors.New(*config.Cors),
		middleware.Error(),
	}
	r.Use(middlewares...)
	r.GET("/warehouses", opts.WarehouseHandler.ListWarehouse)
	r.GET("/suppliers", opts.SupplierHandler.ListSupplier)
	r.GET("/products", opts.ProductHandler.ListProduct)
	r.POST("/warehouses/:warehouse_id/items/in", opts.WarehouseHandler.ItemsIn)
	r.POST("/warehouses/:warehouse_id/items/out", opts.WarehouseHandler.ItemsOut)
	r.GET("/reports", opts.WarehouseHandler.ListReport)

	return r
}
