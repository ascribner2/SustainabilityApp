package services

import (
	"errors"
	"strings"

	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/repos"
	"golang.org/x/crypto/bcrypt"
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
	if len(nu.GetEmail()) <= 5 || len(nu.GetPass()) <= 5 {
		return errors.New("invalid email or password")
	}

	// Make the email lowercase to make it easy to check against
	lowerEmail := strings.ToLower(nu.GetEmail())

	// Generate the hash for the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.GetPass()), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := us.r.InsertUser(lowerEmail, string(hashedPass)); err != nil {
		return err
	}

	return nil
}
