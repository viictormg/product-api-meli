package entity

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/shopspring/decimal"
)

type ProductHistoryEntity struct {
	ID          string          `gorm:"primaryKey"`
	ProductId   string          `gorm:"column:product_id"`
	Price       decimal.Decimal `gorm:"column:price"`
	OrderClosed string          `gorm:"column:order_closed"`
}

func (p ProductHistoryEntity) IsEmpty() bool {
	return p.ID == ""
}

func (ProductHistoryEntity) TableName() string {
	return "price_history"
}

func NewProductHistoryEntity(
	productId string,
	price decimal.Decimal,
	orderClosed string,
) ProductHistoryEntity {
	uid, _ := gonanoid.New(20)

	return ProductHistoryEntity{
		ID:          uid,
		ProductId:   productId,
		Price:       price,
		OrderClosed: orderClosed,
	}
}
