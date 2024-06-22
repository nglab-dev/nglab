package handler

import "github.com/gin-gonic/gin"

type Response struct {
	// 响应状态码
	Code int `json:"code"`
	// 响应消息
	Message string `json:"message"`
	// 响应数据
	Data any `json:"data"`
}

type response struct {
	ctx *gin.Context
}

func newResponse(ctx *gin.Context) *response {
	return &response{
		ctx: ctx,
	}
}

func (r *response) OK(data any) {
	r.ctx.JSON(200, Response{
		Code:    0,
		Message: "OK",
		Data:    data,
	})
}

func (r *response) BadRequest(message string) {
	r.ctx.JSON(200, Response{
		Code:    400,
		Message: message,
		Data:    nil,
	})
}

func (r *response) Unauthorized(message string) {
	r.ctx.JSON(200, Response{
		Code:    401,
		Message: message,
		Data:    nil,
	})
}

func (r *response) Error(message string) {
	r.ctx.JSON(200, Response{
		Code:    1,
		Message: message,
		Data:    nil,
	})
}

func (r *response) CustomError(code int, message string) {
	r.ctx.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
