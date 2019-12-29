package module

import (
	"github.com/lhlyu/yutil"
	"log"
)

// 启动时执行
type initiate struct {
}

func (initiate) seq() int {
	return 1 << 4
}

func (initiate) SetUp() {
	log.Println("init initiate module ->")
	// 工具包不忽略错误
	yutil.NotIgnore()
}



var InitiateModule = initiate{}
