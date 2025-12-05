package dto

import "github.com/gofiber/fiber/v2"

type PaginateRequest struct {
	PageSize   int
	PageNumber int
}

func ParsePaginateRequest(c *fiber.Ctx) PaginateRequest {
	pageSize := c.QueryInt("size")
	if pageSize <= 0 {
		pageSize = 10
	}

	pageNumber := c.QueryInt("page")
	if pageNumber <= 0 {
		pageNumber = 1
	}

	return PaginateRequest{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
}

type PaginateResponse[T any] struct {
	Page  int `json:"page"`
	From  int `json:"from"`
	Total int `json:"total"`
	Items []T `json:"items"`
}
