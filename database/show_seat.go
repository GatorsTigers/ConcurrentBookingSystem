package database

import "github.com/GatorsTigers/ConcurrentBookingSystem/models"

func GetTotalBookingAmount(showSeatIds []uint32) float64 {
	var amount float64
	if err := DbInstance.Db.Model(&models.ShowSeat{}).Select("sum(price)").Where("show_seat_id in ?", showSeatIds).Row().Scan(&amount); err != nil {
		panic(err)
	}
	return amount
}

func UpdateShowSeats(ticketId uint32, showSeatIds []uint32) bool {
	if txn := DbInstance.Db.Model(&models.ShowSeat{}).Where("show_seat_id in ?", showSeatIds).Updates(map[string]interface{}{"ticket_refer_id": ticketId, "is_available": false}); txn.Error != nil {
		panic(txn.Error)
	}
	return true
}

func FetchShowSeats(showSeatIds []uint32) []models.ShowSeat {
	var showSeats []models.ShowSeat
	if txn := DbInstance.Db.Model(&models.ShowSeat{}).Where("show_seat_id in ? and is_available=true", showSeatIds).Find(&showSeats); txn.Error != nil {
		panic(txn.Error)
	}
	return showSeats
}
