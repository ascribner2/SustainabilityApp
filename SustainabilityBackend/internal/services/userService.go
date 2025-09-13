package services

import "github.com/ascribner/sustainabilityapp/internal/repos"

type UserService interface {
	RegisterUser()
}

type UserServiceImpl struct {
	r repos.UserRepo
}

func NewUserService(r repos.UserRepo) UserService {
	return &UserServiceImpl{
		r: r,
	}
}

func (us *UserServiceImpl) RegisterUser() {

}
