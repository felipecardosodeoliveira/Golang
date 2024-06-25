package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	cat1 := Category{
		Name: "Cozinha",
	}
	db.Create(&cat1)

	cat2 := Category{
		Name: "Eletronicos",
	}
	db.Create(&cat2)

	prod1 := Product{
		Name:       "Panela",
		Price:      80.00,
		Categories: []Category{cat1, cat2},
	}
	db.Create(&prod1)
}
