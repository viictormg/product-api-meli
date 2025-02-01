package product

import (
	"context"

	"github.com/viictormg/product-api-meli/internal/application/ports"
	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.ProductRepositoryIF {
	return &ProductRepository{db}
}

func (pr *ProductRepository) UpdatePrice(
	ctx context.Context, productId string, price float64,
) (*gorm.DB, error) {
	tx := pr.db.WithContext(ctx).Begin()

	tx.Where("id = ?", productId).Updates(entity.ProductEntity{Price: price})

	return tx, nil

}
