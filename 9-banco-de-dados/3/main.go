package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// product belongs to category
// category has many products
type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryId   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

// product has one
type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// db.Create(&Category{
	// 	Name: "Eletronicos",
	// })

	// db.Create(&Product{
	// 	Name:       "Moto G 14",
	// 	Price:      1900.00,
	// 	CategoryId: 1,
	// })

	// db.Create(&Product{
	// 	Name:       "Sansung Galax",
	// 	Price:      1500.00,
	// 	CategoryId: 1,
	// })

	// db.Create(&SerialNumber{
	// 	Number:    "1234",
	// 	ProductId: 1,
	// })

	// db.Create(&SerialNumber{
	// 	Number:    "4321",
	// 	ProductId: 2,
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Println(category.Name, product.Name, product.SerialNumber.Number)
		}
	}

}
