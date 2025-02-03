package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	dtoLimits "github.com/viictormg/product-api-meli/internal/application/product/dto"
	"github.com/viictormg/product-api-meli/internal/infra/api/handler/product/dto"
	"github.com/viictormg/product-api-meli/mocks"
)

var (
	ctx         = context.Background()
	requestMock = dto.UpdatePriceRequest{
		ProductID: "MLB4279481140",
		Price:     decimal.NewFromFloat(100.00),
	}

	limitsCache = dtoLimits.PriceLimitsDTO{
		Min:          decimal.NewFromFloat(80.00),
		Max:          decimal.NewFromFloat(120.00),
		CurrentPrice: decimal.NewFromFloat(100.00),
	}
)

func TestUpdatePrice(t *testing.T) {
	t.Run("should return an error price the same", func(t *testing.T) {
		respositoryHistoryMock := mocks.NewProductHistoryRepositoryIF(t)
		repositoryCacheMock := mocks.NewProductCacheHistoryRepositoryIF(t)

		productApp := NewProductUsecase(respositoryHistoryMock, repositoryCacheMock)

		repositoryCacheMock.On("GetProductHistory", ctx, requestMock.ProductID).
			Return(&limitsCache, nil)

		err := productApp.UpdatePrice(ctx, requestMock)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "price is the same")
	})

	t.Run("should return an error save limits calcualted", func(t *testing.T) {
		respositoryHistoryMock := mocks.NewProductHistoryRepositoryIF(t)
		repositoryCacheMock := mocks.NewProductCacheHistoryRepositoryIF(t)

		repositoryCacheMock.On("GetProductHistory", ctx, requestMock.ProductID).
			Return(nil, errors.New("limits not found"))

		respositoryHistoryMock.On("GetAverageAndDeviation", requestMock.ProductID).
			Return(dtoLimits.PriceStatsDTO{
				Average:           decimal.Zero,
				StandardDeviation: decimal.Zero,
			}, nil)

		productApp := NewProductUsecase(respositoryHistoryMock, repositoryCacheMock)

		err := productApp.UpdatePrice(ctx, requestMock)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "no data found")
	})

}
