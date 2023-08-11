package database

import (
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func AddScreenShowScheduleInTheatre(screenShowSchedules *[]models.ScreenShowSchedule) error {
	if txn := DbInstance.Db.Create(screenShowSchedules); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetShowScheduleForTheatre(theaterId uint32, screenShowSchedules *[]models.ScreenShowSchedule) error {
	cur_ts := time.Now()
	if txn := DbInstance.Db.Model(&models.ScreenShowSchedule{}).Joins("Screen").Joins("Show").Where("start_time>=? and theater_comp_refer_id= ?", cur_ts, theaterId).Find(screenShowSchedules); txn.Error != nil {
		return txn.Error
	}
	return nil
}
