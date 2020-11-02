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
}

type logConfig struct {
	Level      string `yaml:"level"`
	IsFile     bool   `yaml:"isFile"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
}

type mysql struct {
	Dsn         string        `yaml:"dsn"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxOpen     int           `yaml:"maxOpen"`
	MaxLifetime time.Duration `yaml:"maxLifetime"`
	LogMode     string        `yaml:"logMode"`
}

type redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

var (
	Config = new(config)
	v      *viper.Viper
)

const defaultConfigFile = "config.yaml"

// 初始化读取配置文件
func Setup() {
	v = viper.New()
	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	setConfig()
}

// 配置文件热加载回调
func OnConfigChange(run func()) {
	v.OnConfigChange(func(in fsnotify.Event) { run() })
	v.WatchConfig()
}

// 构造配置文件到Config结构体上
func setConfig() {
	if err := v.Unmarshal(&Config); err != nil {
		zap.L().Error(err.Error())
	}
	Config.Server.ReadTimeout *= time.Second
	Config.Server.WriteTimeout *= time.Second
	Config.Mysql.MaxLifetime *= time.Minute
}

func Reset() {
	setConfig()
}
