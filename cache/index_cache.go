package cache

import (
	"github.com/lhlyu/libra/trace"
)

type IndexCache struct {
	BaseCache
}

func NewIndexCache(tracker trace.ITracker) *IndexCache {
	return &IndexCache{
		BaseCache{
			ITracker: tracker,
		},
	}
}

// 获取缓存
func (c *IndexCache) Get(name string) bool {
	c.Info("query cache ", name)
	return false
}

// 设置缓存
func (c *IndexCache) Set(name string) bool {
	return true
}
