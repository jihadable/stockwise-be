package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jihadable/stockwise-be/helper"
	"github.com/jihadable/stockwise-be/helper/mapper"
	"github.com/jihadable/stockwise-be/model/entity"
	"github.com/jihadable/stockwise-be/model/request"
	"github.com/jihadable/stockwise-be/services"
	"github.com/jihadable/stockwise-be/validator"
)

type UserHandler interface {
	PostUser(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	PutUserById(ctx *fiber.Ctx) error
	VerifyUser(ctx *fiber.Ctx) error
	UpdatePassword(ctx *fiber.Ctx) error
}

type UserHandlerImpl struct {
	Service   services.UserService
	Validator validator.UserValidator
}

func (handler *UserHandlerImpl) PostUser(ctx *fiber.Ctx) error {
	requestBody := request.UserRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidatePostUserRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := handler.Service.AddUser(&entity.User{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
		Bio:      requestBody.Bio,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Registration failed")
	}

	jwt, err := helper.GetJWT(user.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Fail to generate JWT")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": mapper.UserToResponse(user),
			"jwt":  jwt,
		},
	})
}

func (handler *UserHandlerImpl) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	user, err := handler.Service.GetUserById(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": mapper.UserToResponse(user),
		},
	})
}

func (handler *UserHandlerImpl) PutUserById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)
	requestBody := request.UpdateUserRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid resquest body")
	}

	err = handler.Validator.ValidatePutUserRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid resquest body")
	}

	user, err := handler.Service.UpdateUserById(userId, &entity.User{
		Username: requestBody.Username,
		Bio:      requestBody.Bio,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to update user")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": mapper.UserToResponse(user),
		},
	})
}

func (handler *UserHandlerImpl) VerifyUser(ctx *fiber.Ctx) error {
	requestBody := request.LoginRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidateVerifyUserRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := handler.Service.VerifyUser(requestBody.Email, requestBody.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Incorrect email or password")
	}

	jwt, err := helper.GetJWT(user.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Fail to generate JWT")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": mapper.UserToResponse(user),
			"jwt":  jwt,
		},
	})
}

func (handler *UserHandlerImpl) UpdatePassword(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	requestBody := request.UpdatePasswordRequest{}

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = handler.Validator.ValidateUpdatePasswordRequest(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := handler.Service.UpdatePassword(userId, requestBody.NewPasswrod)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Fail to change password")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
		},
	})
}

func NewUserHandler(service services.UserService, validator validator.UserValidator) UserHandler {
	return &UserHandlerImpl{Service: service, Validator: validator}
}
