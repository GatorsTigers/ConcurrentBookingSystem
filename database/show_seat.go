package database

func GetTotalBookingAmount(showSeatIds []uint32) float64 {
	var amount float64
	var query string = "select sum(amount) as amount from show_seats where show_seat_id in (?)"
	if txn := DbInstance.Db.Raw(query, getParsedIds(showSeatIds)).Scan(&amount); txn.Error != nil {
		panic(txn.Error)
	}
	return amount
}

func UpdateShowSeats(ticketId uint32, showSeatIds []uint32) bool {
	var query string = "update show_seats set ticket_id = ? where show_seats_id in (?)"
	txn := DbInstance.Db.Raw(query, ticketId, getParsedIds(showSeatIds))
	if txn.Error != nil {
		panic(txn.Error)
	}
	return true
}
