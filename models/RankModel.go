package models

import (
	"errors"
	"log"
	"time"

	//"time"

	//"tqgin/pkg/util"

	"github.com/jinzhu/gorm"
)

type RankType int32
type RankSubType int32

const (
	RankType_Rich       RankType = 0 //富豪
	RankType_Charm      RankType = 1 //魅力
	RankType_Room       RankType = 2 //房间
	RankType_Star       RankType = 3 //名人
	RankType_Popularity RankType = 4 //声望
)

const (
	RankSubType_Day   RankSubType = 0 //日
	RankSubType_Week  RankSubType = 1 //周
	RankSubType_Month RankSubType = 2 //月
)

type RankInfo struct {
	PlayerID   int64
	CreateTime time.Time
	Rich       int64
	Charm      int64
	Room       int64
	Star       int64
	Popularity int64
}

func RankinfoSave(rank RankInfo) error {

	if rank.PlayerID <= 0 {
		return errors.New("playerid error")
	}

	now := time.Now()
	rank.CreateTime = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)
	newRankInfo := rank

	tempDB := DB.Model(RankInfo{}).Where("player_id=(?) and create_time = (?)",
		rank.PlayerID, rank.CreateTime).FirstOrCreate(&newRankInfo)
	if tempDB.Error != nil {
		return errors.New("RankinfoSave db error")
	}

	if tempDB.RowsAffected == 1 {
		//新建数据生效，立即返回
		return nil
	}

	if rank.Rich > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and create_time = (?)",
			rank.PlayerID, rank.CreateTime).UpdateColumn("rich", gorm.Expr("rich + ?", rank.Rich)).Error

		if err != nil {
			log.Println("RankinfoSave rich", err)
		}
	}

	if rank.Charm > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and create_time = (?)",
			rank.PlayerID, rank.CreateTime).UpdateColumn("charm", gorm.Expr("charm + ?", rank.Charm)).Error

		if err != nil {
			log.Println("RankinfoSave Charm", err)
		}
	}

	if rank.Room > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and create_time = (?)",
			rank.PlayerID, rank.CreateTime).UpdateColumn("room", gorm.Expr("room + ?", rank.Room)).Error

		if err != nil {
			log.Println("RankinfoSave Room", err)
		}
	}
	if rank.Star > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and create_time = (?)",
			rank.PlayerID, rank.CreateTime).UpdateColumn("star", gorm.Expr("star + ?", rank.Star)).Error

		if err != nil {
			log.Println("RankinfoSave Star", err)
		}
	}
	if rank.Popularity > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and create_time = (?)",
			rank.PlayerID, rank.CreateTime).UpdateColumn("popularity", gorm.Expr("popularity + ?", rank.Popularity)).Error

		if err != nil {
			log.Println("RankinfoSave Popularity", err)
		}
	}

	return nil
}
