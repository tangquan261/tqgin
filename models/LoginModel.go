package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	LoginType_Phone  = 1
	LoginType_QQ     = 2
	LoginType_WeChat = 3
)

type Account struct {
	gorm.Model
	AccountID string `gorm:"primary_key"`
	Password  string
	PlayerID  int64 `gorm:"not null"`
	LoginType int8
	LoginTime time.Time
	//屏蔽时间
	ForbidTime int64
	//屏蔽原因
	ForbidMsg string
}

func LoginAccount(accountID string) *Account {

	var account Account

	err := DB.Find(&account, "account_id = ?", accountID).GetErrors()

	if len(err) > 0 {
		log.Println(err, accountID)
	}

	return &account
}

func Register(account *Account) (status int, retAccount *Account) {

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
		retAccount = account
		DB.Create(account)

	} else {
		status = 1
		retAccount = nil
	}
	return
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
