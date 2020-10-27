package handler

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/util/errno"
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
