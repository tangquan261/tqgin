//房间排行榜model
package models

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type RoomRankType int32

const (
	RoomRankType_Rich  RoomRankType = 0 //财富
	RoomRankType_Charm RoomRankType = 1 //魅力
)

type RoomRankInfo struct {
	gorm.Model
	RoomID   int64
	PlayerID int64
	Rich     int64
	Charm    int64
}

func RoomRankinfoSave(rank RoomRankInfo) error {

	if rank.PlayerID <= 0 {
		return errors.New("playerid error")
	}

	now := time.Now()

	newRankInfo := rank

	tempDB := DB.Model(RoomRankInfo{}).Where("room_id = (?) and player_id=(?) and week(created_at) = (week(?))",
		rank.RoomID, rank.PlayerID, now).FirstOrCreate(&newRankInfo)
	if tempDB.Error != nil {
		return errors.New("RoomRankinfoSave db error")
	}

	if tempDB.RowsAffected == 1 {
		//新建数据生效，立即返回
		return nil
	}

	if rank.Rich > 0 {
		err := DB.Model(RoomRankInfo{}).Where("room_id = (?) and player_id=(?) and week(create_time) = (week(?))",
			rank.RoomID, rank.PlayerID, now).UpdateColumn("rich", gorm.Expr("rich + ?", rank.Rich)).Error

		if err != nil {
			log.Println("RoomRankinfoSave rich", err)
		}
	}
	if rank.Charm > 0 {
		err := DB.Model(RoomRankInfo{}).Where("room_id = (?) and player_id=(?) and week(create_time) = (week(?))",
			rank.RoomID, rank.PlayerID, now).UpdateColumn("charm", gorm.Expr("charm + ?", rank.Charm)).Error

		if err != nil {
			log.Println("RoomRankinfoSave charm", err)
		}
	}

	return nil
}

type RoomRankResult struct {
	Player_id  int64
	Star       int64
	Popularity int64
}

func RoomRankInfoBy(roomID int64, rankType RoomRankType) ([]RoomRankResult, error) {

	var rets []RoomRankResult
	var err error
	if rankType == RoomRankType_Rich { //财富
		err = DB.Model(RankInfo{}).Select("room_id = (?) and player_id, sum(rich) as rich", roomID).
			Where("week(create_time)=week(now())").Group("player_id").Order("sum(rich) desc").
			Limit(100).Scan(&rets).Error
	} else if rankType == RoomRankType_Charm { //魅力
		err = DB.Model(RankInfo{}).Select("room_id = (?) and player_id, sum(charm) as charm", roomID).
			Where("week(create_time)=week(now())").Group("player_id").Order("sum(charm) desc").
			Limit(100).Scan(&rets).Error
	} else {
		//do nothing
		err = errors.New("not find type")
	}

	return rets, err
}
