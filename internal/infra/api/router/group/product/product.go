package group

import (
	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product"
)

const (
	productPathV1 = "/product"
)

type ProductInterfaceRoute struct{}

func NewProductInterfaceRoute(group *echo.Group, handler product.ProductHandlerIF) {
	routes := group.Group(productPathV1)

	routes.PATCH("/:id", handler.UpdatePrice)
}

func NewProductInterfaceRoutes(group *echo.Group, handler product.ProductHandlerIF) *ProductInterfaceRoute {
	NewProductInterfaceRoute(group, handler)

	return &ProductInterfaceRoute{}
}
