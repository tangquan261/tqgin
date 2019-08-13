package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type RoomTags struct {
	ID      int64 `gorm:"primary_key"`
	TagName string
}

type RoomInfo struct {
	gorm.Model
	RoomID        int64 `gorm:"not null; unique"`
	MasterID      int64
	RoomName      string `gorm:"not null;unique"`
	RoomTagName   string
	RoomLevel     int
	RoomTotalStar int64
	RoomPic       string
	RoomIntro     string
	RoomPassword  string
	RoomCount     int64
}

type HotRoomInfo struct {
	RoomID      int64 `gorm:"primary_key"`
	RoomTagName string
	RoomHot     int64
	BeginTime   time.Time
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

func CreateRoom(room *RoomInfo) error {

	err := DB.Save(room).Error

	if err == nil {

		_, notFound := GetPowerRoom(room.MasterID, room.RoomID)
		if notFound {
			//没有找到该数据
			var roomPower RoomPowerMemberInfo
			roomPower.PlayerID = room.MasterID
			roomPower.RoomId = room.RoomID
			roomPower.RoomPower = 1
			DB.Save(&roomPower)
		}

		_, notFound = GetHotRoom(room.RoomID)
		if notFound {
			//没有在热门中找到
			count := getHotRoomCountByTag(room.RoomTagName)
			if count < 100 {
				//改类型的热门数量小于100
				newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTagName: room.RoomTagName, RoomHot: 100, BeginTime: time.Now()}
				addHotRoom(&newhotRoom)
			}
		} else {
			//在热门中找到
			//newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTagName: room.RoomTagName, RoomHot: 100, BeginTime: time.Now()}
			//addHotRoom(&newhotRoom)
		}
	}
	fmt.Println("创建房间的返回", err)
	return err
}

func addHotRoom(hotRoom *HotRoomInfo) {
	err := DB.Save(hotRoom).Error
	if err != nil {
		fmt.Println("add hot room error", err)
	}
}

func getHotRoomCountByTag(tagName string) int {

	var count int
	DB.Raw("select count(1) as total from tq_hot_room_info where room_tag_name = ? ", tagName).Count(&count)

	fmt.Println("getHotRoomCountByTag nums:", count)
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
		fmt.Println("GetRoomByIds", err)
		return nil
	}
	return &room
}
