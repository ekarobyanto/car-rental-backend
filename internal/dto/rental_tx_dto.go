package dto

import (
	"car-rental-backend/internal/model"
	"time"

	"github.com/google/uuid"
)

type CreateRentalTransactionRequest struct {
	RenterID      uuid.UUID `json:"renter_id" validate:"required"`
	CarID         uuid.UUID `json:"car_id" validate:"required"`
	RentalStartAt time.Time `json:"rental_start_at" validate:"required"`
	RentalEndAt   time.Time `json:"rental_end_at" validate:"required"`
}

type ReturnRentalRequest struct {
	CarConditionOnReturn *string    `json:"car_condition_on_return" validate:"omitempty"`
	PenaltyFee           *float64   `json:"penalty_fee" validate:"omitempty,min=0"`
	ActualReturnDate     *time.Time `json:"actual_return_date" validate:"omitempty"`
}

type RentalTransactionPaginateResponse PaginateResponse[model.RentalTransaction]
