package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Telefone", 800.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Telefone", product.Name)
	assert.Equal(t, 800.0, product.Price)

}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 800.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Telefone", 0.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestNewProductValidate(t *testing.T) {
	product, err := NewProduct("Telefone", 800.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
