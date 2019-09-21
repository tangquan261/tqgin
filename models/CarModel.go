/*
	用户model
*/
package models

import (
	//"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type UserCar struct {
	gorm.Model
	Player_ID int64
	CarID     int32
	IsWear    int32
	EndTime   time.Time //结束时间
}

func AddCarinfo(playerID int64, car *CarConfig) bool {

	var userCar UserCar
	userCar.Player_ID = playerID
	userCar.CarID = car.CarID
	userCar.IsWear = 0
	userCar.EndTime = time.Now().Add(time.Duration(car.Duration) * time.Second)
	DBTemp := DB.Model(UserCar{}).Where("player_id = (?) and car_id = (?)", playerID, car.CarID).FirstOrCreate(&userCar)

	if DBTemp.Error != nil {
		return false
	} else if DBTemp.RowsAffected != 0 {
		return true
	} else {
		userCar.EndTime = userCar.EndTime.Add(time.Duration(car.Duration) * time.Second)
		DB.Model(UserCar{}).Where("player_id = (?) and car_id = (?)", playerID, car.CarID).UpdateColumn("end_time", userCar.EndTime)
	}
	return true

}

func LoadCars(playerID int64) []UserCar {
	var cars []UserCar

	err := DB.Model(UserCar{}).Where("player_id = (?) and end_time > (?)", playerID, time.Now()).Find(&cars).Error
	if err != nil {
		return nil
	}
	return cars
}

func loadWearCar(playerID int64) *UserCar {
	var onecar UserCar

	err := DB.Model(UserCar{}).Where("player_id = (?) and is_wear = (?) and end_time > (?)", playerID, 1, time.Now()).Find(&onecar).Error
	if err != nil {
		return nil
	}
	return &onecar
}

func ModifyWearCar(playerID int64, carID int32) {
	DB.Model(UserCar{}).Where("player_id = (?)", playerID).UpdateColumn("is_wear", 0)
	if carID > 0 {
		DB.Model(UserCar{}).Where("player_id = (?) and car_id = (?)", playerID, carID).UpdateColumn("is_wear", 1)
	}
}
