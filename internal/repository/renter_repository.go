package repository

import (
	"car-rental-backend/internal/dto"
	"car-rental-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RenterRepository struct {
	db *gorm.DB
}

func NewRenterRepository(db *gorm.DB) *RenterRepository {
	return &RenterRepository{db: db}
}

func (r *RenterRepository) Create(renter *model.Renter) error {
	return r.db.Create(&renter).Error
}

func (r *RenterRepository) Update(renter *model.Renter) error {
	return r.db.Save(&renter).Error
}

func (r *RenterRepository) Delete(id uuid.UUID) error {
	tx := r.db.Delete(&model.Renter{}, "id = ?", id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func (r *RenterRepository) FindByID(id uuid.UUID) (*model.Renter, error) {
	var renter model.Renter
	err := r.db.First(&renter, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &renter, nil
}

func (r *RenterRepository) Paginate(paginate dto.PaginateRequest) (dto.RenterPaginateResponse, error) {
	var paginatedRenters dto.RenterPaginateResponse
	var renters []model.Renter
	var total int64

	err := r.db.Model(&model.Renter{}).Count(&total).Error
	if err != nil {
		return paginatedRenters, err
	}

	err = r.db.Limit(paginate.PageSize).Offset((paginate.PageNumber - 1) * paginate.PageSize).Order("updated_at DESC").Find(&renters).Error
	if err != nil {
		return paginatedRenters, err
	}

	totalPages := int(total) / paginate.PageSize
	if int(total)%paginate.PageSize != 0 {
		totalPages++
	}

	paginatedRenters = dto.RenterPaginateResponse{
		Page:  paginate.PageNumber,
		Total: int(total),
		From:  totalPages,
		Items: renters,
	}

	return paginatedRenters, nil
}
