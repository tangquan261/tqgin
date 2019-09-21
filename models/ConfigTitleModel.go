package models

import (
	"fmt"
)

type TitleConfig struct {
	ID             int32
	Name           string
	Icon           string
	Icon_animation int
	Duration       int64 //秒
}

var TitleDic map[int32]*TitleConfig

func init() {
	TitleDic = make(map[int32]*TitleConfig)
}

func GetTitlemodel(id int32) *TitleConfig {

	if obj, ok := TitleDic[id]; ok {
		return obj
	}

	rows, _ := DB.Raw("select * from config_title where id = (?);", id).Rows()

	defer rows.Close()

	for rows.Next() {
		model := new(TitleConfig)
		err := rows.Scan(&model.ID, &model.Name, &model.Icon, &model.Icon_animation)
		fmt.Println("loadd titl data:", model, err)
		if err == nil {
			model.Duration = 60 * 60 * 24
			TitleDic[model.ID] = model
			return model
		}
	}
	return nil
}
