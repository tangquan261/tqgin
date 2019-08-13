package models

import (
	"errors"
	"log"
	"time"

	"tqgin/proto"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	AccountID string `gorm:"not null;unique"`
	Password  string
	PlayerID  int64 `gorm:"not null"`
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

	err := DB.Find(&account, "account_id = ?", accountID).Error

	if err != nil {
		log.Println(err, accountID)
	}

	return &account
}

func LoginAccountByPlayerID(PlayerID int64) Account {

	var account Account

	err := DB.Find(&account, "player_id = ?", PlayerID).Error

	if err != nil {
		log.Println(err, PlayerID)
	}

	return account
}

func Register(account *Account) (status int) {

	var accountCurrent Account

	status = 0

	DB.Find(&accountCurrent, "account_id = ?", account.AccountID)

	if accountCurrent.AccountID == "" {
		status = 0

		DB.Last(&accountCurrent)

		if accountCurrent.AccountID == "" {
			account.PlayerID = 20000
		} else {
			account.PlayerID = accountCurrent.PlayerID + 1
		}
		account.LoginTime = time.Now()

		DB.Create(account)

	} else {
		status = 1
	}
	return
}

func AccountSave(accountid string, account Account) error {
	if account.AccountID == "" {
		return errors.New("account is nui")
	}

	return DB.Model(Account{}).Where("account_id = (?)", accountid).Update(account).Error
}
