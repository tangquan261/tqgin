package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type Friend struct {
	gorm.Model
	PlayerID int64 `gorm:"not null"`
	FriendID int64 `gorm:"not null"`
}

type Black struct {
	gorm.Model
	PlayerID int64 `gorm:"not null"`
	BlackID  int64 `gorm:"not null"`
}

type Attention struct {
	gorm.Model
	PlayerID    int64 `gorm:"not null"`
	TarPlayerID int64 `gorm:"not null"`
}

func GetFriends(playerID int64) []Friend {
	if playerID <= 0 {
		return nil
	}
	var user Friend
	user.PlayerID = playerID
	var users []Friend

	err := DB.Where(user).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func GetBlacks(playerID int64) []Black {
	if playerID <= 0 {
		return nil
	}
	var user Black
	user.PlayerID = playerID
	var users []Black

	err := DB.Where(user).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func GetAttentions(playerID int64) []Attention {
	if playerID <= 0 {
		return nil
	}
	var user Attention
	user.PlayerID = playerID
	var users []Attention

	err := DB.Where(user).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func AddFriend(playerID, friendID int64) error {
	if playerID <= 0 || friendID <= 0 {
		return errors.New("addFriend error")
	}

	if playerID == friendID {
		return errors.New("addFriend error self")
	}

	hasUser := UserHasInfo(friendID)
	if !hasUser {
		return errors.New("not find player")
	}

	var friend Friend
	friend.PlayerID = playerID
	friend.FriendID = friendID
	err := DB.Where(&friend).FirstOrCreate(&friend).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func RemoveFriend(playerID, friendID int64) error {
	if playerID <= 0 || friendID <= 0 {
		return errors.New("RemoveFriend error")
	}

	if playerID == friendID {
		return errors.New("RemoveFriend error self")
	}

	hasUser := UserHasInfo(friendID)
	if !hasUser {
		return errors.New("not find player")
	}
	var friend Friend
	friend.PlayerID = playerID
	friend.FriendID = friendID

	return DB.Where(friend).Delete(friend).Error
}

func AddBlack(playerID, blackID int64) error {
	if playerID <= 0 || blackID <= 0 {
		return errors.New("AddBlack error")
	}

	if playerID == blackID {
		return errors.New("AddBlack error self")
	}

	hasUser := UserHasInfo(blackID)
	if !hasUser {
		return errors.New("not find player")
	}

	var black Black
	black.PlayerID = playerID
	black.BlackID = blackID
	err := DB.Where(&black).FirstOrCreate(&black).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func RemoveBlack(playerID, blackID int64) error {
	if playerID <= 0 || blackID <= 0 {
		return errors.New("RemoveBlack error")
	}

	if playerID == blackID {
		return errors.New("RemoveBlack error self")
	}

	hasUser := UserHasInfo(blackID)
	if !hasUser {
		return errors.New("not find player")
	}
	var balck Black
	balck.PlayerID = playerID
	balck.BlackID = blackID

	return DB.Where(balck).Delete(balck).Error
}

func AddAttention(playerID, tarPlayerID int64) error {
	if playerID <= 0 || tarPlayerID <= 0 {
		return errors.New("addFriend error")
	}

	if playerID == tarPlayerID {
		return errors.New("addFriend error self")
	}

	hasUser := UserHasInfo(tarPlayerID)
	if !hasUser {
		return errors.New("not find player")
	}

	var attention Attention
	attention.PlayerID = playerID
	attention.TarPlayerID = tarPlayerID
	err := DB.Where(&attention).FirstOrCreate(&attention).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func RemoveAttention(playerID, tarPlayerID int64) error {
	if playerID <= 0 || tarPlayerID <= 0 {
		return errors.New("RemoveAttention error")
	}

	if playerID == tarPlayerID {
		return errors.New("RemoveAttention error self")
	}

	hasUser := UserHasInfo(tarPlayerID)
	if !hasUser {
		return errors.New("not find player")
	}
	var attention Attention
	attention.PlayerID = playerID
	attention.TarPlayerID = tarPlayerID

	return DB.Where(attention).Delete(attention).Error
}
