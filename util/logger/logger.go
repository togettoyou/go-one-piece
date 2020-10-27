package logger

import (
	"github.com/natefinch/lumberjack"
	"go-one-piece/util/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	filename    string
	errFilename string
)

func init() {
	filename = "log/log.log"
	errFilename = "log/err.log"
}

// 初始化日志
func Setup() {
	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(getConsoleEncoder(), zapcore.Lock(os.Stdout), getLevel()))
	if conf.Config.LogConfig.IsFile {
		file := getFileWriter(filename, conf.Config.LogConfig.MaxSize, conf.Config.LogConfig.MaxBackups, conf.Config.LogConfig.MaxAge)
		errFile := getFileWriter(errFilename, conf.Config.LogConfig.MaxSize, conf.Config.LogConfig.MaxBackups, conf.Config.LogConfig.MaxAge)
		fileEncoder := getFileEncoder()
		cores = append(cores, zapcore.NewCore(fileEncoder, errFile, getHighPriorityLevel()))
		cores = append(cores, zapcore.NewCore(fileEncoder, file, getLowPriorityLevel()))
	}
	lg := zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	defer lg.Sync()
	zap.ReplaceGlobals(lg)
	return
}

var levels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// 日志模式
func getLevel() zapcore.Level {
	if level, ok := levels[conf.Config.LogConfig.Level]; ok {
		return level
	} else {
		return zapcore.DebugLevel
	}
}

func getHighPriorityLevel() zap.LevelEnablerFunc {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && lvl >= getLevel()
	})
	return highPriority
}

func getLowPriorityLevel() zap.LevelEnablerFunc {
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= getLevel()
	})
	return lowPriority
}

// 文件日志格式
func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 控制台日志格式
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
