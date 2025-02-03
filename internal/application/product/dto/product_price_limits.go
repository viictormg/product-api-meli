package dto

import "github.com/shopspring/decimal"

type PriceLimitsDTO struct {
	Min          decimal.Decimal
	Max          decimal.Decimal
	CurrentPrice decimal.Decimal
}
