package handler

import (
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/service"
	"car-rental-backend/internal/utils"
	"car-rental-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RentalTransactionHandler struct {
	service *service.RentalTransactionService
}

func NewRentalTransactionHandler(service *service.RentalTransactionService) *RentalTransactionHandler {
	return &RentalTransactionHandler{service: service}
}

func (h *RentalTransactionHandler) RegisterRoutes(router fiber.Router) {
	transactions := router.Group("/rental-transactions")
	transactions.Get("/", h.Paginate)
	transactions.Post("/", h.Create)
	transactions.Put("/:id/pickup", h.Pickup)
	transactions.Put("/:id/return", h.Return)
	transactions.Put("/:id/cancel", h.Cancel)
}

func (h *RentalTransactionHandler) Paginate(c *fiber.Ctx) error {
	paginate, err := h.service.Paginate(dto.ParsePaginateRequest(c))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return response.Success(c, "Successfully fetch rental transactions", paginate)
}

func (h *RentalTransactionHandler) Create(c *fiber.Ctx) error {
	req, err := utils.ParseValidateRequest[dto.CreateRentalTransactionRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	err = h.service.Create(&req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to create rental transaction", err.Error())
	}

	return response.Created(c, "Rental transaction created successfully", nil)
}

func (h *RentalTransactionHandler) Pickup(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	err = h.service.Pickup(id)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to pickup rental transaction", err.Error())
	}

	return response.Success(c, "Rental transaction picked up successfully", nil)
}

func (h *RentalTransactionHandler) Return(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	req, err := utils.ParseValidateRequest[dto.ReturnRentalRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	err = h.service.Return(id, &req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to update rental transaction", err.Error())
	}

	return response.Success(c, "Rental transaction updated successfully", nil)
}

func (h *RentalTransactionHandler) Cancel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid transaction ID", err.Error())
	}

	err = h.service.Cancel(id)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to cancel rental transaction", err.Error())
	}

	return response.Success(c, "Rental transaction canceled successfully", nil)
}
