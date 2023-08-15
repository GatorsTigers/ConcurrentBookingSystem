package database

import (
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func AddShowsInTheatre(shows *[]models.Show) error {
	if txn := DbInstance.Db.Create(shows); txn.Error != nil {
		return txn.Error
	}
	var showIdSlice []uint32
	for _, show := range *shows {
		showIdSlice = append(showIdSlice, show.ShowId)
	}
	var showSeats []models.ShowSeat
	if txn := DbInstance.Db.Raw("select shows.show_id, seats.seat_id "+
		"from shows "+
		"left join screens on screens.screen_id = shows.screen_refer_id "+
		"left join seats on seats.screen_refer_id = screens.screen_id "+
		"where shows.show_id in ", showIdSlice).Scan(&showSeats); txn.Error != nil {
		return txn.Error
	}
	if txn := DbInstance.Db.Create(showSeats); txn.Error != nil {
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
