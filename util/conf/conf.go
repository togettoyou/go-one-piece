package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type config struct {
	App    app    `yaml:"app"`
	Server server `yaml:"server"`
	Logger logger `yaml:"logger"`
	Mysql  mysql  `yaml:"mysql"`
	Redis  redis  `yaml:"redis"`
}

type app struct {
	JwtSecret string `yaml:"jwtSecret"`
}

type server struct {
	RunMode      string        `yaml:"runMode"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	HttpPort     int           `yaml:"httpPort"`
	Url          string        `yaml:"url"`
	MaxPingCount int           `yaml:"maxPingCount"`
}

type logger struct {
	LoggerLevel            string        `yaml:"loggerLevel"`
	LoggerFile             string        `yaml:"loggerFile"`
	LoggerFileMaxAge       time.Duration `yaml:"loggerFileMaxAge"`
	LoggerFileRotationTime time.Duration `yaml:"loggerFileRotationTime"`
}

type mysql struct {
	Url string `yaml:"url"`
}

type redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

var (
	Config *config
	v      *viper.Viper
)

const defaultConfigFile = "config.yaml"

// 初始化程序配置
func Setup() {
	v = viper.New()
	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	setConfig()
}

func OnConfigChange(run func()) {
	v.OnConfigChange(func(in fsnotify.Event) { run() })
	v.WatchConfig()
}

func setConfig() {
	Config = &config{}
	if err := v.Unmarshal(&Config); err != nil {
		logging().Errorln(defaultConfigFile, err)
	}
	Config.Server.ReadTimeout *= time.Second
	Config.Server.WriteTimeout *= time.Second
	Config.Logger.LoggerFileMaxAge *= time.Hour
	Config.Logger.LoggerFileRotationTime *= time.Hour
}

func Reset() {
	setConfig()
}

var levels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

func logging() *logrus.Entry {
	if level, ok := levels[Config.Logger.LoggerLevel]; ok {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return logrus.WithFields(logrus.Fields{
		"env": Config.Server.RunMode,
	})
}
