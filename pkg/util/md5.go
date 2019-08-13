package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
	"github.com/zheng-ji/goSnowFlake"
)

var (
	iw *goSnowFlake.IdWorker
)

func init() {
	var err error
	iw, err = goSnowFlake.NewIdWorker(1)
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

func Uidu() int64 {

	id, err := iw.NextId()
	if err != nil {
		log.Panic("error Uidu goSnowFlake")
		return 0
	}
	return id

}
