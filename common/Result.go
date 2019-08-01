package tqgin

import (
	"net/http"
	"tqgin/pkg/errorcode"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func Result(ctx *gin.Context, code int, data interface{}, msg string) {

	_, bTure := data.(proto.Message)

	var retData interface{}

	if bTure {
		retData, _ = proto.Marshal(data.(proto.Message))

	} else {
		retData = data
	}

	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": retData, "msg": msg})
}

func ResultOk(ctx *gin.Context, data interface{}) {

	_, bTure := data.(proto.Message)

	var retData interface{}

	if bTure {
		retData, _ = proto.Marshal(data.(proto.Message))

	} else {
		retData = data
	}

	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": retData, "msg": ""})
}

func ResultOkMsg(ctx *gin.Context, data interface{}, msg string) {
	_, bTure := data.(proto.Message)

	var retData interface{}

	if bTure {
		retData, _ = proto.Marshal(data.(proto.Message))

	} else {
		retData = data
	}
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": retData, "msg": msg})
}

func ResultFail(ctx *gin.Context, msg string) {

	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.ERROR, "data": nil, "msg": msg})
}

func ResultFailData(ctx *gin.Context, data interface{}, msg string) {
	_, bTure := data.(proto.Message)

	var retData interface{}

	if bTure {
		retData, _ = proto.Marshal(data.(proto.Message))

	} else {
		retData = data
	}
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.ERROR, "data": retData, "msg": msg})
}
