package dto

type UpdatePriceRequest struct {
	ProductID string  `json:"product_id" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}
