package handlers

import (
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/utils"
	"github.com/jihadable/stockwise-be/validator"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	PostUser(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	PutUserById(ctx *fiber.Ctx) error
	VerifyUser(ctx *fiber.Ctx) error
}

type UserHandlerImpl struct {
	Service   services.UserService
	Validator validator.UserValidator
}

func (handler *UserHandlerImpl) PostUser(ctx *fiber.Ctx) error {
	userRequest := request.UserRequest{}

	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal bro")
	}

	err = handler.Validator.ValidatePostUserRequest(&userRequest)
	if err != nil {
		return err
	}

	user := *utils.RequestToUser(&userRequest)

	userResponse, err := handler.Service.AddUser(user)
	if err != nil {
		return err
	}

	token, err := utils.GenerateJWT(userResponse.Id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user":  userResponse,
			"token": token,
		},
	})
}

func (handler *UserHandlerImpl) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	userResponse, err := handler.Service.GetUserById(userId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": userResponse,
		},
	})
}

func (handler *UserHandlerImpl) PutUserById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)
	userRequest := request.UpdateUserRequest{}

	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal memperbarui pengguna")
	}

	err = handler.Validator.ValidatePutUserRequest(&userRequest)
	if err != nil {
		return err
	}

	user := *utils.UpdateUserRequestToUser(&userRequest)

	userResponse, err := handler.Service.UpdateUserById(userId, user)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": userResponse,
		},
	})
}

func (handler *UserHandlerImpl) VerifyUser(ctx *fiber.Ctx) error {
	loginRequest := request.LoginRequest{}

	err := ctx.BodyParser(&loginRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal masuk")
	}

	err = handler.Validator.ValidateVerifyUserRequest(&loginRequest)
	if err != nil {
		return err
	}

	userResponse, err := handler.Service.VerifyUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return err
	}

	token, err := utils.GenerateJWT(userResponse.Id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user":  userResponse,
			"token": token,
		},
	})
}

func NewUserHandler(service services.UserService, validator validator.UserValidator) UserHandler {
	return &UserHandlerImpl{Service: service, Validator: validator}
}
