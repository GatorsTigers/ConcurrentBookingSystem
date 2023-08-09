package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"golang.org/x/crypto/bcrypt"
)

// AddUser creates a user
func AddUser(user *models.User) error {

	password := []byte(user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = DbInstance.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmailID(email string) (models.User, error) {
	var user models.User
	if txn := DbInstance.Db.Find(&user); txn.Error != nil {
		return user, txn.Error
	}
	return user, nil
}

// ValidateUserCredentials checks if password is valid for a particular email
func ValidateUserCredentials(email string, password string) (bool, error) {
	passwordByte := []byte(password)
	user, e := GetUserByEmailID(email)
	if e == nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), passwordByte)
		if err == nil {
			return true, nil
		}
		return false, err
	}
	return false, e
}
