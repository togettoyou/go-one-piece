package examples

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
)

// @Tags examples
// @Summary Get请求
// @Produce  json
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/get [get]
func GetExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	g.OkResponse()
}
