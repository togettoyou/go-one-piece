package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	SlowThreshold time.Duration
	ZapLogger     *zap.Logger
	LogLevel      gormlogger.LogLevel
}

func New(zapLogger *zap.Logger) gormlogger.Interface {
	return &GormLogger{
		SlowThreshold: 200 * time.Millisecond,
		ZapLogger:     zapLogger,
		LogLevel:      gormlogger.Warn,
	}
}

func (gl *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *gl
	newLogger.LogLevel = level
	return &newLogger
}

func (gl GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= gormlogger.Info {
		gl.ZapLogger.Sugar().Infof(msg, data...)
	}
}

func (gl GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= gormlogger.Warn {
		gl.ZapLogger.Sugar().Warnf(msg, data...)
	}
}

func (gl GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= gormlogger.Error {
		gl.ZapLogger.Sugar().Errorf(msg, data...)
	}
}

func (gl GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel > gormlogger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && gl.LogLevel >= gormlogger.Error:
			sql, rows := fc()
			gl.ZapLogger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		case elapsed > gl.SlowThreshold && gl.SlowThreshold != 0 && gl.LogLevel >= gormlogger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", gl.SlowThreshold)
			gl.ZapLogger.Warn("trace", zap.String("tips", slowLog), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		case gl.LogLevel == gormlogger.Info:
			sql, rows := fc()
			gl.ZapLogger.Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		}
	}
}
