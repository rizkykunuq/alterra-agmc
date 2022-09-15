package database

import (
	"submission/day2/config"
	"submission/day2/models"

	"gorm.io/gorm/clause"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user *models.Users

	if e := config.DB.Where("id = ?", id).Take(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user *models.Users) (interface{}, error) {
	if e := config.DB.Select("Name", "Email", "Password").Create(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUserById(user *models.Users) (interface{}, error) {
	if e := config.DB.Model(&user).Clauses(clause.Returning{}).Where("id = ?", user.Id).Select("Name", "Email", "Password").Updates(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUserById(id int) (interface{}, error) {
	var user *models.Users

	if e := config.DB.Where("id = ?", id).Delete(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
