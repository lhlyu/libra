package logger

import (
    "github.com/lhlyu/libra/common"
    "github.com/sirupsen/logrus"
    "os"
)

func NewEntry() *logrus.Entry{
    lr := logrus.New()
    out := common.Cfg.GetString("log.out")
    level := common.Cfg.GetString("log.level")
    if out != ""{
        f,err := os.OpenFile(out,os.O_CREATE|os.O_APPEND,0666)
        if err != nil{
            panic(err)
            return nil
        }
        lr.SetOutput(f)
        lr.SetFormatter(new(logrus.JSONFormatter))
    }
    lv,err := logrus.ParseLevel(level)
    if err != nil{
        lv = logrus.InfoLevel
    }
    lr.SetLevel(lv)
    entry := logrus.NewEntry(lr)
    return entry
}
