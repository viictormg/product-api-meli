package ports

import (
	"context"

	"gorm.io/gorm"
)

type ProductRepositoryIF interface {
	UpdatePrice(ctx context.Context, productId string, price float64) (*gorm.DB, error)
}
