package handler

import "github.com/gin-gonic/gin"

type ResponseBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty" `
}

type response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *response {
	return &response{
		ctx: ctx,
	}
}

func (r *response) OK(data any) {
	r.ctx.JSON(200, ResponseBody{
		Code: 0,
		Msg:  "OK",
		Data: data,
	})
}

func (r *response) BadRequest(msg string) {
	r.ctx.JSON(200, ResponseBody{
		Code: 400,
		Msg:  msg,
		Data: nil,
	})
}

func (r *response) Unauthorized(msg string) {
	r.ctx.JSON(200, ResponseBody{
		Code: 401,
		Msg:  msg,
		Data: nil,
	})
}

func (r *response) Error(msg string) {
	r.ctx.JSON(200, ResponseBody{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
}

func (r *response) ErrorWithCode(code int, msg string) {
	r.ctx.JSON(200, ResponseBody{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
