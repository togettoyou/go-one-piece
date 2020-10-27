package logger

import (
	"github.com/natefinch/lumberjack"
	"go-one-piece/util/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 初始化日志
func Setup() {
	file := getFileWriter(conf.Config.LogConfig.Filename, conf.Config.LogConfig.MaxSize, conf.Config.LogConfig.MaxBackups, conf.Config.LogConfig.MaxAge)
	console := zapcore.Lock(os.Stdout)
	fileEncoder := getFileEncoder()
	consoleEncoder := getConsoleEncoder()
	l := setLevel()
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, file, l),
		zapcore.NewCore(consoleEncoder, console, l),
	)
	lg := zap.New(core, zap.AddCaller())
	defer lg.Sync()
	zap.ReplaceGlobals(lg)
	return
}

// 设置日志模式
func setLevel() *zapcore.Level {
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(conf.Config.LogConfig.Level))
	if err != nil {
		_ = l.UnmarshalText([]byte("debug"))
	}
	return l
}

// 设置文件日志格式
func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 设置控制台日志格式
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 文件日志保存配置
func getFileWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Reset() {
	Setup()
}
