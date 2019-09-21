package controllers

//称号

import (
	"time"
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	tqgin.Controller
}

func (this *CarController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/car")
	temp.POST("getShopsCar", this.shopsCar)
	temp.POST("applyCars", this.applyMyCars)
	temp.POST("modifyWearCar", this.modifyWearCar)
}

func (r *CarController) modifyWearCar(con *gin.Context) {
	playerID := r.GetPlayerGUID(con)

	type modifyCar struct {
		CarID int32 `json:"carid"`
	}
	var param modifyCar

	err := con.ShouldBindJSON(&param)
	if err != nil {
		tqgin.ResultFail(con, "错误")
		return
	}

	models.ModifyWearCar(playerID, param.CarID)

	tqgin.ResultOk(con, param)
}

func (r *CarController) applyMyCars(con *gin.Context) {
	playerID := r.GetPlayerGUID(con)

	type RetParam struct {
		CarID     int32
		IsWear    int32
		EndTime   int64  //剩余时间
		Name      string //车名字
		Icon      string //车头像
		Animation string //车动画
		CarDetail string //车简介
	}

	var retCars []RetParam
	cars := models.LoadCars(playerID)

	for _, car := range cars {
		var myCar RetParam
		conCar := models.GetCarmodel(car.CarID)

		myCar.CarID = car.CarID
		myCar.Name = conCar.Name
		myCar.Icon = conCar.Icon
		myCar.Animation = conCar.Animation
		myCar.IsWear = car.IsWear
		myCar.EndTime = car.EndTime.Unix() - time.Now().Unix()
		myCar.CarDetail = conCar.CarDetail

		retCars = append(retCars, myCar)
	}
	tqgin.ResultOkMsg(con, retCars, "成功")
}

func (r *CarController) shopsCar(con *gin.Context) {

	cars := models.GetAllCarConfig()

	tqgin.ResultOk(con, cars)
}
