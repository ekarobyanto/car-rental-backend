package service

import (
	"car-rental-backend/internal/constants"
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/model"
	"car-rental-backend/internal/repository"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RenterService struct {
	repo         *repository.RenterRepository
	mediaService *MediaService
}

func NewRenterService(repo *repository.RenterRepository, mediaService *MediaService) *RenterService {
	return &RenterService{
		repo:         repo,
		mediaService: mediaService,
	}
}

func (s *RenterService) Paginate(paginate dto.PaginateRequest) (dto.RenterPaginateResponse, error) {
	return s.repo.Paginate(paginate)
}

func (s *RenterService) Create(req *dto.MutateRenterRequest) error {
	renter := &model.Renter{
		Name:                 req.Name,
		IDCardNumber:         req.IDCardNumber,
		PhoneNumber:          req.PhoneNumber,
		Address:              req.Address,
		DrivingLicenseNumber: req.DrivingLicenseNumber,
	}

	if req.IDCardPhoto != nil {
		url, err := s.mediaService.Upload(constants.RENTER_ID_CARD_COLLECTION, req.IDCardPhoto)
		if err != nil {
			return fmt.Errorf("failed to upload ID card photo: %w", err)
		}
		renter.IDCardPhotoURL = &url
	}

	if req.LicensePhoto != nil {
		url, err := s.mediaService.Upload(constants.RENTER_LICENSE_COLLECTION, req.LicensePhoto)
		if err != nil {
			return fmt.Errorf("failed to upload license photo: %w", err)
		}
		renter.LicensePhotoURL = &url
	}

	if err := s.repo.Create(renter); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("renter with ID card %s or license number %s already exists", req.IDCardNumber, req.DrivingLicenseNumber)
		}
		return fmt.Errorf("failed to create renter: %w", err)
	}

	return nil
}

func (s *RenterService) Update(id uuid.UUID, req *dto.MutateRenterRequest) error {
	renter, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("renter not found")
	}

	renter.Name = req.Name
	renter.IDCardNumber = req.IDCardNumber
	renter.PhoneNumber = req.PhoneNumber
	renter.Address = req.Address
	renter.DrivingLicenseNumber = req.DrivingLicenseNumber

	var prevIdCardPhotoUrl, prevLicensePhotoUrl string
	if req.RemoveIDCardPhoto != nil && *req.RemoveIDCardPhoto {
		if renter.IDCardPhotoURL != nil {
			prevIdCardPhotoUrl = *renter.IDCardPhotoURL
		}
		renter.IDCardPhotoURL = nil
	}

	if req.RemoveLicensePhoto != nil && *req.RemoveLicensePhoto {
		if renter.LicensePhotoURL != nil {
			prevLicensePhotoUrl = *renter.LicensePhotoURL
		}
		renter.LicensePhotoURL = nil
	}

	if req.IDCardPhoto != nil {
		if renter.IDCardPhotoURL != nil {
			prevIdCardPhotoUrl = *renter.IDCardPhotoURL
		}
		url, err := s.mediaService.Upload(constants.RENTER_ID_CARD_COLLECTION, req.IDCardPhoto)
		if err != nil {
			return fmt.Errorf("failed to upload ID card photo: %w", err)
		}
		renter.IDCardPhotoURL = &url
	}

	if req.LicensePhoto != nil {
		if renter.LicensePhotoURL != nil {
			prevLicensePhotoUrl = *renter.LicensePhotoURL
		}
		url, err := s.mediaService.Upload(constants.RENTER_LICENSE_COLLECTION, req.LicensePhoto)
		if err != nil {
			return fmt.Errorf("failed to upload license photo: %w", err)
		}
		renter.LicensePhotoURL = &url
	}

	if err := s.repo.Update(renter); err != nil {
		return fmt.Errorf("failed to update renter: %w", err)
	}

	if prevIdCardPhotoUrl != "" {
		prevIdCardPhotoUrl = strings.Split(prevIdCardPhotoUrl, "/")[len(strings.Split(prevIdCardPhotoUrl, "/"))-1]
		err := s.mediaService.Delete(constants.RENTER_ID_CARD_COLLECTION, prevIdCardPhotoUrl)
		if err != nil {
			slog.Error("failed to delete ID card photo", "error", err)
		}
	}

	if prevLicensePhotoUrl != "" {
		prevLicensePhotoUrl = strings.Split(prevLicensePhotoUrl, "/")[len(strings.Split(prevLicensePhotoUrl, "/"))-1]
		err := s.mediaService.Delete(constants.RENTER_LICENSE_COLLECTION, prevLicensePhotoUrl)
		if err != nil {
			slog.Error("failed to delete license photo", "error", err)
		}
	}

	return nil
}

func (s *RenterService) Delete(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("renter with id %s not found", id)
		}
		return fmt.Errorf("failed to delete renter: %w", err)
	}

	//not deleting images because of using soft delete
	return nil
}
