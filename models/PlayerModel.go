package models

import (
	"errors"
	"fmt"
	"time"

	//"time"
	"tqgin/proto"

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
	Sex        login.SexType //1女，2男，0未知
	BirthDay   time.Time
	Sign       string
	Pic        string
	Loc        int
	Locx       float64
	Locy       float64
	Photos     string `gorm:size=1000`
}

func GetDefaultUserinfo(PlayerID int64, Name string, Sex login.SexType) UserInfo {
	var user UserInfo
	user.PlayerID = PlayerID
	user.PlayerName = Name
	user.Sex = Sex

	return user
}

func CreateUser(user *UserInfo) error {

	if user.PlayerID <= 0 || user.PlayerName == "" {
		return errors.New("param err")
	}

	return DB.Create(user).Error
}

func SaveUser(user *UserInfo) error {
	if user.PlayerID <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(user).Update(user).Error

	return err
}

func GetUser(playerID int64) *UserInfo {

	var user UserInfo
	user.PlayerID = playerID

	err := DB.First(&user, "player_id = ?", playerID).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &user
}

func GetUsers(playerIDs []int64) []UserInfo {

	if len(playerIDs) <= 0 {
		return nil
	}
	var user []UserInfo

	err := DB.Where("player_id in (?)", playerIDs).Find(&user)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}
