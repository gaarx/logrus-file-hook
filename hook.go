package logrus_file_hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(10); ok {
		entry.Data["source"] = fmt.Sprintf("%s:%v->%s", path.Base(file), line, path.Base(runtime.FuncForPC(pc).Name()))
	}
	return nil
}
