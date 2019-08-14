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
	Cid       string
	Uuid      string
	PlayerID  int64
	FType     int32
	SoundRUL  string
	PhotoURLs string
	Content   string
	Ats       string
	LocX      int64
	LocY      int64
	LocString string
}

//根据帖子id获取帖子信息
func CycleGetModel(uuid string) *CycleModel {
	if len(uuid) <= 0 {
		return nil
	}

	var ret CycleModel
	err := DB.Model(CycleModel{}).Where("uuid = (?)", uuid).Find(&ret).Error
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
func CycleDel(uuid string) error {
	if len(uuid) < 0 {
		return errors.New("参数错误")
	}
	err := DB.Where("uuid = (?)", uuid).Delete(CycleModel{}).Error

	return err
}

//获取声音的帖子
func CycleGetSound(index string) []CycleModel {

	var ret []CycleModel

	if len(index) <= 0 {
		DB.Model(CycleModel{}).Where("f_type = 1").Limit(20).Order("id desc").Find(&ret)
		return ret
	}

	var indexdata CycleModel

	err := DB.Model(CycleModel{}).Select("id").Where("uuid = (?)", index).Unscoped().Find(&indexdata).Error

	if err != nil {
		//根据没有找到对应数据，则拉取最新的20条数据
		err = DB.Model(CycleModel{}).Where("f_type = 1").Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}
		return ret
	}

	DB.Model(CycleModel{}).Where("f_type = 1 and id < (?)", indexdata.ID).Limit(20).Order("id desc").Find(&ret)

	return ret
}

//获取朋友圈的帖子列表
func CycleGetFeeds(index string) []CycleModel {
	var ret []CycleModel

	if len(index) <= 0 {
		DB.Model(CycleModel{}).Limit(20).Order("id desc").Find(&ret)
		return ret
	}

	var indexdata CycleModel

	err := DB.Model(CycleModel{}).Select("id").Where("uuid = (?)", index).Unscoped().Find(&indexdata).Error

	if err != nil {
		//根据没有找到对应数据，则拉取最新的20条数据
		err = DB.Model(CycleModel{}).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}
		return ret
	}

	DB.Model(CycleModel{}).Where("id < (?)", indexdata.ID).Limit(20).Order("id desc").Find(&ret)

	return ret
}

//组合查询，暂时不用
func CycleTestFeed(index string) []CycleModel {

	var ret []CycleModel
	if len(index) <= 0 {
		DB.Model(CycleModel{}).Limit(20).Order("id desc").Find(&ret)
		return ret
	}

	err := DB.Model(CycleModel{}).
		Joins("inner join(select id from tq_cycle_model where uuid = (?)) t2 on tq_cycle_model.id < t2.id ",
			index).Limit(20).Order("id desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}

//我关注的人动态
func CycleGetFeedsFollow(playerID int64, index string) []CycleModel {
	var ret []CycleModel

	if len(index) <= 0 {
		err := DB.Model(CycleModel{}).
			Where("`tq_cycle_model`.player_id in (SELECT tarplayer_id FROM tq_relation_ship WHERE  player_id = (?))",
				playerID).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}

		return ret
	}

	var indexdata CycleModel

	err := DB.Model(CycleModel{}).Select("id").Where("uuid = (?)", index).Unscoped().Find(&indexdata).Error

	if err != nil {
		//根据没有找到对应数据，则拉取最新的20条数据
		err := DB.Model(CycleModel{}).
			Where("`tq_cycle_model`.player_id in (SELECT tarplayer_id FROM tq_relation_ship WHERE  player_id = (?))",
				playerID).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}
		return ret
	}

	err = DB.Model(CycleModel{}).
		Where("`tq_cycle_model`.player_id in (SELECT tarplayer_id FROM tq_relation_ship WHERE  player_id = (?))",
			playerID).Where("id < (?)", indexdata.ID).Limit(20).Order("id desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}

//我的粉丝的人动态
func CycleGetFeedsFans(playerID int64, index string) []CycleModel {
	var ret []CycleModel

	if len(index) <= 0 {
		err := DB.Model(CycleModel{}).
			Where("`tq_cycle_model`.player_id in (SELECT player_id FROM tq_relation_ship WHERE  tarplayer_id = (?))",
				playerID).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}

		return ret
	}

	var indexdata CycleModel

	err := DB.Model(CycleModel{}).Select("id").Where("uuid = (?)", index).Unscoped().Find(&indexdata).Error

	if err != nil {
		//根据没有找到对应数据，则拉取最新的20条数据
		err := DB.Model(CycleModel{}).
			Where("`tq_cycle_model`.player_id in (SELECT player_id FROM tq_relation_ship WHERE  tarplayer_id = (?))",
				playerID).Limit(20).Order("id desc").Find(&ret).Error

		if err != nil {
			return nil
		}
		return ret
	}

	err = DB.Model(CycleModel{}).
		Where("`tq_cycle_model`.player_id in (SELECT player_id FROM tq_relation_ship WHERE  tarplayer_id = (?))",
			playerID).Where("id < (?)", indexdata.ID).Limit(20).Order("id desc").Find(&ret).Error

	if err != nil {
		return nil
	}

	return ret
}
