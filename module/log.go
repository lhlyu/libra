package module

import (
	"github.com/lhlyu/libra/common"
	"github.com/sirupsen/logrus"
	"os"
)

type lg struct {
}

func (lg) seq() int {
	return 1 << 0
}

func (lg) SetUp() {
	common.L = NewEntry()
}

var LgModule = lg{}

func NewEntry() *logrus.Entry {
	lr := logrus.New()
	out := common.Cfg.GetString("log.out")
	level := common.Cfg.GetString("log.level")
	if out != "" {
		f, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
			return nil
		}
		lr.SetOutput(f)
		lr.SetFormatter(new(logrus.JSONFormatter))
	}
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		lv = logrus.InfoLevel
	}
	lr.SetLevel(lv)

	// 这里可以给日志加 hook

	entry := logrus.NewEntry(lr)
	return entry
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
