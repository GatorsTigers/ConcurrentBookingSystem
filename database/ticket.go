package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func BookSelectedSeats(ticket *models.Ticket, ShowSeatIds []uint32) error {
	tx := DbInstance.Db.Begin()

	if len(FetchShowSeats(ShowSeatIds)) != len(ShowSeatIds) {
		tx.Rollback()
		panic("Some seats are not available")
	}

	if txn := DbInstance.Db.Create(&ticket); txn.Error != nil {
		tx.Rollback()
		return txn.Error
	}

	if booked := UpdateShowSeats(ticket.TicketId, ShowSeatIds); !booked {
		tx.Rollback()
		panic("Cant Book selected seats.")
	}

	return tx.Commit().Error

}

func GetTicketsForUser(emailId string, tickets *[]models.Ticket) error {
	if txn := DbInstance.Db.Where("email_refer_id = ?", emailId).Find(tickets); txn.Error != nil {
		return txn.Error
	}
	return nil
}
