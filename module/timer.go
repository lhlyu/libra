package module

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

type timer struct {
}

func (timer) seq() int {
	return 1 << 5
}

func (timer) SetUp() {
	log.Println("init timer module ->")
	c := cron.New()

	//c.AddFunc("0 0/5 * * * *", task)

	c.Start()
	log.Println("timer is running...")
}

// 定时器模块
var TimerModule = timer{}

func task() {
	fmt.Printf("%s : testing...\n", time.Now().Format("2006-01-02 15:04:05"))
}
