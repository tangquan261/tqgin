package models

import (
	"errors"
	"fmt"
	"time"
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

//不累计的属性存储,不能用户货币类型的更新
func SaveUser(playerid int64, user *UserInfo) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(UserInfo{}).Where("player_id = (?)", playerid).Update(user).Error
	return err
}

//负数则是减
func ModifyDinamondUser(playerid int64, exDianmond int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(UserInfo{}).Where("player_id=(?)", playerid).UpdateColumn("diamond",
		gorm.Expr("diamond + ?", exDianmond)).Error

	return err
}

func ModifyGoldUser(playerid int64, exGold int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(UserInfo{}).Where("player_id=(?)", playerid).UpdateColumn("gold",
		gorm.Expr("gold + ?", exGold)).Error

	return err
}

func ModifyCashUser(playerid int64, exCash int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}

	err := DB.Model(UserInfo{}).Where("player_id=(?)", playerid).UpdateColumn("cash",
		gorm.Expr("cash + ?", exCash)).Error

	return err
}

func ExChangeGoldToDiamond(playerid int64, exGold, exDiamond int64) error {
	if playerid <= 0 || exGold < 0 || exDiamond < 0 {
		return errors.New("param error")
	}

	affected := DB.Model(UserInfo{}).Where("player_id=(?) and gold >= (?)",
		playerid, exGold).UpdateColumn("gold", gorm.Expr("gold - ?", exGold)).RowsAffected

	var err error
	if affected == 1 {
		err = DB.Model(UserInfo{}).Where("player_id=(?)",
			playerid).UpdateColumn("diamond", gorm.Expr("diamond + ?", exDiamond)).Error
	}

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

func UserHasInfo(playerID int64) bool {

	var user UserInfo
	user.PlayerID = playerID

	var count int32
	DB.Model(UserInfo{}).Where(user).Count(&count)

	return count != 0
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
