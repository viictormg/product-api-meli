package db

import (
	"fmt"

	"github.com/viictormg/product-api-meli/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(config *config.Config) *gorm.DB {
	configDb := config.GetDbConfig()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=%s",
			configDb.DbUser,
			configDb.DbPass,
			configDb.DbHost,
			configDb.DbName,
			configDb.DbPort,
			configDb.SslMode,
		),
		// DSN:                  fmt.Sprintf("user=postgresql password=root host=localhost dbname=products_db port=5432 sslmode=disable"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
