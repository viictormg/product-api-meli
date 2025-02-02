package product

import (
	"context"
	"fmt"

	"github.com/viictormg/product-api-meli/internal/application/product/ports"
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

func (pr *ProductRepository) GetProductByID(
	ctx context.Context, productId string,
) (entity.ProductEntity, error) {
	var product entity.ProductEntity

	err := pr.db.WithContext(ctx).Where("id = ?", productId).First(&product).Error

	return product, err
}

func (pr *ProductRepository) CreateProduct(
	products []entity.ProductEntity,
) (*gorm.DB, error) {

	err := pr.db.Create(&products)

	fmt.Println("ERRRR", err)

	return nil, nil
}
