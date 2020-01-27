package logger

import (
	"context"
	"fmt"
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

func GetLogger(ctx context.Context, skip int) *logrus.Entry {
	lg := ctx.Value(loggerKey{})
	var entry *logrus.Entry
	if lg == nil {
		entry = logrus.NewEntry(logrus.StandardLogger())
	} else {
		entry = lg.(*logrus.Entry)
	}
	funcName, _, line := util.CurrentInfo(2 + skip)
	if common.Cfg.GetString("log.level") == "debug" {
		entry = entry.WithFields(logrus.Fields{
			"position": fmt.Sprintf("%s:%d", funcName, line),
		})
	}
	return entry
}
