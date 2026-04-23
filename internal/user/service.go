package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	Create(userData *InputUser) (*User, error)
}

type userService struct{ db *gorm.DB }

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) Create(userData *InputUser) (*User, error) {
	newUser := &User{Name: userData.Name, Mobile: userData.Mobile}
	us.db.Create(newUser)
	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}
	return newUser, nil

}
