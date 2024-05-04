package database

import "github.com/dualexandre/fullcycle-api-golang/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
