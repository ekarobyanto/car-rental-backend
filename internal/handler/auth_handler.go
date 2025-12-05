package handler

import (
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/service"
	"car-rental-backend/internal/utils"
	"car-rental-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", h.Login)
	auth.Post("/register", h.Register)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req, err := utils.ParseValidateRequest[dto.LoginRequest](c)
	result, err := h.service.Login(&req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, err.Error(), nil)
	}

	return response.Success(c, "Login successful", result)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	req, err := utils.ParseValidateRequest[dto.CreateUserRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	err = h.service.Create(&req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to create user", err.Error())
	}

	return response.Created(c, "User created successfully", nil)
}
