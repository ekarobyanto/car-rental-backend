package dto

import (
	"car-rental-backend/internal/model"
	"fmt"
	"mime/multipart"
)

type MutateCarRequest struct {
	Name              string                `form:"name" validate:"required,min=2"`
	Brand             string                `form:"brand" validate:"required,min=2"`
	ProductionYear    int                   `form:"production_year" validate:"required,min=1886"`
	PassengerCapacity int                   `form:"passenger_capacity" validate:"required,min=1"`
	TransmissionType  string                `form:"transmission_type" validate:"required,oneof=manual automatic"`
	LicensePlate      string                `form:"license_plate" validate:"required"`
	RentalPricePerDay float64               `form:"rental_price_per_day" validate:"required,min=0"`
	Status            string                `form:"status" validate:"required,oneof=available rented maintenance"`
	Photo             *multipart.FileHeader `form:"photo" validate:"omitempty"`
}

func (r *MutateCarRequest) ValidatePhoto() error {
	if r.Photo != nil {
		if r.Photo.Size > 5*1024*1024 {
			return fmt.Errorf("photo size must not exceed 5MB")
		}
		contentType := r.Photo.Header.Get("Content-Type")
		if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/jpg" {
			return fmt.Errorf("photo must be jpeg or png")
		}
	}

	return nil
}

type CarPaginateResponse PaginateResponse[model.Car]
