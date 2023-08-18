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
	if txn := DbInstance.Db.Model(&models.ShowSeat{}).Where("show_seat_id in ?", showSeatIds).UpdateColumns(models.ShowSeat{TicketReferId: ticketId, IsAvailable: false}); txn.Error != nil {
		panic(txn.Error)
	}
	return true
}
