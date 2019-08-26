package models

import (
	"errors"
	"time"

	"tqgin/pkg/define"
	"tqgin/pkg/tqlog"

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
	LoginType define.LoginType
	LoginTime time.Time
	//屏蔽时间
	ForbidTime int64
	//屏蔽原因
	ForbidMsg string
	//登录验证tocken码
	Tocken        string
	TockenTimeOut time.Time
}

type AuthCode struct {
	gorm.Model
	Account  string               `gorm:"not null"`
	CodeType define.PhoneCodeType `gorm:"not null"` //1注册2修改密码3绑定手机s
	CodeText string               `gorm:"not null"`
}

func AuthSaveCode(authCode AuthCode) error {

	return DB.Model(AuthCode{}).Save(&authCode).Error
}

func AuthSaveCount(account string, codeType define.PhoneCodeType) int {

	var count int

	err := DB.Raw("select count(1) from tq_auth_code where account = (?) and code_type = (?) and day(created_at) = (day(?))",
		account, codeType, time.Now()).Count(&count).Error

	if err != nil {
		return -1
	}

	return count

}

func AuthGetCode(account string, codeType define.PhoneCodeType) *AuthCode {

	var auth AuthCode

	DBtemp := DB.Model(AuthCode{}).Where("account= (?) and code_type = (?)",
		account, codeType).Last(&auth)

	if DBtemp.RecordNotFound() {
		return nil
	}
	if DBtemp.Error != nil {
		return nil
	}

	return &auth
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

//返回playerid，0为错误
func Register(account Account) (int64, error) {

	var accountCurrent Account

	tempDB := DB.Where("account_id = (?)", account.AccountID).Find(&accountCurrent)

	if tempDB.Error != nil && !tempDB.RecordNotFound() {
		return 0, errors.New("重复注册")
	}

	if tempDB.RowsAffected != 0 {
		userInfo := GetUser(accountCurrent.PlayerID)
		if userInfo != nil {
			return 0, errors.New("重复注册")
		} else {
			account.PlayerID = accountCurrent.PlayerID
			AccountSave(account.AccountID, account)
			return accountCurrent.PlayerID, nil
		}
	}

	tx := DB.Begin()
	defer tx.Commit()

	tempDB = tx.Model(Account{}).Where("account_id = (?)", account.AccountID).FirstOrCreate(&account)

	if tempDB.Error != nil {
		return 0, errors.New("创建错误")
	}

	if tempDB.RowsAffected == 0 {
		return 0, errors.New("重复注册")
	}

	for {
		playerID := GetPlayerIDNext()

		if canCreatePlayerID(playerID) {
			err := tx.Model(Account{}).Where("account_id = (?)", account.AccountID).
				UpdateColumn("player_id", playerID).Error

			if err != nil {
				tx.Rollback()
				tqlog.TQSysLog.Error("create account error :", err)
				return 0, errors.New("创建错误")
			}
			account.PlayerID = playerID
			return playerID, nil
		}
	}

	return 0, nil
}

func AccountSave(accountid string, account Account) error {
	if len(accountid) <= 0 {
		return errors.New("参数错误")
	}

	return DB.Model(Account{}).Where("account_id = (?)", accountid).Update(&account).Error
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
