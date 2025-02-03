package usecases

import (
	"context"
	"errors"
	"time"

	dtoLimits "github.com/viictormg/product-api-meli/internal/application/product/dto"
	ports "github.com/viictormg/product-api-meli/internal/application/product/ports"
	"github.com/viictormg/product-api-meli/internal/domain/constants"
	"github.com/viictormg/product-api-meli/internal/domain/entity"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product/dto"
)

type ProductUsecaseIF interface {
	UpdatePrice(ctx context.Context, product dto.UpdatePriceRequest) error
}
type productUsecase struct {
	respositoryHistory ports.ProductHistoryRepositoryIF
	repositoryCache    ports.ProductCacheHistoryRepositoryIF
}

func NewProductUsecase(
	respositoryHistory ports.ProductHistoryRepositoryIF,
	repositoryCache ports.ProductCacheHistoryRepositoryIF,

) ProductUsecaseIF {
	return &productUsecase{
		respositoryHistory: respositoryHistory,
		repositoryCache:    repositoryCache,
	}
}

func (p *productUsecase) UpdatePrice(ctx context.Context, product dto.UpdatePriceRequest) error {
	priceIsInrage, err := p.priceIsInrage(ctx, product)
	if err != nil {
		return err
	}

	if !priceIsInrage {
		return errors.New("price is anomalous")
	}

	return p.UpdatePriceProduct(ctx, product)
}

func (p *productUsecase) UpdatePriceProduct(ctx context.Context, product dto.UpdatePriceRequest) error {
	productHistory := entity.NewProductHistoryEntity(
		product.ProductID,
		product.Price,
		time.Now().Format(time.DateOnly),
	)

	trx, err := p.respositoryHistory.CreateProductHistory(ctx, productHistory)

	if err != nil {
		trx.Rollback()
	}

	err = p.SaveLimitsPriceCache(ctx, product)

	if err != nil {
		trx.Rollback()
	}

	return trx.Commit().Error
}

func (p *productUsecase) priceIsInrage(ctx context.Context, product dto.UpdatePriceRequest) (bool, error) {
	limits, err := p.GetLimitsPriceCache(product.ProductID)
	if err != nil || limits == nil {
		p.SaveLimitsPriceCache(ctx, product)
		limits, err = p.GetLimitsPriceCache(product.ProductID)
	}

	if err != nil || limits == nil {
		return false, err
	}

	if limits.CurrentPrice == product.Price {
		return true, errors.New("price is the same")
	}

	return product.Price >= limits.Min && product.Price <= limits.Max, nil
}

func (p *productUsecase) SaveLimitsPriceCache(ctx context.Context, product dto.UpdatePriceRequest) error {
	limits, err := p.GetLimitsPriceDB(product.ProductID)
	if err != nil || limits == nil {
		return err
	}

	productFound, err := p.respositoryHistory.GetLastPrice(ctx, product.ProductID)
	if err != nil || productFound.IsEmpty() {
		return err
	}

	limits.CurrentPrice = productFound.Price

	return p.repositoryCache.SaveProductHistory(product.ProductID, limits)
}

func (p *productUsecase) GetLimitsPriceCache(product string) (*dtoLimits.PriceLimitsDTO, error) {
	return p.repositoryCache.GetProductHistory(product)
}

func (p *productUsecase) GetLimitsPriceDB(productID string) (*dtoLimits.PriceLimitsDTO, error) {
	stats, err := p.respositoryHistory.GetAverageAndDeviation(productID)

	if err != nil {
		return &dtoLimits.PriceLimitsDTO{}, err
	}

	minLimit := stats.Average - constants.FactorLimitMin*stats.StandardDeviation
	maxLimit := stats.Average + constants.FactorLimitMax*stats.StandardDeviation

	if minLimit < 0 {
		minLimit = 0
	}

	return &dtoLimits.PriceLimitsDTO{
		Min: minLimit,
		Max: maxLimit,
	}, nil
}
