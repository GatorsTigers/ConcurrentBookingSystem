package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func BookSelectedSeats(ticket *models.Ticket) error {
	if txn := DbInstance.Db.Create(&ticket); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetTicketsForUser(emailId string, tickets *[]models.Ticket) error {
	if txn := DbInstance.Db.Where("email_refer_id = ?", emailId).Find(tickets); txn.Error != nil {
		return txn.Error
	}
	return nil
}
