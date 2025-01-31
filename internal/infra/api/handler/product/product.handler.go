package product

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type AnomalyHandlerIF interface {
	UpdatePrice(context echo.Context) error
}

type AnomalyHandler struct {
}

func NewAnomalyHandler() AnomalyHandlerIF {
	return &AnomalyHandler{}
}

func (a *AnomalyHandler) UpdatePrice(context echo.Context) error {
	fmt.Println("LLEGA ")
	return nil
}
