package ports

import (
	"context"

	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductRepositoryIF interface {
	UpdatePrice(ctx context.Context, productId string, price float64) (*gorm.DB, error)
	GetProductByID(ctx context.Context, productId string) (entity.ProductEntity, error)
	CreateProduct(product []entity.ProductEntity) (*gorm.DB, error)
}
