package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/request"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/service"
	"github.com/nglab-dev/nglab/pkg/log"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(authService service.AuthService, userService service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

// @Tags Auth
// @Summary Login user
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Login request"
// @Success 0 {object} response.Response{data=model.LoginResponse}
// @Router /login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req = new(model.LoginRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	user, err := h.userService.FindByUsername(req.Username)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	if user == nil {
		response.Unauthorized(ctx, errors.New("invalid username or password"))
		return
	}

	token, err := h.authService.Login(user)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}

	response.Ok(ctx, model.LoginResponse{
		AccessToken: token,
	})
}

// @Tags Auth
// @Summary Get auth user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=model.User}
// @Router /user [get]
func (h *AuthHandler) GetLoginUser(ctx *gin.Context) {
	claims := request.GetUserClaims(ctx)
	if claims.UserID == 0 {
		response.Unauthorized(ctx, errors.New("invalid token"))
	}
	user, err := h.userService.FindByID(claims.UserID)
	if err != nil {
		response.ServerError(ctx, err)
	}
	if user == nil {
		response.Unauthorized(ctx, errors.New("invalid token"))
	}

	log.Logger.Sugar().Infof("GetLoginUser: %v", user)

	response.Ok(ctx, user)
}
