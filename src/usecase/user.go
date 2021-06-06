package usecase

import (
	"domain/model"
	"domain/repository"
)

type UserUseCase interface {
	FindByID(id int64) (*model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

// FindByID taskをIDで取得するときのユースケース
func (uu *userUseCase) FindByID(id int64) (*model.User, error) {
	foundUser, err := uu.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}