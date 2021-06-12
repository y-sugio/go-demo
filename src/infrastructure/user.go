package infrastructure

import (
	"domain/model"
	"domain/repository"
	"infrastructure/database/query"
)

type UserRepository struct{}

func NewTaskRepository() repository.UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindByID(id int64) (user *model.User, err error) {
	user, err = query.Select("user", id)
	if err != nil {
		return
	}
	return
}

func (ur *UserRepository) Create(user *model.User) (*model.User, error) {
	err := query.Ins("user", user)
	return user, err
}
