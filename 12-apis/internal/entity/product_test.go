package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Notebooke Del", 3500.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Notebooke Del", product.Name)
	assert.Equal(t, 3500.0, product.Price)
	assert.Nil(t, product.ValidateProduct())
}

func TestWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 1500.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestWhenPriceIsInvalid(t *testing.T) {
	product1, err := NewProduct("PC Gamer", 0.0)
	assert.Nil(t, product1)
	assert.Equal(t, ErrInvalidPrice, err)

	product2, err := NewProduct("PC Gamer", -1.0)
	assert.Nil(t, product2)
	assert.Equal(t, ErrInvalidPrice, err)
}
