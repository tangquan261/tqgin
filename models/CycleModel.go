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

// func CycleGetSound(index int64) []CycleModel {
// 	return nil
// }

// func CycleGetFeeds(index int64) []CycleModel {

// }
