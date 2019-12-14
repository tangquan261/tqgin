package service

import (
	"tqgin/models"
)

var g_playerModel map[int64]*ServiceUserInfo

func init() {
	g_playerModel = make(map[int64]*ServiceUserInfo)
}

type ServiceUserInfo struct {
	models.UserInfo
}

func GetUserInfo(playerGUID int64) *ServiceUserInfo {
	if playerGUID == 0 {
		return nil
	}

	if obj, ok := g_playerModel[playerGUID]; ok {
		return obj
	} else {
		userinfo := new(ServiceUserInfo)

		obj := models.GetUser(playerGUID)

		userinfo.UserInfo = *obj
		g_playerModel[playerGUID] = userinfo
		return userinfo
	}
}
