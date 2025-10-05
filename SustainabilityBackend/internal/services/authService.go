package services

import (
	"strings"

	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/repos"
	"golang.org/x/crypto/bcrypt"
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
	// Make the email lowercase to make it easy to check against
	lowerEmail := strings.ToLower(u.GetEmail())

	passHash, err := as.r.GetPassword(lowerEmail)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(u.GetPass()))

	// Compare hashes
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}
