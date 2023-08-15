package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateSeats(seats *[]models.Seat) error {

	if txn := DbInstance.Db.Create(seats); txn.Error != nil {
		return txn.Error
	}
	return nil
}

// func GetShowSeats(theaterId int, screenSeats *[]models.Seat) error {
// 	cur_ts := time.Now()
// 	if txn := DbInstance.Db.Model(&models.Show{}).Joins("Screen").Where("screen.theater_refer_id=? and show.start_time=?", theaterId, cur_ts).Find(&screenSeats); txn.Error != nil {
// 		return txn.Error
// 	}
// 	return nil
// }

func GetAllSeatsInTheater(theaterId int, screenSeats *[]models.Seat) error {
	if txn := DbInstance.Db.Model(&models.Seat{}).Joins("Screen").Where("theater_refer_id=?", theaterId).Find(&screenSeats); txn.Error != nil {
		return txn.Error
	}
	return nil
}
