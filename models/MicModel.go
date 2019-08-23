package models

import (
	"fmt"
	"time"
)

//房间麦序状态，
type MicUserModel struct {
	RoomID     int64 //房间id
	PlayerID   int64 //用户id，为0表示该麦序没有人
	MicIndex   int16 //麦序
	MicStatus  int16 //麦序状态0是开启，1关闭不可上
	UpdateDate time.Time
}

//排队麦序
type MicQueue struct {
	RoomID  int64
	PlayeID int64
}

//加入排队序列
func MicQueueAdd(roomID, playerID int64) error {
	//先尝试移除
	MicQueueDel(playerID)

	var micqueue MicQueue
	micqueue.RoomID = roomID
	micqueue.PlayeID = playerID

	err := DB.Model(MicQueue{}).
		Where("player_id = (?) and room_id = (?)", playerID, roomID).
		FirstOrCreate(&micqueue).Error

	return err
}

//从排队序列中移出
func MicQueueDel(playerID int64) error {
	err := DB.Model(MicQueue{}).Where("player_id = (?)", playerID).Delete(MicQueue{}).Error

	return err
}

//获取房间排队序列
func MicQueueInfo(roomID int64) []MicQueue {
	var rets []MicQueue

	err := DB.Model(MicQueue{}).Where("room_id = (?)", roomID).Find(&rets).Error

	if err != nil {
		return rets
	}
	return rets
}

//添加到mic
func MicAdd(roomID, playerGUID int64, micIndex int16) error {

	MicDelByPlayerID(playerGUID)

	var mic MicUserModel
	mic.RoomID = roomID
	mic.PlayerID = playerGUID
	mic.MicIndex = micIndex
	mic.UpdateDate = time.Now()

	DBTemp := DB.Model(MicUserModel{}).
		Where("room_id = (?) and mic_index = (?)", roomID, micIndex).
		FirstOrCreate(&mic)

	if DBTemp.Error != nil {
		return DBTemp.Error
	} else if DBTemp.RowsAffected == 0 {
		//麦序上有人
		if mic.UpdateDate.Add(30 * time.Second).Before(time.Now()) {
			//麦序上人无效
			err := DB.Model(MicUserModel{}).
				Where("room_id = (?) and mic_index = (?)", roomID, micIndex).
				UpdateColumns(MicUserModel{PlayerID: playerGUID, UpdateDate: time.Now()}).Error

			if err != nil {
				return err
			}
		}

	} else {
		//麦序上没有人则什么都不做

	}
	return nil
}

//从mic上移除
func MicDelByPlayerID(playerGUID int64) error {

	err := DB.Where("player_id = (?)", playerGUID).Delete(MicUserModel{}).Error
	if err != nil {

	}
	return err
}

//房间解散，移除所有mic
func MicDelByDismiss(roomID int64) error {
	err := DB.Where("room_id = (?)", roomID).Delete(MicUserModel{}).Error
	if err != nil {

	}
	return err
}

func MicUpdateStatus(roomID int64, MicIndex, MicStatus int16) error {

	err := DB.Model(MicUserModel{}).
		Where("room_id = (?) and  mic_index = (?)", roomID, MicIndex).
		Update("mic_status", MicStatus).Error

	return err
}

//根据房间id获取到所有麦序
func MicGetAllIndex(roomID int64) []MicUserModel {

	var mics []MicUserModel

	lesstime := time.Now().Add(-60 * time.Second)

	fmt.Println(lesstime.Unix(), time.Now().Unix())
	err := DB.Model(MicUserModel{}).Where("room_id = (?) and update_date > (?)",
		roomID, lesstime).Find(&mics).Error

	fmt.Println(err, "----", mics)
	if err != nil {
		return mics
	}
	return mics
}

//更新麦序状态
func MicUpdate(roomID, playeID int64, MicIndex int16) []MicUserModel {

	var newMic MicUserModel

	newMic.RoomID = roomID
	newMic.PlayerID = playeID
	newMic.MicIndex = MicIndex
	newMic.MicStatus = 0

	//尝试第一次创建
	DBTemp := DB.Model(MicUserModel{}).
		Where("room_id = (?) and mic_index = (?)", roomID, MicIndex).
		FirstOrCreate(&newMic)

	if DBTemp.Error != nil {
		return nil
	}

	if DBTemp.RowsAffected != 0 {
		//创建成功
		//某房间的特定麦序插入成功
	} else {

		if newMic.MicStatus == 1 {
			//麦序关闭，不能操作
		} else {
			//创建失败，已经存在了
			DB.Model(MicUserModel{}).
				Where("room_id = (?) and mic_index = (?) and update_date = (?)",
					roomID, MicIndex, time.Now()).Update(&newMic)
		}

	}

	return MicGetAllIndex(roomID)
}
