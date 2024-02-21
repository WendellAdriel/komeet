package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	UUID            string `gorm:"index"`
	Name            string
	Email           string
	EmailVerifiedAt *time.Time
	Password        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func NewUser(name, email, password string) User {
	return User{
		Name:            name,
		Email:           email,
		EmailVerifiedAt: nil,
		Password:        password,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.NewString()
	return
}
