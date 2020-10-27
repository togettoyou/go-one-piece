package examples

import (
	"github.com/gin-gonic/gin"
	. "go-one-piece/handler"
)

type Body struct {
	Username string `json:"username" form:"username" validate:"required,checkUsername"`
}

func GetExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	var body Body
	if !g.ParseQueryRequest(&body) {
		return
	}
	if g.HasError(nil) {
		return
	}
	g.OkWithDataResponse(body)
}
