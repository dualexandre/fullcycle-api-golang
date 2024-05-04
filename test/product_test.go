package test

import (
	"testing"

	"github.com/dualexandre/fullcycle-api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := entity.NewProduct("phone", 100.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.Equal(t, "phone", product.Name)
	assert.Equal(t, 100.00, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := entity.NewProduct("", 100.00)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := entity.NewProduct("phone", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := entity.NewProduct("phone", -100)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := entity.NewProduct("phone", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
