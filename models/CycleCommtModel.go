package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type CycleCommet struct {
	gorm.Model
	PlayerID    int64  //评论者id
	SnowID      int64  //帖子snowid
	TarPlayerID int64  //被评论者id，1级评论是0
	FromID      int64  //评论id，1级评论是0 2级评论是CycleCommet 对一个id
	Conent      string //评论内容
}

//帖子点赞，评论点赞
type CycleLike struct {
	gorm.Model
	SnowID   int64 //帖子id
	UID      int64 //评论id，0表示给帖子点赞
	PlayerID int64 //点赞者
}

//给帖子添加评论
func CycleAddCommet(cycle CycleCommet) error {

	if cycle.SnowID <= 0 || len(cycle.Conent) <= 0 {
		return errors.New("参数错误")
	}

	if cycle.FromID > 0 && cycle.PlayerID > 0 {
		//是二级评论
		//找到对应的评论，如果不是1级评论，则返回错误
		var preCycle CycleCommet
		err := DB.Model(CycleCommet{}).Where("id = (?)", cycle.FromID).Find(&preCycle).Error

		if err != nil {
			return errors.New("参数错误")
		}

		if preCycle.FromID > 0 {
			return errors.New("不能多级评论")
		}
	} else {
		//以及评论
		cycle.FromID = 0
		cycle.PlayerID = 0
	}

	err := DB.Model(CycleCommet{}).Save(&cycle).Error

	if err != nil {
		return errors.New("参数错误")
	}

	return nil
}

//获取该帖子的所有评论
func CycleGetCommet(snowID int64) []CycleCommet {
	if snowID <= 0 {
		return nil
	}
	var ret []CycleCommet
	err := DB.Model(CycleCommet{}).Where("snow_id = (?)", snowID).Where(&ret).Error

	if err != nil {
		return nil
	}
	return ret
}

//根据评论id删除评论
func CycleDelCommet(uid int64) error {
	if uid <= 0 {
		return errors.New("参数错误")
	}

	tx := DB.Begin()

	var commet CycleCommet
	err := tx.Model(CycleCommet{}).Where("id = (?)", uid).Find(&commet).Error

	if err != nil {
		return errors.New("参数错误")
	}

	CycleDelCommetByFromid(tx, commet.ID)

	//移除该评论下的点赞
	CycleDelCommetByCycleUID(tx, commet.ID)
	err = tx.Where("id = (?)", uid).Delete(CycleCommet{}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("参数错误")
	}

	tx.Commit()
	return nil
}

//根据from_id 移除 评论，以及该评论下的所有评论
func CycleDelCommetByFromid(tx *gorm.DB, fromID int64) {

	if fromID <= 0 {
		return
	}

	err := DB.Where("`tq_cycle_like`.uid in (select id from tq_cycle_commet where from_id = (?))", fromID).Delete(CycleLike{}).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = DB.Where("from_id = (?)", fromID).Delete(CycleCommet{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
}

//添加帖子或者评论的点赞
func CycleAddLikeCommet(like CycleLike) error {

	if like.SnowID <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Model(CycleLike{}).Where("snow_id= (?) and uid = (?) and player_id = (?)",
		like.SnowID, like.UID, like.PlayerID).FirstOrCreate(&like).Error

	if err != nil {
		return errors.New("参数错误")
	}

	return nil
}

//点赞查询数量 根据帖子id和评论id、 uid=0表示查询帖子点赞数，不为零表示评论的点赞数
func CycleGetLikeCommet(snowID int64, uid int64) int64 {

	if snowID <= 0 {
		return 0
	}

	if uid <= 0 {
		//查询该动态下所有点赞和评论的点赞
		var count int64
		err := DB.Model(CycleLike{}).Where("snow_id = (?)", snowID).Count(&count).Error
		if err != nil {
			return 0
		}
		return count

	} else {
		//查询该动态下对应评论的点赞
		var count int64
		err := DB.Model(CycleLike{}).Where("id = (?) and snow_id = (?)", uid, snowID).Count(&count).Error
		if err != nil {
			return 0
		}
		return count
	}
}

//移除帖子或者评论的点赞
func CycleDelLikeCommet(id int64) error {

	if id <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Where("id = (?)", id).Delete(CycleLike{}).Error
	if err != nil {
		return errors.New("参数错误")
	}
	return nil
}

//根据uid移除评论对应的点赞
func CycleDelCommetByCycleUID(tx *gorm.DB, uid int64) error {

	if uid <= 0 {
		return errors.New("参数错误")
	}

	err := tx.Where("uid = (?)", uid).Delete(CycleLike{}).Error
	if err != nil {
		return errors.New("参数错误")
	} else {
		return nil
	}
}

//根据uuid删除所有帖子的评论
func CycleDelCommetByCycleuuid(snowID int64) error {

	if snowID <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Where("snow_id = (?)", snowID).Delete(CycleCommet{}).Error
	if err != nil {
		return errors.New("参数错误")
	} else {
		return nil
	}
}

//根据uuid删除所有帖子的点赞
func CycleDelLikesByCycleuuid(snowID int64) error {
	if snowID <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Where("snow_id = (?)", snowID).Delete(CycleLike{}).Error
	if err != nil {
		return errors.New("参数错误")
	} else {
		return nil
	}
}
