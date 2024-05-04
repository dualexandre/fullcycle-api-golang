package test

import (
	"testing"

	"github.com/dualexandre/fullcycle-api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := entity.NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}
