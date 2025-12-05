package model

type Renter struct {
	BaseModel
	Name                 string  `gorm:"type:varchar(255);not null" json:"name"`
	IDCardNumber         string  `gorm:"type:varchar(50);unique;not null" json:"id_card_number"`
	PhoneNumber          string  `gorm:"type:varchar(20);not null" json:"phone_number"`
	Address              string  `gorm:"type:text;not null" json:"address"`
	DrivingLicenseNumber string  `gorm:"type:varchar(50);not null" json:"driving_license_number"`
	IDCardPhotoURL       *string `gorm:"type:varchar(500)" json:"id_card_photo_url"`
	LicensePhotoURL      *string `gorm:"type:varchar(500);column:driving_license_photo_url" json:"license_photo_url"`
}
