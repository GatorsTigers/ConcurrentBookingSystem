package database

import (
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func AddScreenShowScheduleInTheatre(screenShowSchedules []models.ScreenShowSchedule) ([]models.ScreenShowSchedule, error) {
	if txn := DbInstance.Db.Create(&screenShowSchedules); txn.Error != nil {
		return screenShowSchedules, txn.Error
	}
	return screenShowSchedules, nil
}

func GetShowScheduleForTheatre(theaterId uint32) ([]models.ScreenShowSchedule, error) {
	var screenShowSchedules []models.ScreenShowSchedule
	cur_ts := time.Now()
	if txn := DbInstance.Db.Where("start_time>=? and theater_comp_refer_id= ?", cur_ts, theaterId).Find(&screenShowSchedules); txn.Error != nil {
		return screenShowSchedules, txn.Error
	}
	return screenShowSchedules, nil
}
