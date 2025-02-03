package producthistory

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/viictormg/product-api-meli/internal/application/product/dto"
	"github.com/viictormg/product-api-meli/internal/application/product/ports"
)

type ProductCacheHistoryRepository struct {
	db *redis.Client
}

func NewProductCacheHistoryRepository(db *redis.Client) ports.ProductCacheHistoryRepositoryIF {
	return &ProductCacheHistoryRepository{db}
}

func (pr *ProductCacheHistoryRepository) SaveProductHistory(ctx context.Context, productId string, stats *dto.PriceLimitsDTO) error {
	key := productId
	value, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	err = pr.db.Set(ctx, key, value, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductCacheHistoryRepository) GetProductHistory(ctx context.Context, productId string) (*dto.PriceLimitsDTO, error) {
	key := productId
	value, err := pr.db.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("No history found for product ID", productId)
		return nil, errors.New("no history found")
	} else if err != nil {
		return nil, err
	}

	var stats dto.PriceLimitsDTO
	err = json.Unmarshal([]byte(value), &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
