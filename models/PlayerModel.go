package models

import (
	"errors"
	"fmt"
	"strconv"
	"tqgin/pkg/define"

	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	PlayerID    int64  `gorm:"not null;unique"`
	PlayerName  string `gorm:"not null"`
	DisPlayerID string `gorm:"unique"`
	Diamond     int64
	Gold        int64
	Cash        int64
	RoomID      int64
	Sex         define.SexType //1女，2男，0未知
	Sign        string         //签名
	Pic         string         //头像
	Loc         int
	Locx        float64
	Locy        float64
	Photos      string `gorm:size=1000`
	Rich        int64  //财富
	Charm       int64  //魅力
	CityName    string //城市
	BirthDay    string //出生日期
	StarSign    string //星座，由出生日期算出
	Profession  string //职业
	School      string
	MarryState  string //婚姻状态
}

func GetDefaultUserinfo(PlayerID int64, Name string, Sex define.SexType) UserInfo {
	var user UserInfo
	user.PlayerID = PlayerID
	user.PlayerName = Name
	user.DisPlayerID = strconv.FormatInt(PlayerID, 10)
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
func SaveUser(playerid int64, user UserInfo) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(UserInfo{}).Where("player_id = (?)", playerid).Update(&user).Error
	return err
}

//修改钻石数量，负数则是减
func ModifyDinamondUser(playerid int64, exDianmond int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(UserInfo{}).Where("player_id=(?) and diamond >= (?)", playerid, -exDianmond).UpdateColumn("diamond",
		gorm.Expr("diamond + ?", exDianmond)).Error

	if err == nil {
		if DB.RowsAffected != 0 {
			return err
		} else {
			return errors.New("无法满足执行条件")
		}
	} else {
		return err
	}
}

//修改金币数量，负数则是减
func ModifyGoldUser(playerid int64, exGold int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}
	err := DB.Model(UserInfo{}).Where("player_id=(?) and gold >= (?)", playerid, -exGold).UpdateColumn("gold",
		gorm.Expr("gold + ?", exGold)).Error

	if err == nil {
		if DB.RowsAffected != 0 {
			return err
		} else {
			return errors.New("无法满足执行条件")
		}
	} else {
		return err
	}
}

//修改现金数量，负数则是减
func ModifyCashUser(playerid int64, exCash int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}

	err := DB.Model(UserInfo{}).Where("player_id=(?) and cash >= (?)", playerid, -exCash).UpdateColumn("cash",
		gorm.Expr("cash + ?", exCash)).Error

	if err == nil {
		if DB.RowsAffected != 0 {
			return err
		} else {
			return errors.New("无法满足执行条件")
		}
	} else {
		return err
	}
}

//修改财富数量，负数则是减
func ModifyRichUser(playerid int64, exRich int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}

	err := DB.Model(UserInfo{}).Where("player_id=(?) and rich >= (?)", playerid, -exRich).UpdateColumn("rich",
		gorm.Expr("rich + ?", exRich)).Error

	if err == nil {
		if DB.RowsAffected != 0 {
			return err
		} else {
			return errors.New("无法满足执行条件")
		}
	} else {
		return err
	}
}

func ModifyCharmUser(playerid int64, exCharm int64) error {
	if playerid <= 0 {
		return errors.New("playerGUID error")
	}

	err := DB.Model(UserInfo{}).Where("player_id=(?) and charm >= (?)", playerid, -exCharm).UpdateColumn("charm",
		gorm.Expr("charm + ?", exCharm)).Error

	if err == nil {
		if DB.RowsAffected != 0 {
			return err
		} else {
			return errors.New("无法满足执行条件")
		}
	} else {
		return err
	}
}

//金币兑换钻石
func ExChangeGoldToDiamond(playerid int64, exGold, exDiamond int64) error {
	if playerid <= 0 || exGold < 0 || exDiamond < 0 {
		return errors.New("param error")
	}
	//先扣除金币在增再钻石
	err := DB.Model(UserInfo{}).Where("player_id=(?) and gold >= (?)",
		playerid, exGold).UpdateColumn("gold", gorm.Expr("gold - ?", exGold)).Error

	if err != nil {
		return err
	}
	if 0 == DB.RowsAffected {
		return errors.New("无法满足执行条件")
	}

	err = DB.Model(UserInfo{}).Where("player_id=(?)",
		playerid).UpdateColumn("diamond", gorm.Expr("diamond + ?", exDiamond)).Error
	if err != nil {
		return err
	}
	if 0 == DB.RowsAffected {
		return errors.New("无法满足执行条件")
	}
	return nil
}

func GetUser(playerID int64) *UserInfo {

	var user UserInfo

	DBtemp := DB.First(&user, "player_id = ?", playerID)
	if DBtemp.Error != nil {
		return nil
	}

	if DBtemp.RowsAffected <= 0 {
		return nil
	}

	return &user
}

func UserHasInfo(playerID int64) bool {

	var user UserInfo
	user.PlayerID = playerID

	var count int32
	DB.Model(UserInfo{}).Where(user).Select([]string{"player_id"}).Count(&count)

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
