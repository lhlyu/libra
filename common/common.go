package common

import (
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

var (
	Cfg   *viper.Viper
	Email *yuEmail
	L  *logrus.Entry
)
