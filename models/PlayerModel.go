package models

import (
	"fmt"

	//"time"

	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	PlayerID   int64  `gorm:"not null;unique"`
	PlayerName string `gorm:"not null;unique"`
	Diamond    int64
	Gold       int64
	Cash       int64
	RoomID     int64
	Sex        int8 //1女，2男，0未知
}

func GetDefaultUserinfo(PlayerID int64, Name string, Sex int8) UserInfo {

	var user UserInfo
	user.PlayerID = PlayerID
	user.PlayerName = Name
	user.Sex = Sex

	return user

}

func CreateUser(user *UserInfo) bool {

	if user.PlayerID <= 0 || user.PlayerName == "" {
		return false
	}

	err := DB.Create(user).GetErrors()

	if len(err) > 0 {
		return false
	} else {
		return true
	}
}

func SaveUser(user *UserInfo, newUser *UserInfo) bool {

	if user.PlayerID == 0 || user.PlayerID != newUser.PlayerID {
		return false
	}

	err := DB.Model(user).Update(newUser).GetErrors()

	if len(err) > 0 {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
