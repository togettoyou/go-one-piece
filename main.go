package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"go-one-server/model"
	"go-one-server/router"
	"go-one-server/util"
	"go-one-server/util/conf"
	"go-one-server/util/logger"
	"go-one-server/util/tools"
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
	v = pflag.BoolP("version", "v", false, "show version info.")
)

// @title go-one-server
// @version 1.0
// @description 基于Gin进行快速构建RESTful API 服务的脚手架
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
	setup()
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
	timeLocal := time.FixedZone("CST", 8*3600)
	time.Local = timeLocal
	zap.L().Info(time.Now().Format(tools.TimeFormat))
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
	fmt.Printf(`swagger文档地址：http://localhost%s/swagger/index.html
   ____   ____             ____   ____   ____             ______ ______________  __ ___________ 
  / ___\ /  _ \   ______  /  _ \ /    \_/ __ \   ______  /  ___// __ \_  __ \  \/ // __ \_  __ \
 / /_/  >  <_> ) /_____/ (  <_> )   |  \  ___/  /_____/  \___ \\  ___/|  | \/\   /\  ___/|  | \/
 \___  / \____/           \____/|___|  /\___  >         /____  >\___  >__|    \_/  \___  >__|   
/_____/                              \/     \/               \/     \/                 \/       

`, httpPort)
}
