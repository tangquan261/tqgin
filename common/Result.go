package tqgin

import (
	"net/http"

	"tqgin/pkg/errorcode"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func Result(ctx *gin.Context, code int, data proto.Message, msg string) {
	protodata, _ := proto.Marshal(data)
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": protodata, "msg": msg})
}

func ResultOk(ctx *gin.Context, data proto.Message) {
	protodata, _ := proto.Marshal(data)
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": protodata, "msg": ""})
}

func ResultOkMsg(ctx *gin.Context, data proto.Message, msg string) {
	protodata, _ := proto.Marshal(data)
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": protodata, "msg": msg})
}

func ResultFail(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": nil, "msg": err})
}

func ResultFailData(ctx *gin.Context, data proto.Message, err interface{}) {
	protodata, _ := proto.Marshal(data)
	ctx.JSON(http.StatusOK, gin.H{"code": errorcode.SUCCESS, "data": protodata, "msg": err})
}
