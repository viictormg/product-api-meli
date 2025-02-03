package dto

import "github.com/shopspring/decimal"

type UpdatePriceRequest struct {
	ProductID string          `param:"id" validate:"required"`
	Price     decimal.Decimal `json:"price" validate:"required"`
}
