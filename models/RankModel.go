package models

import (
	"errors"
	"log"
	"strconv"
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
	gorm.Model
	PlayerID   int64
	Rich       int64
	Charm      int64
	Room       int64
	Star       int64
	Popularity int64
}

//
var (
	rankDic      map[string][]RankResult
	rankDataTime map[string]time.Time
)

func init() {
	rankDic = make(map[string][]RankResult)
	rankDataTime = make(map[string]time.Time)
}

func RankinfoSave(rank RankInfo) error {

	if rank.PlayerID <= 0 {
		return errors.New("playerid error")
	}

	now := time.Now()

	newRankInfo := rank

	tempDB := DB.Model(RankInfo{}).Where("player_id=(?) and day(created_at) = (day(?))",
		rank.PlayerID, now).FirstOrCreate(&newRankInfo)
	if tempDB.Error != nil {
		return errors.New("RankinfoSave db error")
	}

	if tempDB.RowsAffected == 1 {
		//新建数据生效，立即返回
		return nil
	}

	if rank.Rich > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and day(created_at) = (day(?))",
			rank.PlayerID, now).UpdateColumn("rich", gorm.Expr("rich + ?", rank.Rich)).Error

		if err != nil {
			log.Println("RankinfoSave rich", err)
		}
	}

	if rank.Charm > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and day(created_at) = (day(?))",
			rank.PlayerID, now).UpdateColumn("charm", gorm.Expr("charm + ?", rank.Charm)).Error

		if err != nil {
			log.Println("RankinfoSave Charm", err)
		}
	}

	if rank.Room > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and day(create_time) = (day(?))",
			rank.PlayerID, now).UpdateColumn("room", gorm.Expr("room + ?", rank.Room)).Error

		if err != nil {
			log.Println("RankinfoSave Room", err)
		}
	}
	if rank.Star > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and day(create_time) = (day(?))",
			rank.PlayerID, now).UpdateColumn("star", gorm.Expr("star + ?", rank.Star)).Error

		if err != nil {
			log.Println("RankinfoSave Star", err)
		}
	}
	if rank.Popularity > 0 {
		err := DB.Model(RankInfo{}).Where("player_id=(?) and day(create_time) = (day(?))",
			rank.PlayerID, now).UpdateColumn("popularity", gorm.Expr("popularity + ?", rank.Popularity)).Error

		if err != nil {
			log.Println("RankinfoSave Popularity", err)
		}
	}

	return nil
}

type RankResult struct {
	Player_id  int64
	Rich       int64
	Charm      int64
	Room       int64
	Star       int64
	Popularity int64
}

func RankInfoBy(rankType RankType, rankSubType RankSubType) []RankResult {

	key := strconv.Itoa(int(rankType)) + "_" + strconv.Itoa(int(rankSubType))

	if value, ok := rankDataTime[key]; ok {
		//5分钟刷去一次新
		if time.Now().Unix() < (value.Unix() + 300) {
			ret, _ := rankDic[key]
			return ret
		}
	}

	var rets []RankResult
	var err error
	if rankType == RankType_Rich { //财富
		if RankSubType_Day == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(rich) as rich").
				Where("day(create_time)=day(now())").Group("player_id").Order("sum(rich) desc").
				Limit(100).Scan(&rets).Error
		} else if RankSubType_Week == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(rich) as rich").
				Where("week(create_time)=week(now())").Group("player_id").Order("sum(rich) desc").
				Limit(100).Scan(&rets).Error

		} else {
			err = DB.Model(RankInfo{}).Select("player_id, sum(rich) as rich").
				Where("month(create_time)=month(now())").Group("player_id").Order("sum(rich) desc").
				Limit(100).Scan(&rets).Error
		}
	} else if rankType == RankType_Charm { //魅力
		if RankSubType_Day == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(charm) as charm").
				Where("day(create_time)=day(now())").Group("player_id").Order("sum(charm) desc").
				Limit(100).Scan(&rets).Error
		} else if RankSubType_Week == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(charm) as charm").
				Where("week(create_time)=week(now())").Group("player_id").Order("sum(charm) desc").
				Limit(100).Scan(&rets).Error

		} else {
			err = DB.Model(RankInfo{}).Select("player_id, sum(charm) as charm").
				Where("month(create_time)=month(now())").Group("player_id").Order("sum(charm) desc").
				Limit(100).Scan(&rets).Error
		}
	} else if rankType == RankType_Room { //房间
		if RankSubType_Day == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(room) as room").
				Where("day(create_time)=day(now())").Group("player_id").Order("sum(room) desc").
				Limit(100).Scan(&rets).Error
		} else if RankSubType_Week == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(room) as room").
				Where("week(create_time)=week(now())").Group("player_id").Order("sum(room) desc").
				Limit(100).Scan(&rets).Error

		} else {
			err = DB.Model(RankInfo{}).Select("player_id, sum(room) as room").
				Where("month(create_time)=month(now())").Group("player_id").Order("sum(room) desc").
				Limit(100).Scan(&rets).Error
		}
	} else if rankType == RankType_Star { //名人
		if RankSubType_Day == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(star) as star").
				Where("day(create_time)=day(now())").Group("player_id").Order("sum(star) desc").
				Limit(100).Scan(&rets).Error
		} else if RankSubType_Week == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(star) as star").
				Where("week(create_time)=week(now())").Group("player_id").Order("sum(star) desc").
				Limit(100).Scan(&rets).Error

		} else {
			err = DB.Model(RankInfo{}).Select("player_id, sum(star) as star").
				Where("month(create_time)=month(now())").Group("player_id").Order("sum(star) desc").
				Limit(100).Scan(&rets).Error
		}
	} else if rankType == RankType_Popularity { //声望
		if RankSubType_Day == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(popularity) as popularity").
				Where("day(create_time)=day(now())").Group("player_id").Order("sum(popularity) desc").
				Limit(100).Scan(&rets).Error
		} else if RankSubType_Week == rankSubType {
			err = DB.Model(RankInfo{}).Select("player_id, sum(popularity) as popularity").
				Where("week(create_time)=week(now())").Group("player_id").Order("sum(popularity) desc").
				Limit(100).Scan(&rets).Error

		} else {
			err = DB.Model(RankInfo{}).Select("player_id, sum(popularity) as popularity").
				Where("month(create_time)=month(now())").Group("player_id").Order("sum(popularity) desc").
				Limit(100).Scan(&rets).Error
		}
	} else {
		//do nothing
		err = errors.New("not find type")
	}
	if err != nil {
		return nil
	}

	if len(rets) > 0 {
		rankDataTime[key] = time.Now()
		rankDic[key] = rets
	}

	return rets
}
