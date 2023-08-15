package database

import (
	"fmt"
	"strconv"
	"strings"
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
	var query string = "select shows.show_id AS ShowReferId, seats.seat_id as SeatReferId " +
		"from shows " +
		"left join screens on screens.screen_id = shows.screen_refer_id " +
		"left join seats on seats.screen_refer_id = screens.screen_id " +
		"where shows.show_id in (?)"
	if txn := DbInstance.Db.Raw(query, getParsedShowIds(showIdSlice)).Scan(&showSeats); txn.Error != nil {
		return txn.Error
	}
	if txn := DbInstance.Db.Create(showSeats); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func getParsedShowIds(showIdSlice []uint32) string {
	var showIdStrings []string
	for _, i := range showIdSlice {
		showIdStrings = append(showIdStrings, strconv.FormatUint(uint64(i), 10))
	}
	fmt.Println(strings.Join(showIdStrings, ","))
	return strings.Join(showIdStrings, ",")
}

func GetShowsForTheatre(theaterId uint32, showSeats *[]models.ShowSeat) error {
	cur_ts := time.Now()
	if txn := DbInstance.Db.Table("show_seats").Joins("JOIN shows ON show_seats.show_refer_id=shows.show_id").Joins("JOIN screens ON screens.screen_id=shows.screen_refer_id").Where("theater_refer_id=? and start_time>=?", theaterId, cur_ts).Find(&showSeats); txn.Error != nil {
		return txn.Error
	}
	return nil
}
