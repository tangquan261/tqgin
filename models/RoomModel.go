package models

import (
	"time"
	//"fmt"
	"github.com/jinzhu/gorm"
)

type RoomTags struct {
	ID      int64 `gorm:"primary_key"`
	TagName string
}

type RoomInfo struct {
	gorm.Model
	RoomID        int64 `gorm:"not null;unique"`
	MasterID      int64
	RoomName      string `gorm:"not null;unique"`
	RoomTagName   string
	RoomLevel     int
	RoomTotalStar int64
	RoomHot       int64
}

type HotRoomInfo struct {
	ID             int64 `gorm:"primary_key"`
	RoomID         int64 `gorm:"not null;unique"`
	MasterID       int64
	RoomName       string `gorm:"not null;unique"`
	RoomTagName    string
	RoomLevel      int
	RoomTotalStart int64
	RoomHot        int64
	BeginTime      time.Time
	EndTime        time.Time
}

type RoomPowerMemberInfo struct {
	gorm.Model
	RoomId   int64
	PlayerID int64
	//1 房主，2管理员
	RoomPower int
	//在该房间的星光值
	Star int64
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

func CreateRoom(room *RoomInfo) bool {

	return true
}

func SaveRoom(room *RoomInfo) bool {

	return true
}

func GetRoom(roomID int64) *RoomInfo {

	var info RoomInfo

	return &info
}

func GetHotRoomsByTag(tag string) []*HotRoomInfo {

	var rooms []*HotRoomInfo

	DB.Where("RoomTagName = ?", tag).Find(&rooms)

	return rooms
}
