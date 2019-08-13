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

var gifts []GifInfo
var giftDic map[int64]GifInfo

func init() {
	giftDic = make(map[int64]GifInfo)
}

type GifInfo struct {
	GiftID   int64 `gorm:"primary_key"`
	GiftName string
	GiftIcon string
	CashType GiftCashType //消费类型
	CashNum  int32        //消费数量
}

type GifGiveRecord struct {
	gorm.Model
	GiftID     int64
	FromPlayer int64
	ToPlayer   int64
	Count      int32
	RoomID     int64
	md5String  string //送礼记录的流水md5
}

func GetAllGift() []GifInfo {

	if len(gifts) > 0 {
		return gifts
	}
	err := DB.Model(GifInfo{}).Find(&gifts).Error
	if err != nil {
		return nil
	}

	for _, gif := range gifts {
		giftDic[gif.GiftID] = gif
	}

	return gifts
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
