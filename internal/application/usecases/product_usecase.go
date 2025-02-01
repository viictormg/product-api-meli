package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	repo "github.com/viictormg/product-api-meli/internal/application/ports"
	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product/dto"
)

type ProductUsecaseIF interface {
	UpdatePrice(ctx context.Context, product dto.UpdatePriceRequest) error
}

func NewProductUsecase(
	repository repo.ProductRepositoryIF,
	respositoryHistory repo.ProductHistoryRepositoryIF,
) ProductUsecaseIF {
	return &productUsecase{
		repository:         repository,
		respositoryHistory: respositoryHistory,
	}
}

type productUsecase struct {
	repository         repo.ProductRepositoryIF
	respositoryHistory repo.ProductHistoryRepositoryIF
}

func (p *productUsecase) UpdatePrice(ctx context.Context, product dto.UpdatePriceRequest) error {
	priceIsInrage, err := p.priceIsInrage(product)
	if err != nil {
		return err
	}

	if !priceIsInrage {
		return errors.New("price is anomalous")
	}

	trx, err := p.repository.UpdatePrice(ctx, product.ProductID, product.Price)
	if err != nil {
		trx.Rollback()
		return err
	}
	productHistory := entity.NewProductHistoryEntity(
		product.ProductID,
		product.Price,
		time.Now().Format(time.DateOnly),
	)

	err = p.respositoryHistory.CreateProductHistory(ctx, trx, productHistory)

	if err != nil {
		trx.Rollback()
	}

	trx.Commit()

	return nil
}

func (p *productUsecase) priceIsInrage(product dto.UpdatePriceRequest) (bool, error) {
	stats, err := p.respositoryHistory.GetAverageAndDeviation(product.ProductID)

	if err != nil {
		return false, err
	}

	lowerLimit := stats.Average - 2*stats.StandardDeviation
	upperLimit := stats.Average + 2*stats.StandardDeviation
	fmt.Println("LOWER ", lowerLimit)
	fmt.Println("UPPER", upperLimit)

	return product.Price >= lowerLimit && product.Price <= upperLimit, nil
}
