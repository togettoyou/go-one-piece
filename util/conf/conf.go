package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

type config struct {
	App       app       `yaml:"app"`
	Server    server    `yaml:"server"`
	LogConfig logConfig `yaml:"logConfig"`
	Mysql     mysql     `yaml:"mysql"`
	Redis     redis     `yaml:"redis"`
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

type logConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
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
		zap.L().Error(err.Error())
	}
	Config.Server.ReadTimeout *= time.Second
	Config.Server.WriteTimeout *= time.Second
}

func Reset() {
	setConfig()
}
