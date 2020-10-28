package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go-one-piece/util/errno"
	myValidator "go-one-piece/util/validator"
	"net/http"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) SendResponse(err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	g.Ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
}

func (g *Gin) SendNoDataResponse(err error) {
	g.SendResponse(err, map[string]interface{}{})
}

func (g *Gin) OkResponse() {
	g.Ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "OK",
		Data: map[string]interface{}{},
	})
}

func (g *Gin) OkWithMsgResponse(msg string) {
	g.Ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: map[string]interface{}{},
	})
}

func (g *Gin) OkWithDataResponse(data interface{}) {
	g.Ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "OK",
		Data: data,
	})
}

func (g *Gin) OkCustomResponse(msg string, data interface{}) {
	g.Ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func (g *Gin) HasError(err error) bool {
	if err != nil {
		g.SendNoDataResponse(err)
		return true
	}
	return false
}

func (g *Gin) ParseUriRequest(request interface{}, hideDetails ...bool) bool {
	if err := g.Ctx.ShouldBindUri(request); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request, len(hideDetails) > 0 && hideDetails[0])
}

func (g *Gin) ParseQueryRequest(request interface{}, hideDetails ...bool) bool {
	if err := g.Ctx.ShouldBindQuery(request); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request, len(hideDetails) > 0 && hideDetails[0])
}

func (g *Gin) ParseJSONRequest(request interface{}, hideDetails ...bool) bool {
	if err := g.Ctx.ShouldBindJSON(request); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request, len(hideDetails) > 0 && hideDetails[0])
}

func (g *Gin) ParseFormRequest(request interface{}, hideDetails ...bool) bool {
	if err := g.Ctx.ShouldBindWith(request, binding.Form); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request, len(hideDetails) > 0 && hideDetails[0])
}

// hideDetails可选择隐藏参数校验详细信息
func validatorData(g *Gin, request interface{}, hideDetails bool) bool {
	if err := myValidator.V.Struct(request); err != nil {
		var eno error
		switch err.(type) {
		case validator.ValidationErrors:
			if !hideDetails {
				g.SendResponse(errno.ErrValidation, myValidator.TranslateErrData(err.(validator.ValidationErrors)))
				return false
			}
			eno = errno.ErrValidation
		default:
			eno = err
		}
		g.SendNoDataResponse(eno)
		return false
	}
	return true
}
