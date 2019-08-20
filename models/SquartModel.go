package models

import (
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
	//
	err := DB.Where("room_tag = ? and (end_time IS NULL or end_time < (?))", tagName, time.Now()).Order("room_hot desc").Limit(100).Find(&hotRooms).Error
	if err != nil {
		return nil
	}

	var rooms []RoomInfo
	for i := 0; i < len(hotRooms); i++ {
		oneRoom := GetRoomById(hotRooms[i].RoomID)
		if oneRoom != nil {
			rooms = append(rooms, *oneRoom)
		}
	}

	return rooms
}

func GetHotAllRooms() []RoomInfo {

	var hotRooms []HotRoomInfo

	err := DB.Order("begin_time ASC").Limit(100).Find(&hotRooms).Error
	if err != nil {
		return nil
	}

	var rooms []RoomInfo
	for i := 0; i < len(hotRooms); i++ {
		oneRoom := GetRoomById(hotRooms[i].RoomID)
		if oneRoom != nil {
			rooms = append(rooms, *oneRoom)
		}
	}
	return rooms
}

func GetBanners() []BannerInfo {
	var banners []BannerInfo

	err := DB.Find(&banners).Error
	if err != nil {
		return nil
	}
	return banners
}
