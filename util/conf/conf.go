package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
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

var Config *config

const defaultConfigFile = "config.yaml"

// 初始化程序配置
func Setup() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		log.Panicf("读取%s异常: %s", defaultConfigFile, err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("热加载%s", e.Name)
		loadConfig(v)
	})
	loadConfig(v)
}

func loadConfig(v *viper.Viper) {
	Config = &config{}
	if err := v.Unmarshal(&Config); err != nil {
		log.Fatalf("解析%s异常: %s", defaultConfigFile, err)
	}
	Config.Server.ReadTimeout *= time.Second
	Config.Server.WriteTimeout *= time.Second
	Config.Logger.LoggerFileMaxAge *= time.Hour
	Config.Logger.LoggerFileRotationTime *= time.Hour
}
