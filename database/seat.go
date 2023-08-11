package database

import "github.com/GatorsTigers/ConcurrentBookingSystem/models"

func CreateSeats(seats []*models.Seat) error {

	if txn := DbInstance.Db.Create(seats); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetSeats(theaterId int, screenSeats *[]models.Seat) error {
	if txn := DbInstance.Db.Where("theater_comp_refer_id = ?", theaterId).Find(screenSeats); txn.Error != nil {
		return txn.Error
	}
	return nil
}
