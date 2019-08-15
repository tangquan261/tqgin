package tqgin

import (
	"net/http"
	"strconv"
	"tqgin/pkg/errorcode"

	"tqgin/pkg/tqlog"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func playerID(ctx *gin.Context) int64 {
	myID, _ := ctx.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	return nPlayerID
}

func Result(ctx *gin.Context, code int, data interface{}, msg string) {

	_, bTure := data.(proto.Message)

	var retData interface{}

	if bTure {
		retData, _ = proto.Marshal(data.(proto.Message))

	} else {
		retData = data
	}

	tqlog.TQRequest.Debug("playerID:", playerID(ctx), " code:", code, " data:", retData, " msg:", msg)

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

	tqlog.TQRequest.Debug("playerID:", playerID(ctx), " code:", errorcode.SUCCESS, " data:", retData, " msg:", "")
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

	tqlog.TQRequest.Debug("playerID:", playerID(ctx), " code:", errorcode.SUCCESS, " data:", retData, " msg:", msg)
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": retData, "msg": msg})
}

func ResultFail(ctx *gin.Context, msg string) {

	tqlog.TQRequest.Debug("playerID:", playerID(ctx), " code:", errorcode.ERROR, " data:", " msg:", msg)
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

	tqlog.TQRequest.Debug("playerID:", playerID(ctx), " code:", errorcode.ERROR, " data:", retData, " msg:", msg)
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.ERROR, "data": retData, "msg": msg})
}
