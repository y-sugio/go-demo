package infrastructure

import (
	"domain/model"
	"domain/repository"
)

type UserRepository struct {}

func NewTaskRepository() repository.UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindByID(id int64) (*model.User, error) {
	user := &model.User{Id: id, Name: "sugio", Email: "sugio@test.com"}

	return user, nil
}