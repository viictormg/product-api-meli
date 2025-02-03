package dto

import "github.com/shopspring/decimal"

type ProductResponseDTO struct {
	ItemID   string          `json:"item_id"`
	Price    decimal.Decimal `json:"price"`
	Anomaly  bool            `json:"anomaly"`
	Metadata interface{}     `json:"metadata"`
}
