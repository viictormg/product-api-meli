package entity

type ProductEntity struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Price float64
}

func (ProductEntity) TableName() string {
	return "product"
}

func NewProductEntity(id string, name string, price float64) ProductEntity {
	return ProductEntity{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
