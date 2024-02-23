package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	. "komeet/core"
	"time"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	UUID            string `gorm:"index"`
	Name            string
	Email           string
	EmailVerifiedAt *time.Time
	Password        string
	Active          bool `gorm:"default:false"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func NewUser(name, email, password string) User {
	logger := App.Logger()

	passwordHash, err := hashPassword(password)
	if err != nil {
		logger.Panic().Err(err).Msg("Error hashing password")
	}

	return User{
		Name:            name,
		Email:           email,
		EmailVerifiedAt: nil,
		Password:        passwordHash,
	}
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.NewString()
	return
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
