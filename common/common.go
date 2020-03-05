package common

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Cfg   *viper.Viper
	DB    *sqlx.DB
	Redis *redis.Client
	L     *logrus.Entry
)
