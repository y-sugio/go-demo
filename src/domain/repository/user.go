package repository

import (
	"domain/model"
)

type UserRepository interface {
	FindByID(id int64) (*model.User, error)
}
