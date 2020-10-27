package logger

import (
	"github.com/natefinch/lumberjack"
	"go-one-piece/util/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Setup() {
	file := getLogWriter(conf.Config.LogConfig.Filename, conf.Config.LogConfig.MaxSize, conf.Config.LogConfig.MaxBackups, conf.Config.LogConfig.MaxAge)
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

func setLevel() *zapcore.Level {
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(conf.Config.LogConfig.Level))
	if err != nil {
		_ = l.UnmarshalText([]byte("debug"))
	}
	return l
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
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
