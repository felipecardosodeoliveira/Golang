package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/felipecardosodeoliveira/Golang/12-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Product{})
	product, _ := entity.NewProduct("Iphone 13", 10000.0)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.Name, product.Name)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*10)
		assert.Nil(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Product{})
	product, _ := entity.NewProduct("Iphone 13", 10000.0)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	productMatch, err := productDB.FindById(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, productMatch.Price, product.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Product{})
	product, _ := entity.NewProduct("Iphone 13", 10000.0)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	// product dosent exist yet
	var fakeP entity.Product
	err = productDB.Update(&fakeP)
	assert.NotNil(t, err)

	product.Name = "Azuss"
	err = productDB.Update(product)
	assert.Nil(t, err)
	assert.Equal(t, "Azuss", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Product{})
	product, _ := entity.NewProduct("Iphone 13", 10000.0)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

}
