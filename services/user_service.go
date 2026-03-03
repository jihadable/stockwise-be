package services

import (
	"errors"

	"github.com/jihadable/stockwise-be/config"
	"github.com/jihadable/stockwise-be/helper"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(user *entity.User) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
	UpdateUserById(id string, user *entity.User) (*entity.User, error)
	VerifyUser(email, password string) (*entity.User, error)
	UpdatePassword(id, oldPassword, newPassword string) (*entity.User, error)
}

type UserServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *UserServiceImpl) AddUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = service.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserServiceImpl) GetUserById(id string) (*entity.User, error) {
	user := entity.User{}

	result := service.DB.Where("id = ?", id).First(&user)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserServiceImpl) UpdateUserById(id string, user *entity.User) (*entity.User, error) {
	savedUser := entity.User{}

	err := service.DB.Where("id = ?", id).First(&savedUser).Error
	if err != nil {
		return nil, err
	}

	savedUser.Username = user.Username
	savedUser.Bio = user.Bio

	err = service.DB.Where("id = ?", id).Updates(savedUser).Error
	if err != nil {
		return nil, err
	}

	return &savedUser, nil
}

func (service *UserServiceImpl) VerifyUser(email, password string) (*entity.User, error) {
	user := entity.User{}

	err := service.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserServiceImpl) UpdatePassword(id, oldPassword, newPassword string) (*entity.User, error) {
	if oldPassword == newPassword {
		return nil, errors.New("New password can not be same as old password")
	}

	user := entity.User{}

	err := service.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	hashedNewPassword, err := helper.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}

	user.Password = hashedNewPassword

	err = service.DB.Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserService(config *config.Config) UserService {
	return &UserServiceImpl{DB: config.DB, Redis: config.Redis}
}
