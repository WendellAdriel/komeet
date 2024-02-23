package repositories

import (
	"fmt"
	. "komeet/core"
	. "komeet/models"
)

func CreateUser(user *User) {
	logger := App.Logger()
	result := App.DB.Create(&user)
	if result.Error != nil {
		logger.Panic().Err(result.Error).Msgf("Failed creating user: %v", user)
	}
}

func UpdateUser(user *User) {
	logger := App.Logger()
	result := App.DB.Save(&user)
	if result.Error != nil {
		logger.Panic().Err(result.Error).Msgf("Failed updating user: %v", user)
	}
}

func GetUserBy(field string, value any) (User, bool) {
	var user User

	App.DB.Where(fmt.Sprintf("%s = ?", field), value).
		First(&user)

	return checkReturnUser(user)
}

func GetUserForLogin(email string) (User, bool) {
	var user User

	App.DB.Where("email = ?", email).
		Where("email_verified_at IS NOT NULL").
		Where("active = ?", true).
		First(&user)

	return checkReturnUser(user)
}

func checkReturnUser(user User) (User, bool) {
	if user.ID == 0 {
		return user, false
	}

	return user, true
}
