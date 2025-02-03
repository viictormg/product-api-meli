// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	dto "github.com/viictormg/product-api-meli/internal/infra/api/handler/product/dto"
)

// ProductUsecaseIF is an autogenerated mock type for the ProductUsecaseIF type
type ProductUsecaseIF struct {
	mock.Mock
}

// UpdatePrice provides a mock function with given fields: ctx, product
func (_m *ProductUsecaseIF) UpdatePrice(ctx context.Context, product dto.UpdatePriceRequest) error {
	ret := _m.Called(ctx, product)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePrice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdatePriceRequest) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductUsecaseIF creates a new instance of ProductUsecaseIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductUsecaseIF(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductUsecaseIF {
	mock := &ProductUsecaseIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
