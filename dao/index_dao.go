package dao

import (
	"github.com/lhlyu/libra/trace"
	"math/rand"
	"time"
)

type IndexDao struct {
	trace.BaseTracker
}

func NewIndexDao(tracker trace.ITracker) *IndexDao {
	return &IndexDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

func (d *IndexDao) Query(name string) int {
	// 这里假装查询了
	sql := "select * from user where name = ?"
	// 打印日志
	d.Debug(sql, name)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(36)
}
