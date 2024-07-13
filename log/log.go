package log

import (
	"fmt"

	"k8s.io/klog/v2"
)

func Print(level LogLevel, format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	switch level {
	case Warn:
		Msg := fmt.Sprintf("[WARN] %s", logMsg)
		klog.Warning(Msg)
	case Error:
		Msg := fmt.Sprintf("[Error] %s", logMsg)
		klog.Error(Msg)
	case Info:
		fallthrough
	default:
		Msg := fmt.Sprintf("[Info] %s", logMsg)
		klog.Info(Msg)
	}
}
