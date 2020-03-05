package module

import (
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/yutil/v2"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type lg struct {
}

func (lg) seq() int {
	return 1 << 0
}

func (lg) SetUp() {
	common.L = NewEntry()
}

// 日志模块
var LgModule = lg{}

func NewEntry() *logrus.Entry {
	lr := logrus.New()
	out := common.Cfg.GetString("log.out")
	level := common.Cfg.GetString("log.level")
	if out != "" {
		dir := path.Dir(out)
		exists := yutil.File.IsExists(dir)
		if !exists {
			os.MkdirAll(dir, os.ModePerm)
		}
		f, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
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
