package database

import "github.com/GatorsTigers/ConcurrentBookingSystem/models"

func CreateSeats(seats []*models.Seat) error {

	if txn := DbInstance.Db.Create(seats); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetSeats(theaterId int) ([]models.Seat, error) {
	var seats []models.Seat
	if txn := DbInstance.Db.Find(&seats).Where("theater_comp_refer_id= ?", theaterId).Find(&seats); txn.Error != nil {
		return seats, txn.Error
	}
	return seats, nil
}
