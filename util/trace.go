package util

import (
    "bytes"
    "github.com/iris-contrib/go.uuid"
    "runtime"
    "strings"
)

func GetGID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	uid, _ := uuid.NewV4()
	ns := strings.Split(uid.String(), "-")[4] + string(b)
	return ns
}

// 获取调用栈上的函数信息(函数名、文件名、调用行数)
func CurrentInfo(callerSkip int) (funcName string, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(callerSkip)
	f := runtime.FuncForPC(pc)
	if ok {
		return f.Name(), file, line
	} else {
		return "", "", -1
	}
}
