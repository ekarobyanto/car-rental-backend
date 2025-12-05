package dto

import (
	"car-rental-backend/internal/model"
	"fmt"
	"mime/multipart"
)

type MutateRenterRequest struct {
	Name                 string                `form:"name" validate:"required,min=2"`
	IDCardNumber         string                `form:"id_card_number" validate:"required"`
	PhoneNumber          string                `form:"phone_number" validate:"required"`
	Address              string                `form:"address" validate:"required,min=5"`
	DrivingLicenseNumber string                `form:"driving_license_number" validate:"required"`
	IDCardPhoto          *multipart.FileHeader `form:"id_card_photo" validate:"omitempty"`
	LicensePhoto         *multipart.FileHeader `form:"license_photo" validate:"omitempty"`
	RemoveIDCardPhoto    *bool                 `form:"remove_id_card" validate:"omitempty"`
	RemoveLicensePhoto   *bool                 `form:"remove_license" validate:"omitempty"`
}

func (r *MutateRenterRequest) ValidatePhotos() error {
	if r.IDCardPhoto != nil {
		if r.IDCardPhoto.Size > 5*1024*1024 {
			return fmt.Errorf("ID card photo size must not exceed 5MB")
		}
		contentType := r.IDCardPhoto.Header.Get("Content-Type")
		if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/jpg" {
			return fmt.Errorf("ID card photo must be jpeg or png")
		}
	}
	if r.LicensePhoto != nil {
		if r.LicensePhoto.Size > 5*1024*1024 {
			return fmt.Errorf("license photo size must not exceed 5MB")
		}
		contentType := r.LicensePhoto.Header.Get("Content-Type")
		if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/jpg" {
			return fmt.Errorf("license photo must be jpeg or png")
		}
	}

	return nil
}

type RenterPaginateResponse PaginateResponse[model.Renter]
