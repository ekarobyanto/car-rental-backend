package handler

import (
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/service"
	"car-rental-backend/internal/utils"
	"car-rental-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RenterHandler struct {
	service *service.RenterService
}

func NewRenterHandler(service *service.RenterService) *RenterHandler {
	return &RenterHandler{service: service}
}

func (h *RenterHandler) RegisterRoutes(router fiber.Router) {
	renters := router.Group("/renters")
	renters.Get("/", h.Paginate)
	renters.Post("/", h.Create)
	renters.Put("/:id", h.Update)
	renters.Delete("/:id", h.Delete)
}

func (h *RenterHandler) Create(c *fiber.Ctx) error {
	req, err := utils.ParseValidateRequest[dto.MutateRenterRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	if err := req.ValidatePhotos(); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	err = h.service.Create(&req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, err.Error(), nil)
	}

	return response.Created(c, "Renter created successfully", nil)
}

func (h *RenterHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	req, err := utils.ParseValidateRequest[dto.MutateRenterRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	if err := req.ValidatePhotos(); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	err = h.service.Update(id, &req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, err.Error(), nil)
	}

	return response.Success(c, "Renter updated successfully", nil)
}

func (h *RenterHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := h.service.Delete(id); err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, err.Error(), nil)
	}

	return response.Success(c, "Renter deleted successfully", nil)
}

func (h *RenterHandler) Paginate(c *fiber.Ctx) error {
	paginate := dto.ParsePaginateRequest(c)
	renters, err := h.service.Paginate(paginate)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, err.Error(), nil)
	}

	return response.Success(c, "Renters retrieved successfully", renters)
}
