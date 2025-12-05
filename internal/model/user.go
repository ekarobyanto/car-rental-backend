package model

type User struct {
	BaseModel
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"-"`
}

func (User) TableName() string {
	return "users"
}
