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
	RoomAudioType int32  `gorm:""`                //房间声音类型
	RoomTotalStar int64  `gorm:""`                //房间总星
	RoomPic       string `gorm:""`                //房间头像
	RoomPassword  string `gorm:""`                //房间密码
}

type HotRoomInfo struct {
	RoomID    int64 `gorm:"primary_key"`
	RoomTag   string
	RoomHot   int64
	BeginTime time.Time
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

	DBtemp := DB.Where("room_id =(?)", room.RoomID).FirstOrCreate(&room)

	if DBtemp.Error != nil {
		return errors.New("创建失败")
	}

	if DBtemp.RowsAffected == 0 {
		return errors.New("创建失败")
	}

	_, notFound := GetPowerRoom(room.RoomID, room.RoomID)
	if notFound {
		//没有找到该数据
		var roomPower RoomPowerMemberInfo
		roomPower.PlayerID = room.RoomID
		roomPower.RoomId = room.RoomID
		roomPower.RoomPower = 1
		DB.Save(&roomPower)
	}

	_, notFound = GetHotRoom(room.RoomID)
	if notFound {
		//没有在热门中找到
		count := getHotRoomCountByTag(room.RoomTag)
		if count < 100 {
			//改类型的热门数量小于100
			newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTag: room.RoomTag, RoomHot: 100, BeginTime: time.Now()}
			addHotRoom(&newhotRoom)
		}
	} else {
		//在热门中找到
		//newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTagName: room.RoomTagName, RoomHot: 100, BeginTime: time.Now()}
		//addHotRoom(&newhotRoom)
	}
	return nil
}

func addHotRoom(hotRoom *HotRoomInfo) {
	err := DB.Save(hotRoom).Error
	if err != nil {

	}
}

func getHotRoomCountByTag(tagName string) int {

	var count int
	DB.Raw("select count(1) as total from tq_hot_room_info where room_tag_name = ? ", tagName).Count(&count)

	return count
}

func GetHotRoom(roomID int64) (HotRoomInfo, bool) {
	var hotRoom HotRoomInfo
	notFound := DB.Where("room_id = ?", roomID).Find(&hotRoom).RecordNotFound()
	return hotRoom, notFound
}

func GetPowerRoom(playerID int64, roomID int64) (RoomPowerMemberInfo, bool) {
	var power RoomPowerMemberInfo
	notFound := DB.Where("room_id = (?)  AND player_id = (?)", roomID, playerID).Find(&power).RecordNotFound()
	return power, notFound
}

func GetRoomById(roomid int64) *RoomInfo {
	var room RoomInfo
	err := DB.Where("room_id = (?)", roomid).Find(&room).Error
	if err != nil {

		return nil
	}
	return &room
}
