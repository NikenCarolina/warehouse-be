package handler

import (
	"net/http"

	"github.com/NikenCarolina/warehouse-be/internal/appconst"
	"github.com/NikenCarolina/warehouse-be/internal/dto"
	"github.com/NikenCarolina/warehouse-be/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	useCase usecase.ProductUseCase
}

func NewProductHandler(productUseCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		useCase: productUseCase,
	}
}

func (h *ProductHandler) ListProduct(ctx *gin.Context) {
	name := ctx.Query("name")
	data, err := h.useCase.GetProducts(ctx, name)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: appconst.MsgListProductOk,
		Data:    data,
	})
}
