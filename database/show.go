package database

import (
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func AddShowsInTheatre(shows *[]models.Show) error {
	if txn := DbInstance.Db.Create(shows); txn.Error != nil {
		return txn.Error
	}
	// var showSeats []models.ShowSeat
	var showIds []uint32
	for _, show := range *shows {
		showIds = append(showIds, show.ShowId)
	}
	if txn := DbInstance.Db.Model(&models.Show{}).Joins("Screen").Where("show_id?", showIds); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetShowsForTheatre(theaterId uint32, showSeats *[]models.ShowSeat) error {
	cur_ts := time.Now()
	if txn := DbInstance.Db.Model(&models.Show{}).Joins("Screen").Where("screen.theater_refer_id=? and show.start_time=?", theaterId, cur_ts).Find(&showSeats); txn.Error != nil {
		return txn.Error
	}
	return nil
}
