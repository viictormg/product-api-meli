package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/internal/application/usecases"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product"
	group "github.com/viictormg/product-api-meli/internal/infra/api/router/group/product"
	"github.com/viictormg/product-api-meli/internal/infra/clients/db"
	repo "github.com/viictormg/product-api-meli/internal/infra/repository/product"
	repoHistory "github.com/viictormg/product-api-meli/internal/infra/repository/product_history"

	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(NewHTTPServer),
		fx.Provide(db.NewPostgresConnection),
		fx.Provide(repo.NewProductRepository),
		fx.Provide(repoHistory.NewProductHistoryRepository),

		// route
		fx.Provide(NewEchoGroup),
		fx.Provide(group.NewProductInterfaceRoutes),
		fx.Provide(usecases.NewProductUsecase),
		fx.Provide(product.NewProductHandler),

		// init functions
		fx.Invoke(func(*echo.Echo) {}),
		fx.Invoke(func(*group.ProductInterfaceRoute) {}),
	).Run()
}
