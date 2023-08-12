package database

import (
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func AddScreenShowScheduleInTheatre(screenShowSchedules *[]models.Show) error {
	if txn := DbInstance.Db.Create(screenShowSchedules); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetShowScheduleForTheatre(theaterId uint32, screenShowSchedules *[]models.Show) error {
	cur_ts := time.Now()
	if txn := DbInstance.Db.Model(&models.Show{}).Joins("Screen").Joins("Movie").Where("start_time>=? and theater_comp_refer_id= ?", cur_ts, theaterId).Find(screenShowSchedules); txn.Error != nil {
		return txn.Error
	}
	return nil
}
