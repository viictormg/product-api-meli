package usecases

import (
	"context"
	"errors"
	"fmt"
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
	repository         ports.ProductRepositoryIF
	respositoryHistory ports.ProductHistoryRepositoryIF
	repositoryCache    ports.ProductCacheHistoryRepositoryIF
}

func NewProductUsecase(
	repository ports.ProductRepositoryIF,
	respositoryHistory ports.ProductHistoryRepositoryIF,
	repositoryCache ports.ProductCacheHistoryRepositoryIF,

) ProductUsecaseIF {
	return &productUsecase{
		repository:         repository,
		respositoryHistory: respositoryHistory,
		repositoryCache:    repositoryCache,
	}
}

func (p *productUsecase) UpdatePrice(ctx context.Context, product dto.UpdatePriceRequest) error {
	productFound, err := p.repository.GetProductByID(ctx, product.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	if productFound.Price == product.Price {
		return errors.New("price is the same")
	}

	priceIsInrage, err := p.priceIsInrage(product)
	if err != nil {
		return err
	}

	if !priceIsInrage {
		return errors.New("price is anomalous")
	}

	return p.UpdatePriceProduct(ctx, product)
}

func (p *productUsecase) UpdatePriceProduct(ctx context.Context, product dto.UpdatePriceRequest) error {
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

	err = p.SaveLimitsPriceCache(product)

	if err != nil {
		trx.Rollback()
	}

	return trx.Commit().Error
}

func (p *productUsecase) priceIsInrage(product dto.UpdatePriceRequest) (bool, error) {
	limits, err := p.GetLimitsPriceCache(product.ProductID)
	if err != nil || limits == nil {
		p.SaveLimitsPriceCache(product)
		limits, err = p.GetLimitsPriceCache(product.ProductID)
	}

	fmt.Println(err, limits)

	if err != nil || limits == nil {
		return false, err
	}
	fmt.Println(product.Price, "UP:", limits.Min, "DOWM:", limits.Max)

	return product.Price >= limits.Min && product.Price <= limits.Max, nil
}

func (p *productUsecase) SaveLimitsPriceCache(product dto.UpdatePriceRequest) error {
	limits, err := p.GetAverageAndDeviation(product.ProductID)

	if err != nil || limits == nil {
		return err
	}

	return p.repositoryCache.SaveProductHistory(product.ProductID, limits)
}

func (p *productUsecase) GetLimitsPriceCache(product string) (*dtoLimits.PriceLimitsDTO, error) {
	return p.repositoryCache.GetProductHistory(product)
}

func (p *productUsecase) GetAverageAndDeviation(productID string) (*dtoLimits.PriceLimitsDTO, error) {
	stats, err := p.respositoryHistory.GetAverageAndDeviation(productID)

	if err != nil {
		return &dtoLimits.PriceLimitsDTO{}, err
	}

	minLimit := stats.Average - constants.FactorLimitMin*stats.StandardDeviation
	maxLimit := stats.Average + constants.FactorLimitMax*stats.StandardDeviation

	return &dtoLimits.PriceLimitsDTO{
		Min: minLimit,
		Max: maxLimit,
	}, nil
}
