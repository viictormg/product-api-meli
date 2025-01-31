package group

import (
	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product"
)

const (
	anomalyPathV1 = "/anomaly"
)

type AnomalyInterfaceRoute struct{}

func NewAnomalyInterfaceRoute(group *echo.Group, handler product.AnomalyHandlerIF) {
	routes := group.Group(anomalyPathV1)

	routes.POST("", handler.UpdatePrice)
}

func NewAnomalyInterfaceRoutes(group *echo.Group, handler product.AnomalyHandlerIF) *AnomalyInterfaceRoute {
	NewAnomalyInterfaceRoute(group, handler)

	return &AnomalyInterfaceRoute{}
}
