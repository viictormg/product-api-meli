package dto

type PriceHistory struct {
	ProductID string  `json:"product_id"`
	OrderDate string  `json:"order_date"`
	Price     float64 `json:"price"`
}
