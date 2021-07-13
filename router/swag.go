// +build docs

package router

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-one-server/docs"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
