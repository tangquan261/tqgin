package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"tqgin/proto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	AccountID string `gorm:"primary_key"`
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

	err := DB.Find(&account, "account_id = ?", accountID).GetErrors()

	if len(err) > 0 {
		log.Println(err, accountID)
	}

	return &account
}

func LoginAccountByPlayerID(PlayerID int64) *Account {

	var account Account

	err := DB.Find(&account, "player_id = ?", PlayerID).GetErrors()

	if len(err) > 0 {
		log.Println(err, PlayerID)
	}

	return &account
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

func AccountSaveTocken(account *Account) error {

	if account.AccountID == "" {
		return errors.New("AccountID is nil")
	}

	return DB.Model(account).Update(Account{Tocken: account.Tocken, TockenTimeOut: account.TockenTimeOut}).Error
}

func AccountChangePwd(account *Account, newPassword string) int {

	if account.AccountID == "" {
		return 1
	}

	err := DB.Model(account).Update(Account{Password: newPassword}).GetErrors()

	if len(err) > 0 {
		fmt.Println(err)
		return 1
	} else {
		return 0
	}
}
