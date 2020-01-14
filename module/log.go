package module

import (
    "context"
    "github.com/lhlyu/libra/common"
    "github.com/lhlyu/libra/logger"
    "github.com/sirupsen/logrus"
)

type lg struct {
}

func (lg) seq() int {
	return 1 << 0
}

func (lg) SetUp() {
    common.L = logger.NewEntry()
}

var LgModule = lg{}


type loggerKey struct {

}





func WithLogger(ctx context.Context, entry *logrus.Entry) context.Context {
    entry.WithField("id",1)
    return context.WithValue(ctx, loggerKey{}, entry)
}


func GetLogger(ctx context.Context) *logrus.Entry {
    entry := ctx.Value(loggerKey{})
    if entry == nil {
        return logger.NewEntry()
    }
    return entry.(*logrus.Entry)
}

























/**
一个简易的日志格式:
{
  "CREATEDAT": "2019-12-09 12:05:56",                                       // 创建时间
  "FNAME": "github.com/lhlyu/libra/middleware.Log.func1",                     // 调用函数
  "LAYER": "middleware",                                                    // 包/层级
  "LEVEL": "debug",                                                         // 等级
  "PARAM": [                                                                // 信息
    "[1] 127.0.0.1 ▶ GET:/api/articles"
  ],
  "POSITION": "/iyu/middleware/log.go:15",                                  // 调用定位
  "TRACEID": "6f0b160d490a71"                                               // 本次请求唯一ID，用于日志追踪
}
*/
