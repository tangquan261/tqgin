package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//关注好友关系
type RelationShip struct {
	gorm.Model
	PlayerID    int64 `gorm:"not null"` //自己的id
	TarplayerID int64 `gorm:"not null"` //关注者的id
}

//黑名单
type Black struct {
	gorm.Model
	PlayerID int64 `gorm:"not null"`
	BlackID  int64 `gorm:"not null"`
}

func RelationIsFans(myID, tarID int64) bool {
	if myID <= 0 || tarID <= 0 {
		return false
	}

	var user RelationShip

	notfind := DB.Where("player_id = (?) and tarplayer_id = (?)", tarID, myID).Find(&user).RecordNotFound()
	if notfind {
		return false
	}
	return true
}

func RelationAddFollow(fromID, tarID int64) error {
	if fromID <= 0 || tarID <= 0 {
		return errors.New("参数错误")
	}

	var relation RelationShip
	relation.PlayerID = fromID
	relation.TarplayerID = tarID

	tempDB := DB.Model(RelationShip{}).Where("player_id = (?) and tarplayer_id=(?)",
		fromID, tarID).FirstOrCreate(&relation)

	if tempDB.Error != nil {
		return errors.New("保存错误")
	}

	if tempDB.RowsAffected == 1 {
		//新建数据生效，立即返回
		return nil
	} else {
		return errors.New("已经关注过了")
	}
}

//获取我关注的人
func GetFollow(playerID int64) []RelationShip {

	if playerID <= 0 {
		return nil
	}
	var user RelationShip
	user.PlayerID = playerID
	var users []RelationShip

	err := DB.Where(user).Select([]string{"tarplayer_id"}).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func GetFollowCount(playerID int64) int {
	var count int

	err := DB.Raw("select count(1) as total from tq_relation_ship where player_id = ? ", playerID).Count(&count).Error

	if err != nil {
		return 0
	}

	return count
}

func GetFansCount(playerID int64) int {
	var count int

	err := DB.Raw("select count(1) as total from tq_relation_ship where tarplayer_id = ? ", playerID).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

//获取关注我的人，即我的粉丝
func GetFans(playerID int64) []RelationShip {
	if playerID <= 0 {
		return nil
	}
	var user RelationShip
	user.TarplayerID = playerID
	var users []RelationShip

	err := DB.Where(user).Select([]string{"tarplayer_id"}).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func GetFirend(playerID int64) []RelationShip {
	if playerID <= 0 {
		return nil
	}

	//Joins("left join emails on emails.user_id = users.id").Scan(&results)

	var users []RelationShip
	err := DB.Model(RelationShip{}).Select("player_id").
		Joins("inner join (select tarplayer_id as p2 from tq_relation_ship WHERE player_id = (?)) t2 on tq_relation_ship.player_id = t2.p2 ", playerID).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func RelationDelFollow(fromID, tarID int64) error {
	if fromID <= 0 || tarID <= 0 {
		return errors.New("参数错误")
	}

	var relation RelationShip
	relation.PlayerID = fromID
	relation.TarplayerID = tarID

	err := DB.Unscoped().Where("player_id = (?) and tarplayer_id=(?)",
		fromID, tarID).Delete(RelationShip{}).Error
	if err != nil {
		return errors.New("数据错误")
	} else {
		return nil
	}
}

func GetBlacks(playerID int64) []Black {
	if playerID <= 0 {
		return nil
	}
	var user Black
	user.PlayerID = playerID
	var users []Black

	err := DB.Where(user).Select([]string{"player_id", "black_id"}).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func AddBlack(playerID, blackID int64) error {
	if playerID <= 0 || blackID <= 0 {
		return errors.New("参数错误")
	}

	if playerID == blackID {
		return errors.New("参数错误")
	}

	hasUser := UserHasInfo(blackID)
	if !hasUser {
		return errors.New("参数错误")
	}

	var black Black
	black.PlayerID = playerID
	black.BlackID = blackID
	err := DB.Where(&black).FirstOrCreate(&black).Error
	return err
}

func RemoveBlack(playerID, blackID int64) error {
	if playerID <= 0 || blackID <= 0 {
		return errors.New("参数错误")
	}

	if playerID == blackID {
		return errors.New("参数错误")
	}

	hasUser := UserHasInfo(blackID)
	if !hasUser {
		return errors.New("参数错误")
	}
	var balck Black
	balck.PlayerID = playerID
	balck.BlackID = blackID

	return DB.Where(balck).Delete(balck).Error
}
