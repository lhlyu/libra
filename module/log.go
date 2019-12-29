package module

import "github.com/lhlyu/libra/common"

type lg struct {
}

func (lg) seq() int {
	return 1 << 0
}

func (lg) SetUp() {
	common.Ylog = common.NewYlog(common.Cfg.GetString("log.level"),
		common.Cfg.GetString("log.timeFormat"),
		common.Cfg.GetString("log.outFile"))
}

var LgModule = lg{}

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
