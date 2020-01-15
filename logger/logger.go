package logger

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/libra/util"
	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

func WithLogger(ctx context.Context) context.Context {
	traceId := util.GetGID()
	entry := common.L
	if entry == nil {
		entry = logrus.NewEntry(logrus.StandardLogger())
	}
	return context.WithValue(ctx, loggerKey{}, entry.WithFields(logrus.Fields{
		"traceId": traceId,
	}))
}

func GetLogger(ctx context.Context) *logrus.Entry {
	entry := ctx.Value(loggerKey{})
	if entry == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	return entry.(*logrus.Entry)
}

// for iris context
func Log(ctx iris.Context) *logrus.Entry {
	funcName, _, line := util.CurrentInfo(2)
	entry := GetLogger(ctx.Request().Context())
	if common.Cfg.GetString("log.level") == "debug" {
		entry = entry.WithFields(logrus.Fields{
			"position": fmt.Sprintf("%s:%d", funcName, line),
		})
	}
	return entry
}
