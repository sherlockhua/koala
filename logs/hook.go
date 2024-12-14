package logs

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

type MyHook struct{}

func (h *MyHook) Fire(entry *logrus.Entry) error {
	// 获取调用者的行号
	pc, _, line, ok := runtime.Caller(8)
	if ok {
		entry.Data["line"] = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
	} else {
		entry.Data["line"] = "unknown:-1"
	}
	return nil
}

func (h *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
