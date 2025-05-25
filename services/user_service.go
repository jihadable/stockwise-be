package services

import (
	"stockwise-be/model/entity"
	"stockwise-be/model/response"
	"stockwise-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(user entity.User) (response.UserResponse, error)
	GetUserById(id string) (response.UserResponse, error)
	UpdateUserById(id string, user entity.User) (response.UserResponse, error)
	VerifyUser(email, password string) (response.UserResponse, error)
}

type UserServiceImpl struct {
	DB *gorm.DB
}

func (service *UserServiceImpl) AddUser(user entity.User) (response.UserResponse, error) {
	user.Id = uuid.NewString()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return response.UserResponse{}, err
	}
	user.Password = hashedPassword
	result := service.DB.Create(&user)

	err = result.Error
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusBadRequest, "Gagal menambahkan pengguna")
	}

	return *utils.UserToResponse(&user), nil
}

func (service *UserServiceImpl) GetUserById(id string) (response.UserResponse, error) {
	user := entity.User{}

	result := service.DB.Where("id = ?", id).First(&user)

	err := result.Error
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusNotFound, "Gagal mendapatkan pengguna")
	}

	return *utils.UserToResponse(&user), nil
}

func (service *UserServiceImpl) UpdateUserById(id string, user entity.User) (response.UserResponse, error) {
	savedUser := entity.User{}

	result := service.DB.Where("id = ?", id).First(&savedUser)

	err := result.Error
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusNotFound, "Pengguna tidak ditemukan")
	}

	savedUser.Username = user.Username
	savedUser.Bio = user.Bio

	result = service.DB.Where("id = ?", id).Updates(&savedUser)

	err = result.Error
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusNotFound, "Gagal memperbarui pengguna")
	}

	return *utils.UserToResponse(&savedUser), nil
}

func (service *UserServiceImpl) VerifyUser(email, password string) (response.UserResponse, error) {
	user := entity.User{}

	result := service.DB.Where("email = ?", email).First(&user)

	err := result.Error
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusNotFound, "Pengguna tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return response.UserResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Password salah")
	}

	return *utils.UserToResponse(&user), nil
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{DB: db}
}
