package services

import (
	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/repos"
)

type UserService interface {
	RegisterUser(entity.UserEntity) error
}

type UserServiceImpl struct {
	r repos.UserRepo
}

func NewUserService(r repos.UserRepo) UserService {
	return &UserServiceImpl{
		r: r,
	}
}

func (us *UserServiceImpl) RegisterUser(nu entity.UserEntity) error {
	if err := us.r.InsertUser(nu.GetEmail(), nu.GetPass()); err != nil {
		return err
	}

	return nil
}
