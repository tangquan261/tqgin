package models

import (
	"time"
	//"errors"
	"fmt"
	//"time"
	//"fmt"
	//	"github.com/jinzhu/gorm"
)

type RoomTags struct {
	ID      int64 `gorm:"primary_key"`
	TagName string
}

type BannerInfo struct {
	BannerID    int32 `gorm:"primary_key"`
	Target_type int32
	Start_time  time.Time
	End_time    time.Time
	Bg_img      string
	Click_url   string
}

func GetTagList() []RoomTags {

	var tags []RoomTags
	err := DB.Find(&tags).Error

	if err != nil {
		fmt.Println("GetTagList err:", err)
	}
	return tags
}

func GetHotRoomsByTag(tagName string) []HotRoomInfo {

	var hotRooms []HotRoomInfo
	err := DB.Where("room_tag_name = ?", tagName).Find(&hotRooms).Error
	if err != nil {
		fmt.Println("GetHotRoomsByTag err:", err)
	}
	return hotRooms
}

func GetHotAllRooms() []HotRoomInfo {

	var hotRooms []HotRoomInfo

	err := DB.Find(&hotRooms).Error
	if err != nil {
		fmt.Println("GetHotRoomsByTag err:", err)
	}
	return hotRooms
}

func GetBanners() []BannerInfo {
	var banners []BannerInfo

	err := DB.Find(&banners).Error
	if err != nil {
		fmt.Println("GetBanners", err)
	}
	return banners
}
