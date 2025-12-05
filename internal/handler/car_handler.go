package handler

import (
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/service"
	"car-rental-backend/internal/utils"
	"car-rental-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CarHandler struct {
	service *service.CarService
}

func NewCarHandler(service *service.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) RegisterRoutes(router fiber.Router) {
	cars := router.Group("/cars")
	cars.Get("/", h.Paginate)
	cars.Post("/", h.Create)
	cars.Put("/:id", h.Update)
	cars.Delete("/:id", h.Delete)
}

func (h *CarHandler) Paginate(c *fiber.Ctx) error {
	paginate, err := h.service.Paginate(dto.ParsePaginateRequest(c))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return response.Success(c, "Successfully fetch cars", paginate)
}

func (h *CarHandler) Create(c *fiber.Ctx) error {
	req, err := utils.ParseValidateRequest[dto.MutateCarRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := req.ValidatePhoto(); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid photo", err.Error())
	}

	err = h.service.Create(&req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to create car", err.Error())
	}

	return response.Created(c, "Car created successfully", nil)
}

func (h *CarHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	carUUID, err := uuid.Parse(id)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid car ID", err.Error())
	}

	req, err := utils.ParseValidateRequest[dto.MutateCarRequest](c)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := req.ValidatePhoto(); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid photo", err.Error())
	}

	err = h.service.Update(carUUID, &req)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to update car", err.Error())
	}

	return response.Success(c, "Car updated successfully", nil)
}

func (h *CarHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	carUUID, err := uuid.Parse(id)
	if err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid car ID", err.Error())
	}

	err = h.service.Delete(carUUID)
	if err != nil {
		return response.Failed(c, fiber.StatusUnprocessableEntity, "Failed to delete car", err.Error())
	}

	return response.Success(c, "Car deleted successfully", nil)
}
