package models

import (
	//"errors"
	"fmt"
	"time"
	//"fmt"
	//	"github.com/jinzhu/gorm"
)

type RoomTags struct {
	ID      int64 `gorm:"primary_key"`
	TagName string
}

type RoomInfo struct {
	RoomID        int64 `gorm:"primary_key"`
	MasterID      int64
	RoomName      string `gorm:"not null;unique"`
	RoomTagName   string
	RoomLevel     int
	RoomTotalStar int64
	RoomHot       int64
}

type HotRoomInfo struct {
	RoomID      int64 `gorm:"primary_key"`
	RoomTagName string
	BeginTime   time.Time
	RoomHot     int64
}

type RoomPowerMemberInfo struct {
	RoomId   int64
	PlayerID int64
	//1 房主，2管理员
	RoomPower int
}

//扶持房间的ID信息
type SupportRoom struct {
	RoomID    int64 `gorm:"not null;unique"`
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

		roomPower := GetPowerRoom(room.MasterID, room.RoomID)

		var roomPower RoomPowerMemberInfo
		roomPower.PlayerID = room.MasterID
		roomPower.RoomId = room.RoomID
		roomPower.RoomPower = 1
		DB.Save(&roomPower)

		_, err := GetHotRoom(room.RoomID)
		if err == nil {
			newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTagName: room.RoomTagName, RoomHot: 100, BeginTime: time.Now()}
			addHotRoom(&newhotRoom)
		} else {
			count := getHotRoomCountByTag(room.RoomTagName)
			if count < 100 {
				newhotRoom := HotRoomInfo{RoomID: room.RoomID, RoomTagName: room.RoomTagName, RoomHot: 100, BeginTime: time.Now()}
				addHotRoom(&newhotRoom)
			}
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

func GetHotRoom(roomID int64) (HotRoomInfo, error) {

	var hotRoom HotRoomInfo

	err := DB.Where("room_id = ?", roomID).Find(&hotRoom).Error
	return hotRoom, err
}

func GetPowerRoom(playerID int64, roomID int64) (RoomPowerMemberInfo, error) {
	var power RoomPowerMemberInfo
	err := DB.Where("room_id = ?  AND player_id = ?", roomID, playerID).Find(&power).Error

	return power, err
}
