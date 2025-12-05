package repository

import (
	"car-rental-backend/internal/constants"
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RentalTransactionRepository struct {
	db *gorm.DB
}

func NewRentalTransactionRepository(db *gorm.DB) *RentalTransactionRepository {
	return &RentalTransactionRepository{db: db}
}

func (r *RentalTransactionRepository) Paginate(paginate dto.PaginateRequest) (dto.RentalTransactionPaginateResponse, error) {
	var paginatedTx dto.RentalTransactionPaginateResponse
	var transactions []model.RentalTransaction
	var total int64

	err := r.db.Model(&model.RentalTransaction{}).Count(&total).Error
	if err != nil {
		return paginatedTx, err
	}

	err = r.db.Preload("Renter").Preload("Car").Limit(paginate.PageSize).Offset((paginate.PageNumber - 1) * paginate.PageSize).Order("updated_at DESC").Find(&transactions).Error
	if err != nil {
		return paginatedTx, err
	}

	totalPages := int(total) / paginate.PageSize
	if int(total)%paginate.PageSize != 0 {
		totalPages++
	}

	paginatedTx = dto.RentalTransactionPaginateResponse{
		Page:  paginate.PageNumber,
		Total: int(total),
		From:  totalPages,
		Items: transactions,
	}

	return paginatedTx, nil
}

func (r *RentalTransactionRepository) Create(req *model.RentalTransaction) error {
	tx := r.db.Begin()
	defer tx.Rollback()

	err := tx.Create(&req).Error
	if err != nil {
		return err
	}

	if req.Status == constants.RENTAL_STATUS_IN_PROGRESS {
		err = tx.Model(&model.Car{}).Where("id = ?", req.CarID).
			Update("status", constants.CAR_STATUS_RENTED).Error
		if err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

func (r *RentalTransactionRepository) Update(req *model.RentalTransaction) error {
	return r.db.Save(&req).Error
}

func (r *RentalTransactionRepository) Delete(id uuid.UUID) error {
	req := r.db.Delete(&model.RentalTransaction{}, "id = ?", id)
	if req.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return req.Error
}

func (r *RentalTransactionRepository) FindByID(id uuid.UUID) (*model.RentalTransaction, error) {
	var req model.RentalTransaction
	err := r.db.Preload("Renter").Preload("Car").Where("id = ?", id).First(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *RentalTransactionRepository) IsCarAvailableForRentByPeriod(carID uuid.UUID, startDate, endDate time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&model.RentalTransaction{}).
		Where("car_id = ?", carID).
		Where("status IN (?, ?)", constants.RENTAL_STATUS_BOOKED, constants.RENTAL_STATUS_IN_PROGRESS).
		Where("NOT (rental_end_date < ? OR rental_start_date > ?)", startDate, endDate).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
