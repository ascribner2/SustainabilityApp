package services

import (
	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/repos"
)

type AuthService interface {
	UserLogin(entity.UserEntity) (bool, error)
}

type AuthServiceImpl struct {
	r repos.AuthRepo
}

func NewAuthService(r repos.AuthRepo) AuthService {
	return &AuthServiceImpl{
		r: r,
	}
}

func (as *AuthServiceImpl) UserLogin(u entity.UserEntity) (bool, error) {
	passHash, err := as.r.GetPassword(u.GetEmail())
	if err != nil {
		return false, err
	}

	if u.GetPass() == passHash {
		return true, nil
	} else {
		return false, nil
	}
}
