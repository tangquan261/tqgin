/*
	用户model
*/
package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type UserTitle struct {
	gorm.Model
	Player_ID int64
	Title_ID  int32
	IsWear    int32
	EndTime   time.Time //结束时间
}

func TitlesGetByPlayerID(playerID int64) []UserTitle {

	var models []UserTitle
	DBTemp := DB.Model(UserTitle{}).Where("player_id = (?) and end_time > (?)", playerID, time.Now()).Find(&models)

	if DBTemp.Error != nil {
		return nil
	}
	return models
}

func TitleSave(playerID int64, config *TitleConfig) bool {
	if playerID <= 0 || config == nil {
		return false
	}

	var model UserTitle
	model.Player_ID = playerID
	model.IsWear = 0
	model.Title_ID = config.ID
	model.EndTime = time.Now().Add(time.Duration(config.Duration) * time.Second)

	DBtemp := DB.Model(UserTitle{}).Where("player_id = (?) and title_id = (?)",
		playerID, config.ID).FirstOrCreate(&model)
	if DBtemp.Error != nil {
		return false
	} else if DBtemp.RowsAffected != 0 {
		//加入成功
		return true
	} else {
		//已经存在，尝试更新时间
		fmt.Println("----------", model.EndTime)
		model.EndTime = model.EndTime.Add(time.Duration(config.Duration) * time.Second)

		fmt.Println("----------", model.EndTime)
		DB.Model(UserTitle{}).Where("player_id = (?) and title_id = (?)",
			playerID, config.ID).Update(&model)
		return true
	}
}

func ModifyWearTitle(playerID int64, titleID int32) {
	DB.Model(UserTitle{}).Where("player_id = (?)", playerID).UpdateColumn("is_wear", 0)
	if titleID > 0 {
		DB.Model(UserTitle{}).Where("player_id = (?) and title_id = (?)", playerID, titleID).UpdateColumn("is_wear", 1)
	}
}

//id, name, icon,iconanimation
