package models

import (
	"errors"

	"time"

	"tqgin/pkg/tqlog"
	"tqgin/proto"

	"github.com/jinzhu/gorm"
)

type MoneyAccount struct {
	MoneyPlayerID int64 `gorm:"primary_key"`
}

type Account struct {
	gorm.Model
	AccountID string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	PlayerID  int64  `gorm:"not null"`
	LoginType login.LoginType
	LoginTime time.Time
	//屏蔽时间
	ForbidTime int64
	//屏蔽原因
	ForbidMsg string
	//登录验证tocken码
	Tocken        string
	TockenTimeOut time.Time
}

func LoginAccount(accountID string) *Account {

	var account Account

	DBtemp := DB.Find(&account, "account_id = ?", accountID)
	if DBtemp.Error != nil {
		return nil
	}
	if DBtemp.RowsAffected <= 0 {
		return nil
	}

	return &account
}

func LoginAccountByPlayerID(PlayerID int64) *Account {

	var account Account

	DBtemp := DB.Find(&account, "player_id = ?", PlayerID)
	if DBtemp.Error != nil {
		return nil
	}
	if DBtemp.RowsAffected <= 0 {
		return nil
	}
	return &account
}

func Register(account Account) error {

	var accountCurrent Account

	tempDB := DB.Where("account_id = (?)", account.AccountID).Find(&accountCurrent)

	if tempDB.Error != nil {
		return errors.New("错误")
	}

	if tempDB.RowsAffected != 0 {

		userInfo := GetUser(accountCurrent.PlayerID)
		if userInfo != nil {
			return errors.New("已经存在")
		} else {
			account.PlayerID = accountCurrent.PlayerID
			AccountSave(account.AccountID, account)
			return nil
		}
	}

	tx := DB.Begin()
	defer tx.Commit()

	tempDB = tx.Model(Account{}).Where("account_id = (?)", account.AccountID).FirstOrCreate(&account)

	if tempDB.Error != nil {
		return errors.New("创建错误")
	}

	if tempDB.RowsAffected == 0 {
		return errors.New("已经存在")
	}

	for {
		playerID := GetPlayerIDNext()

		if canCreatePlayerID(playerID) {
			err := tx.Model(Account{}).Update("player_id = (?)").
				Where("account_id = (?)", account.AccountID).Error

			if err != nil {
				tx.Rollback()
				tqlog.TQSysLog.Error("create account error :", err)
				return errors.New("创建错误")
			}
			account.PlayerID = playerID
			break
		}
	}

	return nil
}

func AccountSave(accountid string, account Account) error {
	if len(account.AccountID) <= 0 {
		return errors.New("参数错误")
	}

	return DB.Model(Account{}).Where("account_id = (?)", accountid).Update(account).Error
}

//判断当前playerid是否为钱号，不能自动创建
func canCreatePlayerID(PlayerID int64) bool {

	var count int
	err := DB.Model(MoneyAccount{}).Where("money_player_id = (?)", PlayerID).Count(&count).Error

	if err != nil {
		return false
	}
	if count > 0 {
		return false
	}
	return true
}
