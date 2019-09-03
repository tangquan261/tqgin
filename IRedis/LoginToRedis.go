package IRedis

import (
	"fmt"
	"tqgin/pkg/gredis"

	//"github.com/gomodule/redigo/redis"
)

func SetUserAccessTocken(user_id, app_id int64, tocken string) bool {

	conn := gredis.RedisConn.Get()
	defer conn.Close()

	key := fmt.Sprintf("access_token_%s", tocken)

	_, err := conn.Do("HMSET", key, "user_id", user_id, "app_id", app_id)

	if err != nil {
		return false
	}
	return true
}
