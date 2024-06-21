package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/api/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) HandleLogin(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用登录服务进行验证

	// 验证通过后，生成 JWT 令牌并返回给客户端

	// 示例返回
	ctx.JSON(200, gin.H{"token": "your_jwt_token"})
}
