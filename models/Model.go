package models

import (
	"fmt"
	"sync"
	"tqgin/pkg/tqlog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	PlayerID_MAX int64
	lockPlayerID *sync.Mutex
	DB           *gorm.DB
)

func init() {
	lockPlayerID = new(sync.Mutex)
}

func ConfigDB() {

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
	DB.SetLogger(tqlog.NewDBLogger())
	DB.SingularTable(true)
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Account{}, &UserInfo{}, &RoomInfo{}, &RoomPowerMemberInfo{},
		&RoomTags{}, &HotRoomInfo{}, &SupportRoom{}, &BannerInfo{},
		&MicUserModel{}, &Black{}, &GifInfo{}, &GifGiveRecord{}, &RankInfo{},
		&RoomRankInfo{}, &RelationShip{}, &CycleModel{}, &CycleCommet{},
		&CycleLike{}, &MoneyAccount{}, &ConsumeUserCount{}, &GfitUserCount{}, &MicQueue{})

	loadConf()

	PlayerID_MAX = LoginLastPlayerID()
	fmt.Println("db init success")
}

func loadConf() {
	GetAllGift()
}

func GetPlayerIDNext() int64 {
	lockPlayerID.Lock()
	defer lockPlayerID.Unlock()
	PlayerID_MAX++
	tqlog.TQSysLog.Info("get playerID next", PlayerID_MAX)
	return PlayerID_MAX
}

func LoginLastPlayerID() int64 {

	var account Account

	tempDB := DB.Model(Account{}).Last(&account)

	if tempDB.Error != nil {
		tqlog.TQSysLog.Panic("LoginLastPlayerID error", tempDB.Error)
		return 0
	}

	if tempDB.RowsAffected == 0 {
		tqlog.TQSysLog.Warn("LoginLastPlayerID 创建第一个用户，id 1000000")
		return 100000
	}
	return account.PlayerID
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
