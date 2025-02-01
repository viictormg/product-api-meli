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

func (p *ProductHandler) UpdatePrice(context echo.Context) error {
	var request dto.UpdatePriceRequest

	if err := context.Bind(&request); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	err := p.usecase.UpdatePrice(context.Request().Context(), request)

	var metadata interface{} = nil
	if err != nil {
		metadata = map[string]interface{}{
			"error": err.Error(),
		}
	}

	return context.JSON(http.StatusOK, dto.ProductResponseDTO{
		ItemID:   request.ProductID,
		Price:    request.Price,
		Anomaly:  err != nil,
		Metadata: metadata,
	})

}
