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
	Level       string `yaml:"level"`
	IsFile      bool   `yaml:"isFile"`
	FilePath    string `yaml:"filePath"`
	ErrFilePath string `yaml:"errFilePath"`
	MaxSize     int    `yaml:"maxSize"`
	MaxAge      int    `yaml:"maxAge"`
	MaxBackups  int    `yaml:"maxBackups"`
}

type mysql struct {
	Dsn         string        `yaml:"dsn"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxOpen     int           `yaml:"maxOpen"`
	MaxLifetime time.Duration `yaml:"maxLifetime"`
	LogMode     string        `yaml:"logMode"`
}

var (
	Config = new(config)
	v      *viper.Viper
)

var DefaultConfigFile string

// Setup 读取配置文件设置
func Setup() {
	v = viper.New()
	v.SetConfigFile(DefaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	setConfig()
}

// OnConfigChange 配置文件热加载回调
func OnConfigChange(run func()) {
	v.OnConfigChange(func(in fsnotify.Event) { run() })
	v.WatchConfig()
}

// setConfig 构造配置文件到Config结构体上
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
