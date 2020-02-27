package trace

import (
	"fmt"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/libra/util"
	"github.com/sirupsen/logrus"
)

const TRACKER = "TRACKER"

type BaseTracker struct {
	ITracker
}

func NewBaseTracker(tracker ITracker) BaseTracker {
	return BaseTracker{tracker}
}

type ITracker interface {
	Info(v ...interface{})             // 普通信息
	Debug(v ...interface{})            // 调试信息
	Error(err error, v ...interface{}) // 错误信息
	Alert(v ...interface{})            // 消息发送到邮件或者钉钉，可自己定义
}

type Tracker struct {
	entry *logrus.Entry
}

func NewTracker() *Tracker {
	traceId := util.GetGID()
	entry := common.L
	if entry == nil {
		entry = logrus.NewEntry(logrus.StandardLogger())
	}
	return &Tracker{
		entry: entry.WithFields(logrus.Fields{
			"traceId": traceId,
		}),
	}
}

func (t *Tracker) Info(v ...interface{}) {
	t.entry.Infoln(v...)
}

func (t *Tracker) Debug(v ...interface{}) {
	funcName, _, line := util.CurrentInfo(2)
	t.entry.WithFields(logrus.Fields{
		"position": fmt.Sprintf("%s:%d", funcName, line),
	}).Debugln(v...)
}

func (t *Tracker) Error(err error, v ...interface{}) {
	if err == nil {
		return
	}
	t.entry.WithFields(logrus.Fields{
		"error": err.Error(),
	}).Errorln(v...)
}

func (t *Tracker) Alert(v ...interface{}) {
	// 可以发送到邮件 钉钉 ...
}
