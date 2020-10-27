package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go-one-piece/util"
	"go-one-piece/util/conf"
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

func (g *Gin) ParseUriRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindUri(request); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request)
}

func (g *Gin) ParseQueryRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindQuery(request); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request)
}

func (g *Gin) ParseJSONRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindJSON(request); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request)
}

func (g *Gin) ParseFormRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindWith(request, binding.Form); err != nil {
		g.SendNoDataResponse(errno.ErrBind)
		return false
	}
	return validatorData(g, request)
}

func validatorData(g *Gin, request interface{}) bool {
	if err := myValidator.V.Struct(request); err != nil {
		var errStr string
		switch err.(type) {
		case validator.ValidationErrors:
			// 正式版不显示参数异常详情
			if conf.Config.Server.RunMode == util.ReleaseMode {
				g.SendNoDataResponse(errno.ErrValidation)
				return false
			}
			errStr = myValidator.TranslateErr(err.(validator.ValidationErrors))
		default:
			errStr = errors.New("unknown error").Error()
		}
		g.SendNoDataResponse(errno.New(errno.ErrValidation, err).Add(errStr))
		return false
	}
	return true
}
