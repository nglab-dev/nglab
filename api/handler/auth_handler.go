package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/api/schema"
	"github.com/nglab-dev/nglab/api/service"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(authService service.AuthService, userService service.UserService) AuthHandler {
	return AuthHandler{
		authService,
		userService,
	}
}

// @Tags Auth
// @Summary Login user
// @Accept json
// @Produce json
// @Param request body schema.LoginRequest true "Login request"
// @Success 200 {object} ResponseBody{data=schema.LoginResponse}
// @Router /login [post]
func (a *AuthHandler) HandleLogin(ctx *gin.Context) {
	var req schema.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		NewResponse(ctx).BadRequest(err.Error())
		return
	}

	user, err := a.userService.Verify(req.Username, req.Password)
	if err != nil {
		NewResponse(ctx).Error(err.Error())
		return
	}

	// Generate JWT token
	token, err := a.authService.GenerateToken(user.ID)
	if err != nil {
		NewResponse(ctx).Error(err.Error())
		return
	}

	NewResponse(ctx).OK(schema.LoginResponse{
		AccessToken: fmt.Sprintf("Bearer %s", token),
	})
}

// @Tags Auth
// @Summary Register user
// @Accept json
// @Produce json
// @Param request body schema.RegisterRequest true "Register request"
// @Success 200 {object} ResponseBody{data=schema.RegisterResponse}
// @Router /register [post]
func (a *AuthHandler) HandleRegister(ctx *gin.Context) {
	req := &schema.RegisterRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := a.userService.Create(user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, schema.RegisterResponse{
		UserID: user.ID,
	})
}
