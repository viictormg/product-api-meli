package dto

import "github.com/shopspring/decimal"

type PriceStatsDTO struct {
	Average           decimal.Decimal
	StandardDeviation decimal.Decimal
}
