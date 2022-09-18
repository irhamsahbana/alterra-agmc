package models

import (
	"day02/config"
	"day02/models/entity"
)

func GetUsers() []entity.User {
	var users []entity.User

	config.App.Sql.Find(&users)

	return users
}

func GetUser(id int) *entity.User {
	var user entity.User

	config.App.Sql.First(&user, id)

	return &user
}

func GetUserByUsername(username string) *entity.User {
	var user entity.User

	config.App.Sql.Where("username = ?", username).First(&user)

	return &user
}

func CreateUser(user *entity.User) *entity.User {
	config.App.Sql.Create(&user)

	return user
}

func UpdateUser(id int, user *entity.User) *entity.User {
	tx := config.App.Sql.Model(&user).Where("id = ?", id).Updates(user)

	if tx.Error != nil {
		return nil
	}

	tx.First(&user, id)

	return user
}

func DeleteUser(id int) *entity.User {
	var user entity.User

	config.App.Sql.First(&user, id)
	config.App.Sql.Delete(&user, id)

	return &user
}
