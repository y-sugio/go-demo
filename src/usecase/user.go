package usecase

import (
	"domain/model"
	"domain/repository"
)

type UserUseCase interface {
	FindByID(id int64) (*model.User, error)
	Create(user *model.User) (*model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) FindByID(id int64) (*model.User, error) {
	foundUser, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uc *userUseCase) Create(user *model.User) (*model.User, error) {
	createdUser, err := uc.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
