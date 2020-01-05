/*
	获取配置信息
*/
package models

import (
	"fmt"
)

//礼物配置表
type GifInfoConfig struct {
	ID       int32
	GiftID   int32 `gorm:"primary_key"`
	GiftName string
	GiftIcon string
	CashType GiftCashType //消费类型
	CashNum  int32        //消费数量
}

type GiftInfoBatchConfig struct {
	GifInfoConfig
	Num     int32 //礼物数量
	World   int32 //礼物是否发世界
	Quality int32 //礼物质量
}

var giftDicBatch map[int32]*GiftInfoBatchConfig
var giftArray []*GifInfoConfig

func init() {
	giftDicBatch = make(map[int32]*GiftInfoBatchConfig)
}

func GetGiftmodel(giftID, giftNum int32) *GiftInfoBatchConfig {

	if obj, ok := giftDicBatch[giftID*10000+giftNum]; ok {
		return obj
	} else {
		GetAllGift()
		if obj, ok := giftDicBatch[giftID*10000+giftNum]; ok {
			return obj
		}
	}

	return nil
}

//礼物配置表
func GetAllGift() []*GifInfoConfig {

	if len(giftArray) > 0 {
		return giftArray
	}

	rows, _ := DB.Raw("select * from config_gift;").Rows()

	defer rows.Close()

	giftArray = giftArray[0:0]
	for rows.Next() {
		batch := new(GiftInfoBatchConfig)
		err := rows.Scan(&batch.ID, &batch.GiftID, &batch.GiftName, &batch.GiftIcon, &batch.CashType,
			&batch.CashNum, &batch.Num, &batch.World, &batch.Quality)
		if err == nil {
			giftDicBatch[batch.GiftID*10000+batch.Num] = batch
			if batch.Num == 1 {
				giftArray = append(giftArray, &batch.GifInfoConfig)
			}

		}
	}
	fmt.Println(giftArray)
	fmt.Println(giftDicBatch)
	return giftArray
}
