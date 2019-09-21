package models

type CarConfig struct {
	CarID     int32  //车id
	Name      string //车名字
	Icon      string //车头像
	Animation string //车动画
	CarDetail string //车简介
	CashType  int32  //消费类型
	CashCount int32  //消费数量
	Type      int32  //获取方式
	Duration  int64  //持有时间
}

var CarDic map[int32]*CarConfig

func init() {
	CarDic = make(map[int32]*CarConfig)
}

func GetCarmodel(id int32) *CarConfig {

	if obj, ok := CarDic[id]; ok {
		return obj
	} else {
		GetAllCarConfig()
		if obj, ok := CarDic[id]; ok {
			return obj
		}
	}

	return nil
}

func GetAllCarConfig() []CarConfig {

	rows, _ := DB.Raw("select * from config_car;").Rows()

	defer rows.Close()

	var rets []CarConfig

	for rows.Next() {
		model := new(CarConfig)
		err := rows.Scan(&model.CarID, &model.Name, &model.Icon,
			&model.Animation, &model.CarDetail, &model.CashType,
			&model.CashCount, &model.Type, &model.Duration)

		if err == nil {
			CarDic[model.CarID] = model
			rets = append(rets, *model)
		}
	}
	return rets
}
