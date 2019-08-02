package models

import (
	"fmt"
	"time"
)

type BannerInfo struct {
	BannerID    int32 `gorm:"primary_key"`
	Target_type int32
	Start_time  time.Time
	End_time    time.Time
	Bg_img      string
	Click_url   string
}

func GetHotRoomsByTag(tagName string) []RoomInfo {

	var hotRooms []HotRoomInfo
	err := DB.Where("room_tag_name = ?", tagName).Find(&hotRooms).Error
	if err != nil {
		fmt.Println("GetHotRoomsByTag err:", err)
		return nil
	}

	var rooms []RoomInfo
	for i := 0; i < len(hotRooms); i++ {
		oneRoom := GetRoomById(hotRooms[i].RoomID)
		rooms = append(rooms, *oneRoom)
	}

	return rooms
}

func GetHotAllRooms() []RoomInfo {

	var hotRooms []HotRoomInfo

	err := DB.Find(&hotRooms).Error
	if err != nil {
		fmt.Println("GetHotRoomsByTag err:", err)
		return nil
	}

	var rooms []RoomInfo
	for i := 0; i < len(hotRooms); i++ {
		oneRoom := GetRoomById(hotRooms[i].RoomID)
		rooms = append(rooms, *oneRoom)
	}
	return rooms
}

func GetBanners() []BannerInfo {
	var banners []BannerInfo

	err := DB.Find(&banners).Error
	if err != nil {
		fmt.Println("GetBanners", err)
	}
	return banners
}
