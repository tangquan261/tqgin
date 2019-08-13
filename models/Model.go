package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func init() {

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tq_" + defaultTableName
	}

	var err error
	DB, err = gorm.Open("mysql", "root:tangquan@tcp(127.0.0.1:3306)/dbtest?charset=utf8&parseTime=True&loc=Local")

	if nil != err {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(true)
	DB.SingularTable(true)
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Account{}, &UserInfo{}, &RoomInfo{}, &RoomPowerMemberInfo{},
		&RoomTags{}, &HotRoomInfo{}, &SupportRoom{}, &BannerInfo{},
		&MicModel{}, &Black{}, &GifInfo{}, &GifGiveRecord{}, &RankInfo{},
		&RoomRankInfo{}, &RelationShip{})

	loadConf()
	fmt.Println("db init success")
}

func loadConf() {
	GetAllGift()
}

/*
DB.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

DB.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

DB.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';

DB.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;

*/
