package repository

import (
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (r *CarRepository) Create(car *model.Car) error { return r.db.Create(&car).Error }
func (r *CarRepository) Update(car *model.Car) error { return r.db.Save(&car).Error }
func (r *CarRepository) Delete(id uuid.UUID) error {
	tx := r.db.Delete(&model.Car{}, "id = ?", id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func (r *CarRepository) GetByID(id uuid.UUID) (*model.Car, error) {
	var car model.Car
	err := r.db.Where("id = ?", id).First(&car).Error
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *CarRepository) Paginate(paginate dto.PaginateRequest) (dto.CarPaginateResponse, error) {
	var paginatedCars dto.CarPaginateResponse
	var cars []model.Car
	var total int64

	err := r.db.Model(&model.Car{}).Count(&total).Error
	if err != nil {
		return paginatedCars, err
	}

	err = r.db.Limit(paginate.PageSize).Offset((paginate.PageNumber - 1) * paginate.PageSize).Order("updated_at DESC").Find(&cars).Error
	if err != nil {
		return paginatedCars, err
	}

	totalPages := int(total) / paginate.PageSize
	if int(total)%paginate.PageSize != 0 {
		totalPages++
	}

	paginatedCars = dto.CarPaginateResponse{
		Page:  paginate.PageNumber,
		Total: int(total),
		From:  totalPages,
		Items: cars,
	}

	return paginatedCars, nil
}
