package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type CycleCommet struct {
	gorm.Model
	PlayerID int64  //评论者id
	Uuid     string //帖子id
	FromID   int64  //评论id，1级评论是0
	Conent   string //评论内容
}

//帖子点赞，评论点赞
type CycleLike struct {
	gorm.Model
	Uuid     string //帖子id
	UID      int64  //评论id，0表示给帖子点赞
	PlayerID int64  //点赞者
}

//给帖子添加评论
func CycleAddCommet(cycle CycleCommet) error {

	if len(cycle.Uuid) <= 0 || len(cycle.Conent) <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Model(CycleCommet{}).Save(&cycle).Error

	if err != nil {
		return errors.New("参数错误")
	}

	return nil
}

//获取该帖子的所有评论
func CycleGetCommet(uuid string) []CycleCommet {
	if len(uuid) <= 0 {
		return nil
	}
	var ret []CycleCommet
	err := DB.Model(CycleCommet{}).Where("uuid = (?)", uuid).Where(&ret).Error

	if err != nil {
		return nil
	}
	return ret
}

//根据评论id删除评论，只删除当前评论，并没有删除二级，三级评论
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

	var commet []CycleCommet
	err := DB.Model(CycleCommet{}).Where("from_id = (?)", fromID).Find(&commet).Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, obj := range commet {
		CycleDelCommetByFromid(tx, obj.ID)
	}

	err = DB.Where("from_id = (?)", fromID).Delete(CycleCommet{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	//移除该评论下的点赞
	CycleDelCommetByCycleUID(tx, fromID)
}

//添加帖子或者评论的点赞
func CycleAddLikeCommet(like CycleLike) error {

	if len(like.Uuid) <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Model(CycleLike{}).Where("uuid= (?) and uid = (?) and player_id = (?)",
		like.Uuid, like.UID, like.PlayerID).FirstOrCreate(&like).Error

	if err != nil {
		return errors.New("参数错误")
	}

	return nil
}

//点赞查询数量 根据帖子id和评论id、 uid=0表示查询帖子点赞数，不为零表示评论的点赞数
func CycleGetLikeCommet(uuid string, uid int64) int64 {

	if len(uuid) <= 0 {
		return 0
	}

	if uid <= 0 {
		//查询该动态下所有点赞和评论的点赞
		var count int64
		err := DB.Model(CycleLike{}).Where("uuid = (?)", uuid).Count(&count).Error
		if err != nil {
			return 0
		}
		return count

	} else {
		//查询该动态下对应评论的点赞
		var count int64
		err := DB.Model(CycleLike{}).Where("id = (?) and uuid = (?)", uid, uuid).Count(&count).Error
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
func CycleDelCommetByCycleuuid(uuid string) error {

	if len(uuid) <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Where("uuid = (?)", uuid).Delete(CycleCommet{}).Error
	if err != nil {
		return errors.New("参数错误")
	} else {
		return nil
	}
}

//根据uuid删除所有帖子的点赞
func CycleDelLikesByCycleuuid(uuid string) error {
	if len(uuid) <= 0 {
		return errors.New("参数错误")
	}

	err := DB.Where("uuid = (?)", uuid).Delete(CycleLike{}).Error
	if err != nil {
		return errors.New("参数错误")
	} else {
		return nil
	}
}
