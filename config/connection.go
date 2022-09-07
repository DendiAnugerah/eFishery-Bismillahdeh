package config

import (
	"Fish/model/category"
	"Fish/model/product"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conection() (*gorm.DB, error) {
	var dsn = "host=localhost user=postgres password=dendi dbname=day7 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&product.Product{}, &category.Category{})

	return db, nil
}
