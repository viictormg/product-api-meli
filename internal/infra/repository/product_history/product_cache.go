package producthistory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func (pr *ProductCacheHistoryRepository) SaveProductHistory(productId string, stats *dto.PriceLimitsDTO) error {
	//	Key: "price:MLB4432316952"
	//
	// Value: { "last_prices": [22.01, 19.809, 21.5], "avg": 21.1, "stddev": 1.2 }
	// TTL: 24 horas
	key := fmt.Sprintf("price:%s", productId)
	value, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	// Set the value in Redis with a TTL of 24 hours
	err = pr.db.Set(context.Background(), key, value, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductCacheHistoryRepository) GetProductHistory(productId string) (*dto.PriceLimitsDTO, error) {
	key := fmt.Sprintf("price:%s", productId)
	value, err := pr.db.Get(context.Background(), key).Result()
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
