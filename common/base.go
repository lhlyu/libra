package common

type base struct {
	Tracker *Tracker
}

func (b *base) SetTracker(tracker *Tracker) {
	b.Tracker = tracker
}

// 基础dao
type BaseDao struct {
	base
}

func (s *BaseDao) Error(err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(3, "error", s.Tracker.GetTraceId(), "repository", err.Error(), param)
	return true
}

func (s *BaseDao) Info(param ...interface{}) {
	Ylog.Log(3, "info", s.Tracker.GetTraceId(), "repository", param...)
}

// 基础服务
type BaseService struct {
	base
}

func (s *BaseService) Error(err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(3, "error", s.Tracker.GetTraceId(), "service", err.Error(), param)
	return true
}

func (s *BaseService) Info(param ...interface{}) {
	Ylog.Log(3, "info", s.Tracker.GetTraceId(), "service", param...)
}

// 基础缓存
type BaseCache struct {
	base
}

func (s *BaseCache) Error(err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(3, "error", s.Tracker.GetTraceId(), "cache", err.Error(), param)
	return true
}

func (s *BaseCache) Info(param ...interface{}) {
	Ylog.Log(3, "info", s.Tracker.GetTraceId(), "cache", param...)
}

// 基础控制器
type BaseController struct {
	base
}

func (s BaseController) Error(traceId string, err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(4, "error", traceId, "controller", err.Error(), param)
	return true
}

func (s BaseController) Info(traceId string, param ...interface{}) {
	Ylog.Log(4, "info", traceId, "controller", param...)
}

type MSF = map[string]interface{}
type MSS = map[string]string
type MSI = map[string]int
type MIF = map[int]interface{}
type MIS = map[int]string
type MII = map[int]int
