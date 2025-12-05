package service

import (
	"car-rental-backend/internal/constants"
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/model"
	"car-rental-backend/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RentalTransactionService struct {
	repo       *repository.RentalTransactionRepository
	carRepo    *repository.CarRepository
	renterRepo *repository.RenterRepository
}

func NewRentalTransactionService(
	repo *repository.RentalTransactionRepository,
	carRepo *repository.CarRepository,
	renterRepo *repository.RenterRepository,
) *RentalTransactionService {
	return &RentalTransactionService{
		repo:       repo,
		carRepo:    carRepo,
		renterRepo: renterRepo,
	}
}

func (s *RentalTransactionService) Paginate(paginate dto.PaginateRequest) (dto.RentalTransactionPaginateResponse, error) {
	return s.repo.Paginate(paginate)
}

func (s *RentalTransactionService) Create(req *dto.CreateRentalTransactionRequest) error {
	if req.RentalEndAt.Before(req.RentalStartAt) {
		return fmt.Errorf("rental end date must be after start date")
	}

	renter, err := s.renterRepo.FindByID(req.RenterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("renter not found")
		}
		return fmt.Errorf("failed to find renter: %w", err)
	}

	car, err := s.carRepo.GetByID(req.CarID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("car not found")
		}
		return fmt.Errorf("failed to find car: %w", err)
	}

	if !car.IsCarAvailable() {
		return fmt.Errorf("car is currently not available for rental")
	}

	available, err := s.repo.IsCarAvailableForRentByPeriod(req.CarID, req.RentalStartAt, req.RentalEndAt)
	if err != nil {
		return fmt.Errorf("failed to check car availability: %w", err)
	}
	if !available {
		return fmt.Errorf("car is not available for the selected dates")
	}

	days := max(int(req.RentalEndAt.Sub(req.RentalStartAt).Hours()/24), 1)
	totalCost := float64(days) * car.RentalPricePerDay

	currentTime := time.Now()
	status := constants.RENTAL_STATUS_BOOKED
	if req.RentalStartAt.Before(currentTime) || req.RentalStartAt.Equal(currentTime) {
		status = constants.RENTAL_STATUS_IN_PROGRESS
	}

	tx := &model.RentalTransaction{
		RenterID:        renter.ID,
		CarID:           car.ID,
		RentalStartDate: req.RentalStartAt,
		RentalEndDate:   req.RentalEndAt,
		TotalRentalCost: totalCost,
		Status:          status,
	}

	if err := s.repo.Create(tx); err != nil {
		return fmt.Errorf("failed to create rental transaction: %w", err)
	}

	return nil
}

func (s *RentalTransactionService) Pickup(id uuid.UUID) error {
	tx, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("rental transaction not found")
		}
		return fmt.Errorf("failed to find rental transaction: %w", err)
	}

	if tx.Status != constants.RENTAL_STATUS_BOOKED {
		return fmt.Errorf("only booked transactions can be picked up")
	}

	if tx.RentalStartDate.After(time.Now()) {
		return fmt.Errorf("cannot pick up before rental start date")
	}

	tx.Status = constants.RENTAL_STATUS_IN_PROGRESS
	if err := s.repo.Update(tx); err != nil {
		return fmt.Errorf("failed to update rental transaction: %w", err)
	}

	return nil
}

func (s *RentalTransactionService) Return(id uuid.UUID, req *dto.ReturnRentalRequest) error {
	tx, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("rental transaction not found")
		}
		return fmt.Errorf("failed to find rental transaction: %w", err)
	}

	if tx.Status != constants.RENTAL_STATUS_IN_PROGRESS {
		return fmt.Errorf("only in-progress transactions can be returned")
	}

	tx.CarConditionOnReturn = req.CarConditionOnReturn
	tx.ActualReturnDate = req.ActualReturnDate
	tx.Status = constants.RENTAL_STATUS_COMPLETED

	if req.PenaltyFee != nil {
		tx.PenaltyFee = *req.PenaltyFee
	}

	if tx.Status == constants.RENTAL_STATUS_COMPLETED && tx.ActualReturnDate != nil {
		finalTotal := tx.TotalRentalCost + tx.PenaltyFee
		tx.FinalTotalPayment = &finalTotal
	}

	if err := s.repo.Update(tx); err != nil {
		return fmt.Errorf("failed to update rental transaction: %w", err)
	}

	return nil
}

func (s *RentalTransactionService) Cancel(id uuid.UUID) error {
	tx, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("rental transaction not found")
		}
		return fmt.Errorf("failed to find rental transaction: %w", err)
	}

	if tx.Status != constants.RENTAL_STATUS_BOOKED && tx.Status != constants.RENTAL_STATUS_IN_PROGRESS {
		return fmt.Errorf("only booked or in-progress transactions can be cancelled")
	}

	tx.Status = constants.RENTAL_STATUS_CANCELLED

	if err := s.repo.Update(tx); err != nil {
		return fmt.Errorf("failed to update rental transaction: %w", err)
	}

	return nil
}
