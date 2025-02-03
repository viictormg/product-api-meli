package ports

import (
	"context"

	"github.com/viictormg/product-api-meli/internal/application/product/dto"
	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductHistoryRepositoryIF interface {
	GetAverageAndDeviation(productID string) (dto.PriceStatsDTO, error)

	CreateProductHistory(
		ctx context.Context,
		productHistory entity.ProductHistoryEntity,
	) (*gorm.DB, error)
	GetLastPrice(ctx context.Context, productId string) (entity.ProductHistoryEntity, error)
}

type ProductCacheHistoryRepositoryIF interface {
	SaveProductHistory(productId string, limits *dto.PriceLimitsDTO) error
	GetProductHistory(productId string) (*dto.PriceLimitsDTO, error)
}
