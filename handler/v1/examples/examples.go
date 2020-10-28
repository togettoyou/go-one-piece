package examples

import (
	"github.com/gin-gonic/gin"
	. "go-one-piece/handler"
)

type ValidateBody struct {
	Validate string `json:"validate" form:"validate" validate:"required,checkUsername"`
}

func GetExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	var body ValidateBody
	if !g.ParseQueryRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}
