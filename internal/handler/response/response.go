package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
		Data: nil,
	})
}

func Unauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, Response{
		Code: http.StatusUnauthorized,
		Msg:  err.Error(),
		Data: nil,
	})
}

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Msg:  err.Error(),
		Data: nil,
	})
}
