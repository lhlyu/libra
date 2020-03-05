package module

import (
	"github.com/lhlyu/yutil/v2"
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
	// 工具包打印错误
	yutil.NotIgnore()

}

var InitiateModule = initiate{}
