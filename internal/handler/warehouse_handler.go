package handler

import (
	"net/http"

	"github.com/NikenCarolina/warehouse-be/internal/appconst"
	"github.com/NikenCarolina/warehouse-be/internal/apperror"
	"github.com/NikenCarolina/warehouse-be/internal/dto"
	"github.com/NikenCarolina/warehouse-be/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	useCase usecase.WarehouseUseCase
}

func NewWarehouseHandler(warehouseUseCase usecase.WarehouseUseCase) *WarehouseHandler {
	return &WarehouseHandler{
		useCase: warehouseUseCase,
	}
}

func (h *WarehouseHandler) ListWarehouse(ctx *gin.Context) {
	name := ctx.Query("name")
	data, err := h.useCase.GetWarehouses(ctx, name)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgListWarehouseOk,
		Data:    data,
	})
}

func (h *WarehouseHandler) ItemsIn(ctx *gin.Context) {
	var uri dto.WarehouseUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}
	var req dto.WarehouseProductInReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	err := h.useCase.ItemsIn(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: appconst.MsgItemsInOk,
	})

}

func (h *WarehouseHandler) ItemsOut(ctx *gin.Context) {
	var uri dto.WarehouseUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(apperror.ErrBadRequest)
		return
	}
	var req dto.WarehouseProductInReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	err := h.useCase.ItemsOut(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: appconst.MsgItemsOutOk,
	})

}

func (h *WarehouseHandler) ListReport(ctx *gin.Context) {
	data, err := h.useCase.GetReport(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgListWarehouseOk,
		Data:    data,
	})
}
