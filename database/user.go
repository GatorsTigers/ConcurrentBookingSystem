package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"golang.org/x/crypto/bcrypt"
)

// AddUser creates a user
func AddUser(user *models.User) (*models.User, error) {

	password := []byte(user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, er := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if er != nil {
		return *models.User{}, er
	}

	user.Password = string(hashedPassword)

	err := DbInstance.Db.Create(&user).Error
	if err != nil {
		return *models.User{}, err
	}
	return GetUserByEmail(db, user.Email)
}

func GetUserByEmailID(email string) (*models.User, error) {
	var user *models.User
	if txn := DbInstance.Db.Find(&user); txn.Error != nil {
		return nil, txn.Error
	}
	return user, nil
}
