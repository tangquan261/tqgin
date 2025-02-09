package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"time"
	"tqgin/conf"
	"tqgin/pkg/tqlog"

	"github.com/satori/go.uuid"
	"github.com/zheng-ji/goSnowFlake"
	//Snowflake 算法是Twitter的分布式ID自增算法,用于生成可以跨数据中心的全局唯一ID(不连续)。
)

var (
	iw     *goSnowFlake.IdWorker
	iwSnow *goSnowFlake.IdWorker
)

func init() {
	var err error

	workid := int64(config.GetConfigInt(config.Snow_Work_Id))

	iw, err = goSnowFlake.NewIdWorker(workid)
	if err != nil {
		log.Panic("error init goSnowFlake")
	}

	iwSnow, err = goSnowFlake.NewIdWorker(workid)
	if err != nil {
		log.Panic("error init goSnowFlake")
	}
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func Uids() string {
	var ret string

	bytes, err := uuid.NewV4()

	if err == nil {
		return bytes.String()
	}

	u1, err := iw.NextId()

	if err != nil {
		return EncodeMD5("tq" + time.Now().String())
	}

	ret = strconv.FormatInt(u1, 10)

	return ret
}

func UUID() int64 {

	id, err := iw.NextId()
	if err != nil {
		tqlog.TQSysLog.Error("error UUID goSnowFlake")
		return 0
	}
	return id

}

func SnowFlakeUUID() int64 {
	id, err := iw.NextId()
	if err != nil {
		tqlog.TQSysLog.Error("error SnowFlakeUUID goSnowFlake")
		return 0
	}
	return id
}
