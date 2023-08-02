package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func AddUser(user *models.User) (*models.User, error) {
	if txn := DbInstance.Db.Create(&user); txn.Error != nil {
		return nil, txn.Error
	}
	return user, nil
}

func GetUserByEmailID(email string) (*models.User, error) {
	var user *models.User
	if txn := DbInstance.Db.Find(&user); txn.Error != nil {
		return nil, txn.Error
	}
	return user, nil
}
