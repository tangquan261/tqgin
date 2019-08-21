/*
	获取配置信息
*/
package models

//礼物配置表
type GifInfo struct {
	GiftID   int64 `gorm:"primary_key"`
	GiftName string
	GiftIcon string
	CashType GiftCashType //消费类型
	CashNum  int32        //消费数量
}

var gifts []GifInfo
var giftDic map[int64]GifInfo

func init() {
	giftDic = make(map[int64]GifInfo)
}

//礼物配置表
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
