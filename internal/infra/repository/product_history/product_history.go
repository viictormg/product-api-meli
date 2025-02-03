package producthistory

import (
	"context"
	"fmt"

	"github.com/viictormg/product-api-meli/internal/application/product/dto"
	"github.com/viictormg/product-api-meli/internal/application/product/ports"
	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductHistoryRepository struct {
	db *gorm.DB
}

func NewProductHistoryRepository(db *gorm.DB) ports.ProductHistoryRepositoryIF {
	return &ProductHistoryRepository{db}
}

func (pr *ProductHistoryRepository) GetAverageAndDeviation(productId string) (dto.PriceStatsDTO, error) {
	var priceStats dto.PriceStatsDTO
	query := fmt.Sprintf(`SELECT AVG(price) as average, STDDEV(price) as standard_deviation FROM price_history WHERE product_id = '%s'`, productId)

	if err := pr.db.Raw(query).Scan(&priceStats).Error; err != nil {
		return priceStats, err
	}

	return priceStats, nil
}

func (pr *ProductHistoryRepository) CreateProductHistory(
	ctx context.Context,
	productHistory entity.ProductHistoryEntity,
) (*gorm.DB, error) {
	tx := pr.db.WithContext(ctx).Begin()

	err := tx.Create(&productHistory).Error

	return tx, err

}

func (pr *ProductHistoryRepository) GetLastPrice(
	ctx context.Context, productId string,
) (entity.ProductHistoryEntity, error) {

	var productHistory entity.ProductHistoryEntity

	err := pr.db.WithContext(ctx).
		Where("product_id = ?", productId).
		Order("created_at DESC, order_closed DESC").
		First(&productHistory).
		Error

	if err != nil {
		return productHistory, err
	}

	return productHistory, nil
}
