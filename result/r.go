package result

import (
	"fmt"
	"github.com/lhlyu/libra/common"
)

// 统一响应处理
type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewR(code int, msg string) *R {
	return &R{
		Code: code,
		Msg:  msg,
	}
}

func (r *R) IsSuccess() bool {
	if r.Code == 0 {
		return true
	}
	return false
}

func (r *R) String() string {
	return fmt.Sprintf("code=%d,msg=%s,data=%v", r.Code, r.Msg, r.Data)
}

func (r *R) WithData(data interface{}) *R {
	nr := NewR(r.Code, r.Msg)
	nr.Data = data
	return nr
}

func (r *R) WithPage(data interface{}, page *common.Page) *R {
	nr := NewR(r.Code, r.Msg)
	nr.Data = map[string]interface{}{
		"list": data,
		"page": page,
	}
	return nr
}

func (r *R) WithMsg(msg ...interface{}) *R {
	nr := NewR(r.Code, r.Msg)
	if nr.Msg == "" {
		nr.Msg = fmt.Sprint(msg...)
	} else {
		nr.Msg += ":" + fmt.Sprint(msg...)
	}
	return nr
}

var (
	Error   = NewR(-1, "系统异常")
	Success = NewR(0, "成功")
	Failure = NewR(1, "失败")

	EmptyData     = NewR(1000, "数据为空")
	NotExistsData = NewR(1001, "数据不存在")
	IllegalParam  = NewR(1002, "参数不合法")

	QueryError  = NewR(10001, "查询失败")
	InsertError = NewR(10002, "插入失败")
	UpdateError = NewR(10003, "更新失败")
	DeleteError = NewR(10004, "删除失败")
)
