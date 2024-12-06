package handler

import (
	"net/http"

	"github.com/NikenCarolina/warehouse-be/internal/appconst"
	"github.com/NikenCarolina/warehouse-be/internal/dto"
	"github.com/NikenCarolina/warehouse-be/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	useCase usecase.SupplierUseCase
}

func NewSupplierHandler(warehouseUseCase usecase.SupplierUseCase) *SupplierHandler {
	return &SupplierHandler{
		useCase: warehouseUseCase,
	}
}

func (h *SupplierHandler) ListSupplier(ctx *gin.Context) {
	name := ctx.Query("name")
	data, err := h.useCase.GetSuppliers(ctx, name)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgListSupplierOk,
		Data:    data,
	})
}
