package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*
FType     int32    `json:"ftype"`     //1,2声音，普通
	SoundRUL  string   `json:"soundurl"`  //声音地址
	PhotoURLs []string `json:"photourl"`  //图片地址列表
	Content      string   `json:"content"`      //文本
	Ats       []string `json:"at"`        //at的人列表
	LocX      int64    `json:"locx"`      //x位置
	LocY      int64    `json:"locy"`      //y位置
	LocString string    `json:"locstring"` //位置名称
*/

type CycleModel struct {
	gorm.Model
	Cid       string //客户端uid
	SnowID    int64  //服务器生成自增id
	PlayerID  int64  //发布者id
	FType     int32  //发布类型，1，2，3普通，声音，视频
	SoundBG   string //声音背景图
	SoundRUL  string //声音地址
	PhotoURLs string //图片地址列表
	Content   string //内容
	Ats       string //at的人列表
	LocX      int64  //经度
	LocY      int64  //纬度
	LocString string //位置字符串
}

//根据帖子id获取帖子信息
func CycleGetModel(snowID int64) *CycleModel {
	if snowID <= 0 {
		return nil
	}

	var ret CycleModel
	err := DB.Model(CycleModel{}).Where("snow_id = (?)", snowID).Find(&ret).Error
	if err != nil {
		return nil
	}
	return &ret
}

//添加帖子
func CycleAdd(cycle CycleModel) error {

	if len(cycle.Cid) <= 0 {
		return errors.New("参数错误")
	}

	tempDB := DB.Model(CycleModel{}).Where("cid = (?)", cycle.Cid).FirstOrCreate(&cycle)

	if tempDB.Error != nil {
		return errors.New("参数错误")
	}

	if tempDB.RowsAffected == 1 {
		return nil

	} else {
		return errors.New("慢点发哦！")
	}

}

//移除帖子
func CycleDel(snowID int64) error {
	if snowID < 0 {
		return errors.New("参数错误")
	}
	err := DB.Where("snow_id = (?)", snowID).Delete(CycleModel{}).Error

	return err
}

//获取声音的帖子
func CycleGetSound(snowID int64) []CycleModel {

	var ret []CycleModel

	if snowID <= 0 {
		DB.Model(CycleModel{}).Where("f_type = 2").Limit(20).Order("id desc").Find(&ret)
		return ret
	}

	DB.Model(CycleModel{}).Where("f_type = 2 and snow_id < (?)", snowID).Limit(20).Order("id desc").Find(&ret)

	return ret
}

//获取视频的帖子
func CycleGetAudio(snowID int64) []CycleModel {

	var ret []CycleModel

	if snowID <= 0 {
		DB.Model(CycleModel{}).Where("f_type = 3").Limit(20).Order("id desc").Find(&ret)
		return ret
	}

	DB.Model(CycleModel{}).Where("f_type = 3 and snow_id < (?)", snowID).Limit(20).Order("id desc").Find(&ret)

	return ret
}

//获取朋友圈的帖子列表
func CycleGetFeeds(snowID int64) []CycleModel {

	var ret []CycleModel

	if snowID <= 0 {
		DB.Model(CycleModel{}).Limit(20).Order("id desc").Find(&ret)
		return ret
	}

	DB.Model(CycleModel{}).Where("snow_id < (?)", snowID).Limit(20).Order("id desc").Find(&ret)

	return ret
}

// //组合查询，暂时不用
// func CycleTestFeed(snowID int64) []CycleModel {

// 	var ret []CycleModel
// 	if snowID <= 0 {
// 		DB.Model(CycleModel{}).Limit(20).Order("id desc").Find(&ret)
// 		return ret
// 	}

// 	err := DB.Model(CycleModel{}).
// 		Joins("inner join(select id from tq_cycle_model where uuid = (?)) t2 on tq_cycle_model.id < t2.id ",
// 			index).Limit(20).Order("id desc").Find(&ret).Error

// 	if err != nil {
// 		return nil
// 	}

// 	return ret
// }

//我关注的人动态
func CycleGetFeedsFollow(playerID int64, snowID int64) []CycleModel {
	var ret []CycleModel

	if snowID <= 0 {
		err := DB.Model(CycleModel{}).
			Where("`tq_cycle_model`.player_id in (SELECT tarplayer_id FROM tq_relation_ship WHERE  player_id = (?))",
				playerID).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}

		return ret
	}

	err := DB.Model(CycleModel{}).
		Where("`tq_cycle_model`.player_id in (SELECT tarplayer_id FROM tq_relation_ship WHERE  player_id = (?))",
			playerID).Where("snow_id < (?)", snowID).Limit(20).Order("id desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}

//我的粉丝的人动态
func CycleGetFeedsFans(playerID int64, snowID int64) []CycleModel {
	var ret []CycleModel

	if snowID <= 0 {
		err := DB.Model(CycleModel{}).
			Where("`tq_cycle_model`.player_id in (SELECT player_id FROM tq_relation_ship WHERE  tarplayer_id = (?))",
				playerID).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}

		return ret
	}

	err := DB.Model(CycleModel{}).
		Where("`tq_cycle_model`.player_id in (SELECT player_id FROM tq_relation_ship WHERE  tarplayer_id = (?))",
			playerID).Where("snow_id < (?)", snowID).Limit(20).Order("id desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}

//获取某个人的动态
func CycleGetSingleFeeds(playerID int64, snowID int64, limitCount int) []CycleModel {

	if limitCount <= 0 || limitCount > 30 {
		return nil
	}

	var ret []CycleModel

	if snowID <= 0 {
		err := DB.Model(CycleModel{}).Where("player_id = (?)", playerID).
			Limit(limitCount).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}

		return ret
	}

	err := DB.Model(CycleModel{}).Where("player_id = (?) and snow_id < (?)", playerID, snowID).Limit(limitCount).Order("id desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}
