package dto

import "github.com/shopspring/decimal"

type PriceHistory struct {
	ProductID string          `json:"product_id"`
	OrderDate string          `json:"order_date"`
	Price     decimal.Decimal `json:"price"`
}
