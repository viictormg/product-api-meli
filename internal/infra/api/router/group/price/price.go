package group

import (
	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/price"
)

const (
	PricePathV1 = "/price"
)

type PriceInterfaceRoute struct{}

func NewPriceInterfaceRoute(group *echo.Group, handler price.PriceHandlerIF) {
	routes := group.Group(PricePathV1)

	routes.POST("/upload-prices", handler.UploadPriceFile)
}

func NewPriceInterfaceRoutes(group *echo.Group, handler price.PriceHandlerIF) *PriceInterfaceRoute {
	NewPriceInterfaceRoute(group, handler)

	return &PriceInterfaceRoute{}
}
