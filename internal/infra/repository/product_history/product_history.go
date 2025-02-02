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
	tx *gorm.DB,
	productHistory entity.ProductHistoryEntity,
) error {

	if err := tx.Create(&productHistory).Error; err != nil {
		return err
	}

	return nil

}
