package price

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/internal/application/price/usecases"
)

type PriceHandlerIF interface {
	UploadPriceFile(c echo.Context) error
}

func NewPriceHandler(usecase usecases.PriceUsecaseIF) PriceHandlerIF {
	return &PriceHandler{
		usecase: usecase,
	}
}

type PriceHandler struct {
	usecase usecases.PriceUsecaseIF
}

func (h *PriceHandler) UploadPriceFile(c echo.Context) error {
	file, err := c.FormFile("file")

	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.usecase.UploadPriceFile(file)

	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, nil)
}
