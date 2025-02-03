package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/viictormg/product-api-meli/config"
	usecasePrice "github.com/viictormg/product-api-meli/internal/application/price/usecases"
	"github.com/viictormg/product-api-meli/internal/application/product/usecases"
	handlerPrice "github.com/viictormg/product-api-meli/internal/infra/api/handler/price"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product"
	groupPrice "github.com/viictormg/product-api-meli/internal/infra/api/router/group/price"
	group "github.com/viictormg/product-api-meli/internal/infra/api/router/group/product"
	"github.com/viictormg/product-api-meli/internal/infra/clients/db"
	"github.com/viictormg/product-api-meli/internal/infra/clients/producer"
	"github.com/viictormg/product-api-meli/internal/infra/events"
	repoHistory "github.com/viictormg/product-api-meli/internal/infra/repository/product_history"

	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(NewHTTPServer),
		fx.Provide(db.NewPostgresConnection),
		fx.Provide(db.NewRedisConnection),
		fx.Provide(producer.NewKafkaProducer),
		// respositories
		fx.Provide(repoHistory.NewProductHistoryRepository),
		fx.Provide(repoHistory.NewProductCacheHistoryRepository),
		fx.Provide(events.NewPriceEvent),

		// usecases
		fx.Provide(usecasePrice.NewPriceUsecase),
		fx.Provide(usecases.NewProductUsecase),

		// handlers
		fx.Provide(NewEchoGroup),
		fx.Provide(product.NewProductHandler),
		fx.Provide(handlerPrice.NewPriceHandler),

		// routes
		fx.Provide(group.NewProductInterfaceRoutes),
		fx.Provide(groupPrice.NewPriceInterfaceRoutes),

		// init functions
		fx.Invoke(func(*echo.Echo) {}),
		fx.Invoke(func(*group.ProductInterfaceRoute) {}),
		fx.Invoke(func(*groupPrice.PriceInterfaceRoute) {}),
	).Run()
}
