package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/dualexandre/fullcycle-api-golang/internal/entity"
	"github.com/dualexandre/fullcycle-api-golang/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Phone", 1000.00)
	productDB := database.NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	produdctDB := database.NewProduct(db)
	products, err := produdctDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = produdctDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = produdctDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 25", products[4].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Phone", 1000.00)
	db.Create(product)
	productDB := database.NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Phone", 1000.00)
	db.Create(product)

	productDB := database.NewProduct(db)
	product.Name = "Phone Updated"
	err = productDB.Update(product)
	assert.NoError(t, err)

	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Phone Updated", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Phone", 1000.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := database.NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
}
