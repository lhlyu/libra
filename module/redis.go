package module

import (
	"github.com/go-redis/redis"
	"github.com/lhlyu/libra/common"
	"log"
	"time"
)

type rds struct {
}

func (rds) seq() int {
	return 1 << 1
}

func (rds) SetUp() {
	log.Println("init redis module ->")
	c := &redisConf{}
	if err := common.Cfg.UnmarshalKey("redis", c); err != nil {
		log.Fatal("redis setup is err:", err)
	}
	setRedis(c)
}

// redis模块
var RedisModule = rds{}

type redisConf struct {
	Addr        string `json:"addr"`
	Password    string `json:"password"`
	Database    int    `json:"database"`
	IdleTimeout int    `json:"idleTimeout"`
}

func setRedis(r *redisConf) {

	client := redis.NewClient(&redis.Options{
		Addr:        r.Addr,
		Password:    r.Password,
		DB:          r.Database,
		IdleTimeout: time.Duration(r.IdleTimeout) * time.Second,
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("redis connect is fail,err:", err)
	}
	common.Redis = client
}
