package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type RoomTags struct {
	ID      int64 `gorm:"primary_key"`
	TagName string
}

type RoomInfo struct {
	gorm.Model
	RoomID        int64  `gorm:"not null"`        //房间id=playerid
	RoomName      string `gorm:"not null;unique"` //房间名字
	RoomIntro     string `gorm:""`                //房间介绍
	RoomNotice    string `gorm:""`                //介绍公告
	RoomTag       string `gorm:""`                //房间tag
	RoomAudioType int32  `gorm:""`                //房间声音类型1：娱乐模式，2：天籁之音，3：超高音质
	RoomTotalStar int64  `gorm:""`                //房间总星
	RoomPic       string `gorm:""`                //房间头像
	RoomBGPic     string `gorm:""`                //房间背景图
	RoomPassword  string `gorm:""`                //房间密码
	RoomOrderMic  int32  `gorm:""`                //上麦类型，1，2自由上麦，排序上麦

}

type HotRoomInfo struct {
	RoomID    int64 `gorm:"primary_key"`
	RoomTag   string
	RoomHot   int64
	BeginTime time.Time //开始时间
	EndTime   time.Time //关闭时间
}

type RoomPowerMemberInfo struct {
	gorm.Model
	RoomId   int64
	PlayerID int64
	//1 房主，2管理员
	RoomPower int
}

//扶持房间的ID信息
type SupportRoom struct {
	RoomID    int64 `gorm:"primary_key"`
	BeginTime time.Time
	EndTime   time.Time
	Msg       string
}

func GetTagList() []*RoomTags {

	var tags []*RoomTags
	DB.Find(&tags)

	return tags
}

func CreateRoom(room RoomInfo) error {

	if room.RoomID <= 0 {
		return errors.New("创建失败")
	}

	createRoom := room
	DBtemp := DB.Where("room_id =(?)", room.RoomID).FirstOrCreate(&createRoom)

	if DBtemp.Error != nil {
		return errors.New("创建失败")
	}

	if DBtemp.RowsAffected == 0 {
		//已经存在了，更新
		err := DB.Model(RoomInfo{}).Where("room_id =(?)", room.RoomID).Update(&room).Error

		if err != nil {
			return errors.New("创建失败")
		}
	}

	support := GetSupportRoomById(room.RoomID)
	if support != nil {
		//官方支持的房间
		newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTag: room.RoomTag, RoomHot: 999999, BeginTime: time.Now()}
		addHotRoom(&newhotRoom)
		OpenHotRoom(room.RoomID)
		return nil
	}

	hotroom := GetHotRoom(room.RoomID)
	if hotroom == nil {
		count := GetHotRoomCountByTag(room.RoomTag)
		//数量小于100
		if count < 100 {
			//改类型的热门数量小于100
			newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTag: room.RoomTag, RoomHot: 100, BeginTime: time.Now()}
			addHotRoom(&newhotRoom)
			return nil
		}
	} else {
		OpenHotRoom(room.RoomID)
	}

	return nil
}

//添加热门房间
func addHotRoom(hotRoom *HotRoomInfo) {
	DB.Where("room_id = ?", hotRoom.RoomID).FirstOrCreate(&hotRoom)
}

//打开热门房间
func OpenHotRoom(roomid int64) error {
	return DB.Model(HotRoomInfo{}).Where("room_id = ?", roomid).
		Updates(map[string]interface{}{"begin_time": time.Now(), "end_time": nil}).Error
}

//关闭热门房间
func CloseHotRoom(roomid int64) error {
	return DB.Model(HotRoomInfo{}).Where("room_id = ?", roomid).
		UpdateColumn("end_time", time.Now()).Error
}

//更新热门值
func UpdateHotRomhot(roomid int64, addHot int64) error {
	return DB.Model(HotRoomInfo{}).Where("room_id=(?)", roomid).
		UpdateColumn("room_hot", gorm.Expr("room_hot + ?", addHot)).Error
}

//清空所有热门值
func ClearHotRooms() {
	DB.Model(HotRoomInfo{}).UpdateColumn("room_hot", 100)
}

//根据tag类型，获取所有该类型房间数量
func GetHotRoomCountByTag(tagName string) int {

	var count int
	DB.Raw("select count(1) as total from tq_hot_room_info where room_tag_name = ? ", tagName).Count(&count)

	return count
}

//根据房间id获取热门房间
func GetHotRoom(roomID int64) *HotRoomInfo {
	var hotRoom HotRoomInfo
	DBTemp := DB.Where("room_id = ?", roomID).Find(&hotRoom)

	if DBTemp.Error != nil || DBTemp.RowsAffected == 0 {
		return nil
	}

	return &hotRoom
}

func GetPowerRoom(playerID int64, roomID int64) (RoomPowerMemberInfo, bool) {
	var power RoomPowerMemberInfo
	notFound := DB.Where("room_id = (?)  AND player_id = (?)", roomID, playerID).Find(&power).RecordNotFound()
	return power, notFound
}

func GetRoomById(roomid int64) *RoomInfo {

	var room RoomInfo

	DBtemp := DB.Where("room_id = (?)", roomid).Find(&room)

	if DBtemp.Error != nil || DBtemp.RowsAffected == 0 {

		return nil
	}
	return &room
}

func GetSupportRoomById(roomid int64) *SupportRoom {

	var room SupportRoom

	DBtemp := DB.Where("room_id = (?)", roomid).Find(&room)

	if DBtemp.Error != nil || DBtemp.RowsAffected == 0 {

		return nil
	}
	return &room
}

//修改房间信息
func SaveRoominfo(roomID int64, roominfo RoomInfo) error {
	if roomID <= 0 {
		return errors.New("SaveRoominfo error")
	}
	err := DB.Model(RoomInfo{}).Where("room_id = (?)", roomID).Update(&roominfo).Error
	return err
}
