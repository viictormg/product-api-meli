package entity

import "github.com/shopspring/decimal"

type ProductEntity struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Price decimal.Decimal
}

func (ProductEntity) TableName() string {
	return "product"
}

func NewProductEntity(id string, name string, price decimal.Decimal) ProductEntity {
	return ProductEntity{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
