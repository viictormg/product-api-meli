package dto

type UpdatePriceRequest struct {
	ProductID string  `param:"id" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}
