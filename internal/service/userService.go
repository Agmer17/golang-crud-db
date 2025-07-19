package service

import (
	"context"

	"github.com/Agmer17/golang-crud-db.git/internal/model"
	"github.com/Agmer17/golang-crud-db.git/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func (svc *UserService) GetAllData(userCtx context.Context) ([]model.UserModel, error) {
	data, err := svc.repo.GetAllData(userCtx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
