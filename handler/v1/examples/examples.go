package examples

import (
	"github.com/gin-gonic/gin"
	. "go-one-piece/handler"
)

func GetExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	g.OkResponse()
}
