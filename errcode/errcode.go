package errcode

import (
	"fmt"
	"github.com/lhlyu/libra/common"
)

type ErrCode struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e *ErrCode) IsSuccess() bool {
	if e.Code == 0 {
		return true
	}
	return false
}

func (e *ErrCode) new() *ErrCode {
	return &ErrCode{
		Code: e.Code,
		Data: e.Data,
		Msg:  errCodeMap[e.Code],
	}
}

func (e *ErrCode) String() string {
	return fmt.Sprintf("code=%d,msg=%s,data=%v", e.Code, e.Msg, e.Data)
}

func (e *ErrCode) GetErrCode() *ErrCode {
	return e
}

func (e *ErrCode) WithData(data interface{}) *ErrCode {
	ne := e.new()
	ne.Data = data
	return ne
}

func (e *ErrCode) WithPage(page *common.Page, data interface{}) *ErrCode {
	ne := e.new()
	ne.Data = map[string]interface{}{
		"page": page,
		"list": data,
	}
	return ne
}

func (e *ErrCode) AddMsg(msg ...interface{}) *ErrCode {
	ne := e.new()
	ne.Msg += ":" + fmt.Sprint(msg...)
	return ne
}

func NewErrcode(code int, data interface{}) *ErrCode {
	return &ErrCode{
		Code: code,
		Data: data,
		Msg:  errCodeMap[code],
	}
}

const (
	ERROR   = iota - 1 // -1
	SUCCESS            // 0
	FAILURE
)

const (
	NOLOGIN      = 402
	NOPERMISSION = 403
	NOFOUND      = 404
)

const (
	EMPTY_DATA = iota + 1000
	EXISTS_DATA
	NO_EXISTS_DATA
	ILLEGAL_PARAM
)

const (
	QUERY_ERROR = iota + 10000
	UPDATE_ERROR
	INSERT_ERROR
	DELETE_ERROR
)

const (
	NO_OPEN_CMNT = iota + 100000
)

// 常用
var (
	Error   = NewErrcode(ERROR, nil)
	Success = NewErrcode(SUCCESS, nil)
	Failure = NewErrcode(FAILURE, nil)

	NoLogin      = NewErrcode(NOLOGIN, nil)
	NoPermission = NewErrcode(NOPERMISSION, nil)
	NofoundError = NewErrcode(NOFOUND, nil)

	EmptyData    = NewErrcode(EMPTY_DATA, nil)
	ExsistData   = NewErrcode(EXISTS_DATA, nil)
	NoExsistData = NewErrcode(NO_EXISTS_DATA, nil)
	IllegalParam = NewErrcode(ILLEGAL_PARAM, nil)

	QueryError  = NewErrcode(QUERY_ERROR, nil)
	UpdateError = NewErrcode(UPDATE_ERROR, nil)
	InsertError = NewErrcode(INSERT_ERROR, nil)
	DeleteError = NewErrcode(DELETE_ERROR, nil)

	NoOpenCmntError = NewErrcode(NO_OPEN_CMNT, nil)
)

var errCodeMap = map[int]string{
	ERROR:   "异常",
	SUCCESS: "成功",
	FAILURE: "失败",

	NOLOGIN:      "没有登录",
	NOPERMISSION: "没有权限",
	NOFOUND:      "404",

	EMPTY_DATA:     "数据为空",
	EXISTS_DATA:    "数据已存在",
	NO_EXISTS_DATA: "数据不存在",
	ILLEGAL_PARAM:  "参数不合法",

	QUERY_ERROR:  "查询错误",
	UPDATE_ERROR: "更新错误",
	INSERT_ERROR: "添加错误",
	DELETE_ERROR: "删除错误",

	NO_OPEN_CMNT: "暂未开放评论",
}
