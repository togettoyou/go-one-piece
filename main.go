package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/togettoyou/gtools"
	"go-one-server/model"
	"go-one-server/router"
	"go-one-server/util"
	"go-one-server/util/conf"
	"go-one-server/util/logger"
	"go-one-server/util/validator"
	"go-one-server/util/version"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func setup() {
	conf.Setup()
	logger.Setup()
	validator.Setup()
	model.Setup()
}

var (
	v      = pflag.BoolP("version", "v", false, "显示版本信息")
	config = pflag.StringP("config", "c", "conf/config.yaml", "指定配置文件路径")
)

// @title go-server
// @version 1.0
// @description 基于Gin进行快速构建 RESTFUL API 服务的项目模板
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	pflag.Parse()
	if *v {
		info := version.Get()
		marshalled, err := json.MarshalIndent(&info, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}
	conf.DefaultConfigFile = *config
	setup()
	defer func() {
		zap.L().Sync()
		zap.S().Sync()
	}()
	startServer()
	reload := make(chan int, 1)
	conf.OnConfigChange(func() { reload <- 1 })
	for {
		select {
		case <-reload:
			util.Reset()
		}
	}
}

func startServer() {
	time.Local = time.FixedZone("CST", 8*3600)
	zap.L().Info(time.Now().Format(gtools.TimeFormat))
	gin.SetMode(conf.Config.Server.RunMode)
	httpPort := fmt.Sprintf(":%d", conf.Config.Server.HttpPort)
	server := &http.Server{
		Addr:           httpPort,
		Handler:        router.InitRouter(),
		ReadTimeout:    conf.Config.Server.ReadTimeout,
		WriteTimeout:   conf.Config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	if router.HasDocs() {
		fmt.Printf(`swagger 文档地址 : http://%s%s/swagger/index.html
   ____   ____             ____   ____   ____             ______ ______________  __ ___________ 
  / ___\ /  _ \   ______  /  _ \ /    \_/ __ \   ______  /  ___// __ \_  __ \  \/ // __ \_  __ \
 / /_/  >  <_> ) /_____/ (  <_> )   |  \  ___/  /_____/  \___ \\  ___/|  | \/\   /\  ___/|  | \/
 \___  / \____/           \____/|___|  /\___  >         /____  >\___  >__|    \_/  \___  >__|   
/_____/                              \/     \/               \/     \/                 \/       

`, gtools.GetCurrentIP().String(), httpPort)
	}
}
