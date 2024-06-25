package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// product belongs to category
type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
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

	// insert(db)
	// getFirst(db)

	// BelongsTo
	// category := Category{Name: "Eletrônicos"}
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Moto G 14",
	// 	Price:      1900.00,
	// 	CategoryId: category.ID,
	// })
	// var products []Product
	// db.Preload("Category").Where("category_id = 1").Find(&products)

	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Moto G 14",
		Price:      1900.00,
		CategoryId: category.ID,
	})

	db.Create(&SerialNumber{
		Number:    "1234",
		ProductId: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	// change all categories to 1
	// var products []Product
	// db.Preload("Category").Find(&products)
	// for _, product := range products {
	// 	// if product.CategoryId != 1 {
	// 	product.CategoryId = 1
	// 	db.Save(&product)
	// 	// }
	// }

}

// func insert(db *gorm.DB) {
// 	products := []Product{
// 		{Name: "Carro Polo", Price: 50000.00},
// 		{Name: "Moto", Price: 1200.00}}
// 	db.Create(&products)
// }

// func getFirst(db *gorm.DB) {
// var product Product
// db.First(&product)
// db.First(&product, "name = ?", "Smart TV Del")
// fmt.Println(product)

// var products []Product
// db.Find(&products)
// fmt.Println(products)

// var product Product
// db.Where("name = ?", "Sansung Galax").Find(&product)
// fmt.Println(product)

// var products []Product
// db.Where("name LIKE ?", "%Del%").Find(&products)
// fmt.Println(products)

// var p Product
// db.Find(&p, 2)
// fmt.Println(p.Name)
// p.Name = "Notebook Positivo"
// p.Price = 2000
// db.Save(&p)
// fmt.Println(p.Name)

// var p Product
// db.Find(&p, 2)
// db.Delete(&p)
// }
