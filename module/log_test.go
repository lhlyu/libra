package module

import (
    "context"
    "github.com/sirupsen/logrus"
    "testing"
)

func TestLog(t *testing.T){
    entry := logrus.NewEntry(logrus.StandardLogger())
    ctx := context.Background()
    WithLogger(ctx,entry)
    GetLogger(ctx).WithField("tt",12).Infoln("xxx")
    GetLogger(ctx).WithField("tt",12).Infoln("yyy")
}
