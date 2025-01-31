package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product"
	group "github.com/viictormg/product-api-meli/internal/infra/api/router/group/product"
	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(NewHTTPServer),
		// route
		fx.Provide(NewEchoGroup),
		fx.Provide(group.NewAnomalyInterfaceRoutes),
		fx.Provide(product.NewAnomalyHandler),

		// init functions
		fx.Invoke(func(*echo.Echo) {}),
		fx.Invoke(func(*group.AnomalyInterfaceRoute) {}),
	).Run()
}
