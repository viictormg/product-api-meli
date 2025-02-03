// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	dto "github.com/viictormg/product-api-meli/internal/application/product/dto"
)

// ProductCacheHistoryRepositoryIF is an autogenerated mock type for the ProductCacheHistoryRepositoryIF type
type ProductCacheHistoryRepositoryIF struct {
	mock.Mock
}

// GetProductHistory provides a mock function with given fields: ctx, productId
func (_m *ProductCacheHistoryRepositoryIF) GetProductHistory(ctx context.Context, productId string) (*dto.PriceLimitsDTO, error) {
	ret := _m.Called(ctx, productId)

	if len(ret) == 0 {
		panic("no return value specified for GetProductHistory")
	}

	var r0 *dto.PriceLimitsDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*dto.PriceLimitsDTO, error)); ok {
		return rf(ctx, productId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *dto.PriceLimitsDTO); ok {
		r0 = rf(ctx, productId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.PriceLimitsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveProductHistory provides a mock function with given fields: ctx, productId, limits
func (_m *ProductCacheHistoryRepositoryIF) SaveProductHistory(ctx context.Context, productId string, limits *dto.PriceLimitsDTO) error {
	ret := _m.Called(ctx, productId, limits)

	if len(ret) == 0 {
		panic("no return value specified for SaveProductHistory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *dto.PriceLimitsDTO) error); ok {
		r0 = rf(ctx, productId, limits)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductCacheHistoryRepositoryIF creates a new instance of ProductCacheHistoryRepositoryIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductCacheHistoryRepositoryIF(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductCacheHistoryRepositoryIF {
	mock := &ProductCacheHistoryRepositoryIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
