package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
	productUsecase "github.com/viictormg/product-api-meli/internal/application/usecases"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product/dto"
)

type ProductHandlerIF interface {
	UpdatePrice(context echo.Context) error
}

type ProductHandler struct {
	usecase productUsecase.ProductUsecaseIF
}

func NewProductHandler(usecase productUsecase.ProductUsecaseIF) ProductHandlerIF {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (a *ProductHandler) UpdatePrice(context echo.Context) error {
	var request dto.UpdatePriceRequest

	productId := context.Param("id")
	if err := context.Bind(&request); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	ctx := context.Request().Context()

	request.ProductID = productId

	err := a.usecase.UpdatePrice(ctx, request)

	return context.JSON(http.StatusOK, err)

}
