package model

import (
	"time"

	"github.com/google/uuid"
)

type RentalTransaction struct {
	BaseModel
	RenterID             uuid.UUID  `gorm:"type:uuid;not null" json:"renter_id"`
	Renter               Renter     `gorm:"foreignKey:RenterID;constraint:OnDelete:RESTRICT" json:"renter"`
	CarID                uuid.UUID  `gorm:"type:uuid;not null" json:"car_id"`
	Car                  Car        `gorm:"foreignKey:CarID;constraint:OnDelete:RESTRICT" json:"car"`
	RentalStartDate      time.Time  `gorm:"not null" json:"rental_start_date"`
	RentalEndDate        time.Time  `gorm:"not null" json:"rental_end_date"`
	TotalRentalCost      float64    `gorm:"type:decimal(10,2);not null" json:"total_rental_cost"`
	Status               string     `gorm:"type:varchar(20);not null;default:'booked'" json:"status"`
	CarConditionOnReturn *string    `gorm:"type:text" json:"car_condition_on_return"`
	PenaltyFee           float64    `gorm:"type:decimal(10,2);default:0" json:"penalty_fee"`
	FinalTotalPayment    *float64   `gorm:"type:decimal(10,2)" json:"final_total_payment"`
	ActualReturnDate     *time.Time `gorm:"type:timestamp" json:"actual_return_date"`
}
