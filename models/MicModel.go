package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type MicModel struct {
	gorm.Model
	RoomID   int64
	MicIndex int16
	PlayerID int64
}

//添加到mic
func MicAdd(roomID, playerGUID int64, micIndex int16) error {

	var mic MicModel
	mic.RoomID = roomID
	mic.PlayerID = playerGUID
	mic.MicIndex = micIndex

	err := DB.Model(MicModel{}).Where("room_id = (?) and mic_index = (?)", roomID, micIndex).FirstOrCreate(&mic).Error
	if err != nil {
		log.Println("mic add  error", err)
	}
	return err
}

func MicDelByPlayerID(playerGUID int64) error {
	err := DB.Unscoped().Where("player_id = (?)", playerGUID).Delete(MicModel{}).Error
	if err != nil {
		log.Println("MicDelByPlayerID error ", err)
	}
	return err
}

func MicDelByDismiss(roomID int64) error {
	err := DB.Where("room_id = (?)", roomID).Delete(MicModel{}).Error
	if err != nil {
		log.Println("MicDelByDismiss error ", err)
	}
	return err
}

func MicGetAllIndex(roomID int64) []MicModel {

	var mics []MicModel

	DB.Model(MicModel{}).Where("room_id = (?)", roomID).Find(&mics)

	return mics
}
