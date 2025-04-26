package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"strconv"
	"strings"
	"time"

	"bosh-admin/global"

	"gopkg.in/natefinch/lumberjack.v2"
)

// CustomEncoder 自定义解码器
func CustomEncoder() zapcore.Encoder {
	var encoder zapcore.Encoder
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "M",
		LevelKey:      "L",
		TimeKey:       "T",
		NameKey:       "N",
		CallerKey:     "C",
		FunctionKey:   "F",
		StacktraceKey: zapcore.OmitKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format(global.Config.Log.TimestampFormat))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 解码器格式
	if strings.ToLower(global.Config.Log.Format) == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

// CustomLogger 自定义日志
type CustomLogger struct {
	logger *lumberjack.Logger
	date   string // 日志日期
	suffix string // 日志后缀
}

// Write 自定义日志Write
func (l *CustomLogger) Write(p []byte) (n int, err error) {
	now := time.Now().Local()
	nowDate := now.Format(time.DateOnly)
	if l.date != nowDate {
		defer func(logger *lumberjack.Logger) {
			_ = logger.Close()
		}(l.logger)
		l.date = nowDate
		l.logger.Filename = fmt.Sprintf("%s/%s/%s/%s/%s%s.log", global.Config.Log.RootDir, strconv.Itoa(now.Year()), now.Format("2006-01"), nowDate, nowDate, l.suffix)
	}
	return l.logger.Write(p)
}

// CustomLoggerWriter 自定义日志写入器
func CustomLoggerWriter(suffix string) *CustomLogger {
	now := time.Now().Local()
	nowDate := now.Format(time.DateOnly)
	return &CustomLogger{
		logger: &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s/%s/%s/%s%s.log", global.Config.Log.RootDir, strconv.Itoa(now.Year()), now.Format("2006-01"), nowDate, nowDate, suffix),
			MaxSize:    global.Config.Log.MaxSize,    // 每个日志文件保存的最大尺寸 单位：MB
			MaxAge:     global.Config.Log.MaxAge,     // 文件最多保存多少天
			MaxBackups: global.Config.Log.MaxBackups, // 日志文件最多保存多少个备份
			Compress:   global.Config.Log.Compress,   // 是否压缩
		},
		date:   nowDate,
		suffix: suffix,
	}
}
