package entity

type ProductEntity struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Price float64
}

func (ProductEntity) TableName() string {
	return "product"
}
