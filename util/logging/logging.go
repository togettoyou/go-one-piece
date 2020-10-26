package logging

import (
	"github.com/sirupsen/logrus"
	"go-one-piece/util/conf"
)

func Setup() {
	setLevel()
}

var levels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

func setLevel() {
	if level, ok := levels[conf.Config.Logger.LoggerLevel]; ok {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Get() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"env": conf.Config.Server.RunMode,
	})
}

func Reset() {
	setLevel()
}
