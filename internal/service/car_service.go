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

type CarService struct {
	media *MediaService
	repo  *repository.CarRepository
}

func NewCarService(media *MediaService, repo *repository.CarRepository) *CarService {
	return &CarService{media: media, repo: repo}
}

func (s *CarService) Paginate(paginate dto.PaginateRequest) (dto.CarPaginateResponse, error) {
	return s.repo.Paginate(paginate)
}

func (s *CarService) Create(car *dto.MutateCarRequest) error {
	if car.Photo == nil {
		return fmt.Errorf("car photo is required")
	}
	imagePath, err := s.media.Upload(constants.CAR_IMAGE_COLLECTION, car.Photo)
	if err != nil {
		return err
	}
	return s.repo.Create(&model.Car{
		Name:              car.Name,
		Brand:             car.Brand,
		ProductionYear:    car.ProductionYear,
		PassengerCapacity: car.PassengerCapacity,
		TransmissionType:  car.TransmissionType,
		LicensePlate:      car.LicensePlate,
		RentalPricePerDay: car.RentalPricePerDay,
		Status:            car.Status,
		PhotoURL:          imagePath,
	})
}

func (s *CarService) Update(id uuid.UUID, req *dto.MutateCarRequest) error {
	existingCar, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("car not exists")
		}
		return err
	}
	existingCar.Name = req.Name
	existingCar.Brand = req.Brand
	existingCar.ProductionYear = req.ProductionYear
	existingCar.PassengerCapacity = req.PassengerCapacity
	existingCar.TransmissionType = req.TransmissionType
	existingCar.LicensePlate = req.LicensePlate
	existingCar.RentalPricePerDay = req.RentalPricePerDay
	existingCar.Status = req.Status

	var previousImage string
	if req.Photo != nil {
		previousImage = existingCar.PhotoURL
		imagePath, err := s.media.Upload(constants.CAR_IMAGE_COLLECTION, req.Photo)
		if err != nil {
			return err
		}
		existingCar.PhotoURL = imagePath
	}

	if err := s.repo.Update(existingCar); err != nil {
		return err
	}

	if previousImage != "" {
		previousImage = strings.Split(previousImage, "/")[len(strings.Split(previousImage, "/"))-1]
		err := s.media.Delete(constants.CAR_IMAGE_COLLECTION, previousImage)
		if err != nil {
			slog.Error("failed to delete unused car image", "error", err)
		}
	}
	return nil
}

func (s *CarService) Delete(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("car with id %s not found", id)
		}
		return fmt.Errorf("failed to delete car: %w", err)
	}
	//not deleting image because of using soft delete
	return nil
}
