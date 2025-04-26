package log

import "bosh-admin/global"

// Debug debug日志
func Debug(args ...interface{}) {
	global.Logger.Debug(args...)
}

// Info info日志
func Info(args ...interface{}) {
	global.Logger.Info(args...)
}

// Warn warn日志
func Warn(args ...interface{}) {
	global.Logger.Warn(args...)
}

// Error error日志
func Error(args ...interface{}) {
	global.Logger.Error(args...)
}

// Panic panic日志
func Panic(args ...interface{}) {
	global.Logger.Panic(args...)
}

// Fatal fatal日志
func Fatal(args ...interface{}) {
	global.Logger.Fatal(args...)
}
