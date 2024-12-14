package logs

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type MyHook struct{}

func (h *MyHook) Fire(entry *logrus.Entry) error {
	// 获取调用者的行号
	pc, _, line, ok := runtime.Caller(8)
	if ok {
		entry.Data["line_number"] = line
		entry.Data["function_name"] = runtime.FuncForPC(pc).Name()
	} else {
		entry.Data["line_number"] = -1
	}
	return nil
}

func (h *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
