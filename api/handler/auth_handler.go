package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/api/service"
	"github.com/nglab-dev/nglab/internal/constant"
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
// @Param request body model.LoginRequest true "Login request"
// @Success 200 {object} ResponseBody{data=model.LoginResponse}
// @Router /login [post]
func (a *AuthHandler) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		NewResponse(ctx).BadRequest()
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

	NewResponse(ctx).OK(model.LoginResponse{
		AccessToken: fmt.Sprintf("Bearer %s", token),
	})
}

// @Tags Auth
// @Summary Register user
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "Register request"
// @Success 200 {object} ResponseBody{data=model.RegisterResponse}
// @Router /register [post]
func (a *AuthHandler) Register(ctx *gin.Context) {
	req := &model.RegisterRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		NewResponse(ctx).BadRequest()
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := a.userService.Create(user); err != nil {
		NewResponse(ctx).Error(err.Error())
		return
	}

	NewResponse(ctx).OK(model.RegisterResponse{
		UserID: user.ID,
	})
}

// @Tags Auth
// @Summary Get auth user
// @Accept json
// @Produce json
// @Success 200 {object} ResponseBody{data=model.User}
// @Router /user [get]
func (a *AuthHandler) GetUser(ctx *gin.Context) {
	claims, exist := ctx.Get(constant.CurrentUserKey)
	if !exist {
		NewResponse(ctx).Unauthorized()
		return
	}

	user, err := a.userService.Get(claims.(*model.JWTClaims).UserID)

	if err != nil {
		NewResponse(ctx).Error(err.Error())
		return
	}

	NewResponse(ctx).OK(user)
}
