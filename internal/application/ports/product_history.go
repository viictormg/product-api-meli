package ports

import (
	"context"

	"github.com/viictormg/product-api-meli/internal/application/dto"
	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductHistoryRepositoryIF interface {
	GetAverageAndDeviation(productID string) (dto.PriceStatsDTO, error)
	CreateProductHistory(
		ctx context.Context,
		tx *gorm.DB,
		productHistory entity.ProductHistoryEntity,
	) error
}

type ProductCacheHistoryRepositoryIF interface {
	SaveProductHistory(productId string, limits *dto.PriceLimitsDTO) error
	GetProductHistory(productId string) (*dto.PriceLimitsDTO, error)
}
