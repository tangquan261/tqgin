package models

import (
	"errors"
	"time"

	//"log"
	"tqgin/pkg/util"

	"github.com/jinzhu/gorm"
)

type GiftCashType int32

const (
	GiftCashType_Diamond GiftCashType = 0
	GiftCashType_Gold    GiftCashType = 1
)

//送礼流水记录
type GifGiveRecord struct {
	gorm.Model
	GiftID     int64
	FromPlayer int64
	ToPlayer   int64
	Count      int32
	RoomID     int64
	md5String  string //送礼记录的流水md5
}

//获取到的礼物统计
type GfitUserCount struct {
	gorm.Model
	PlayerID int64
	GiftID   int64
	Count    int32
}

//礼物在人身上的数量，每种礼物获得多少个
type ConsumeUserCount struct {
	gorm.Model
	PlayerID int64
	GiftID   int64
	Count    int32
}

func GetGiftByID(giftID int64) GifInfo {
	return giftDic[giftID]
}

func AddGiveGiftLog(giftID int64, playerGUID int64, roomID int64, players []int64, nCount int32) error {

	if _, ok := giftDic[giftID]; !ok {
		return errors.New("not find gift")
	}

	md5String := util.EncodeMD5(util.DateTimeToString(time.Now()))

	for _, targetPlayer := range players {

		var giftRecord GifGiveRecord
		giftRecord.GiftID = giftID
		giftRecord.FromPlayer = playerGUID
		giftRecord.RoomID = roomID
		giftRecord.Count = nCount
		giftRecord.ToPlayer = targetPlayer
		giftRecord.md5String = md5String

		DB.Model(GifGiveRecord{}).Save(&giftRecord)
	}

	return nil
}

func AddConsumeUserCount(playerID, giftID int64, count int32) error {

	var giftCount ConsumeUserCount
	giftCount.PlayerID = playerID
	giftCount.GiftID = giftID
	giftCount.Count = count

	DBtemp := DB.Model(ConsumeUserCount{}).Where("player_id = (?) and gift_id = (?)", playerID, giftID).FirstOrCreate(&giftCount)

	if DBtemp.Error != nil {
		return errors.New("失败")
	}

	if DBtemp.RowsAffected == 0 {
		// 已经存在，则更新

		err := DB.Model(ConsumeUserCount{}).Where("player_id = (?) and gift_id = (?)", playerID, giftID).
			UpdateColumn("count", gorm.Expr("count + ?", count)).Error

		return err
	}

	return nil
}

func AddGiftUserCount(playerID, giftID int64, count int32) error {

	var giftCount GfitUserCount
	giftCount.PlayerID = playerID
	giftCount.GiftID = giftID
	giftCount.Count = count

	DBtemp := DB.Model(GfitUserCount{}).Where("player_id = (?) and gift_id = (?)", playerID, giftID).FirstOrCreate(&giftCount)

	if DBtemp.Error != nil {
		return errors.New("失败")
	}

	if DBtemp.RowsAffected == 0 {
		// 已经存在，则更新

		err := DB.Model(GfitUserCount{}).Where("player_id = (?) and gift_id = (?)", playerID, giftID).
			UpdateColumn("count", gorm.Expr("count + ?", count)).Error

		return err
	}

	return nil

}

//获取某个人的获取的礼物细腻,按照获取量降序排列返回
func GiftGetUserCount(playerID int64, limitCount int) []GfitUserCount {

	var ret []GfitUserCount
	if limitCount <= 0 {
		err := DB.Model(GfitUserCount{}).Where("player_id = (?)", playerID).Order("count desc").Find(&ret).Error

		if err != nil {
			return nil
		}

		return ret
	}

	err := DB.Model(GfitUserCount{}).Where("player_id = (?)", playerID).Limit(limitCount).Order("count desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}
