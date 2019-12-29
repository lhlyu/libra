package common

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/pio"
	"github.com/lhlyu/libra/util"
	"github.com/lhlyu/yutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type ylog struct {
	g          *golog.Logger
	outWay     string
	timeFormat string
	level      string
}

func NewYlog(level, timeFormat, outFile string) *ylog {
	g := golog.New()
	yg := &ylog{
		timeFormat: yutil.DEFAULT_TIME_FORMAT,
		g:          g,
	}
	if level != "" {
		yg.g.SetLevel(level)
		yg.level = level
	}
	if timeFormat != "" {
		yg.g.SetTimeFormat("")
		yg.timeFormat = timeFormat
	}
	if outFile != "" {
		// 写入文件不打印前面的日志等级标志
		yg.g.Hijack(func(ctx *pio.Ctx) {
			l, ok := ctx.Value.(*golog.Log)
			if !ok {
				ctx.Next()
				return
			}
			line := golog.GetTextForLevel(golog.DisableLevel, ctx.Printer.IsTerminal)
			if line != "" {
				line += " "
			}
			if t := l.FormatTime(); t != "" {
				line += t + " "
			}
			line += l.Message
			var b []byte
			if pref := l.Logger.Prefix; len(pref) > 0 {
				b = append(pref, []byte(line)...)
			} else {
				b = []byte(line)
			}
			ctx.Store(b, nil)
			ctx.Next()
		})
		fl, e := os.OpenFile(outFile, os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			panic(e)
		}
		g.SetOutput(fl)
	}
	yg.g.Println("service is running : ", time.Now().Format(timeFormat))
	return yg
}

// 封装一个简易的日志追踪 根据traceId追踪日志
func (y *ylog) Log(skip int, level, traceId, layer string, param ...interface{}) {
	if y == nil {
		return
	}
	fname, fileName, line := util.CurrentInfo(skip)
	m := map[string]interface{}{
		"1.traceid":   traceId,
		"2.level":     level,
		"3.layer":     layer,
		"4.fname":     fname,
		"6.position":  strings.Join([]string{fileName, strconv.Itoa(line)}, ":"),
		"0.createdat": time.Now().Format(y.timeFormat),
	}
	if len(param) > 0 {
		m["5.param"] = fmt.Sprint(param...)
	}
	bytes, _ := json.Marshal(m)
	y.g.Log(golog.ParseLevel(level), string(bytes))

	// 统一处理错误日志，发送邮件
}
