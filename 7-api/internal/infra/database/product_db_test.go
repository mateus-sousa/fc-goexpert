package database

import (
	"fmt"
	"github.com/mateus-sousa/goexpert/7-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func BeforeEach(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	return db
}
func TestCreateNewProduct(t *testing.T) {
	db := BeforeEach(t)
	product, err := entity.NewProduct("Telefone", 60.0)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db := BeforeEach(t)
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Produto %d", i), rand.Float64()*100.0)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDb := NewProduct(db)
	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, "Produto 10", products[9].Name)

	products, err = productDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 11", products[0].Name)
	assert.Equal(t, "Produto 20", products[9].Name)

	products, err = productDb.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Produto 21", products[0].Name)
	assert.Equal(t, "Produto 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db := BeforeEach(t)
	product, err := entity.NewProduct("Notebook", 1000.00)
	assert.NoError(t, err)
	db.Create(product)
	productDb := NewProduct(db)
	storedProduct, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Notebook", storedProduct.Name)
}

func TestUpdateProduct(t *testing.T) {
	db := BeforeEach(t)
	product, err := entity.NewProduct("Notebook", 1000.00)
	assert.NoError(t, err)
	db.Create(product)
	productDb := NewProduct(db)
	product.Name = "Notebook Atualizado"
	err = productDb.Update(product)
	assert.NoError(t, err)
	updatedProduct, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Notebook Atualizado", updatedProduct.Name)
}

func TestDeleteProduct(t *testing.T) {
	db := BeforeEach(t)
	product, err := entity.NewProduct("Notebook", 1000.00)
	assert.NoError(t, err)
	db.Create(product)
	productDb := NewProduct(db)
	product.Name = "Notebook Atualizado"
	err = productDb.Delete(product.ID.String())
	assert.NoError(t, err)
	storedProduct, err := productDb.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, storedProduct)
}
