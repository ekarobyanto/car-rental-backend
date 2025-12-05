package model

import "car-rental-backend/internal/constants"

type Car struct {
	BaseModel
	Name              string  `gorm:"not null" json:"name"`
	Brand             string  `gorm:"not null" json:"brand"`
	ProductionYear    int     `gorm:"not null" json:"production_year"`
	PassengerCapacity int     `gorm:"not null" json:"passenger_capacity"`
	TransmissionType  string  `gorm:"not null" json:"transmission_type"`
	LicensePlate      string  `gorm:"uniqueIndex;not null" json:"license_plate"`
	RentalPricePerDay float64 `gorm:"not null" json:"rental_price_per_day"`
	Status            string  `gorm:"not null;default:'available'" json:"status"`
	PhotoURL          string  `json:"photo_url"`
}

func (c Car) IsCarAvailable() bool {
	return c.Status == constants.CAR_STATUS_AVAILABLE
}
