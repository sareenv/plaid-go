package services

import (
	"errors"

	"github.com/sareenv/plaid-go/models"
	"github.com/sareenv/plaid-go/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	GetOrCreateUser(email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) GetOrCreateUser(email string) (*models.User, error) {
	user, err := us.repo.FindUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err // real error, bail out
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user doesn't exist, create them
		user = &models.User{Email: email}
		if err = us.repo.CreateUser(user); err != nil {
			return nil, err
		}
	}
	return user, nil
}
