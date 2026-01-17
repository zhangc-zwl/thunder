package database

import (
	"github.com/zhangc-zwl/thunder/config"
	"github.com/zhangc-zwl/thunder/db"
)

var (
	RedisCli *db.Redis
)

func InitRedis(redisConf *config.Redis) {
	if redisConf == nil {
		return
	}

	r := db.Redis{}
	r.Init(redisConf)
	RedisCli = &r
}