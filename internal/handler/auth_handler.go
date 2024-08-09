package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/request"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model/dto"
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
// @Param request body dto.LoginRequest true "Login request"
// @Success 0 {object} response.Response{data=dto.LoginResponse}
// @Router /login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req = new(dto.LoginRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	user, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		response.Unauthorized(ctx, err)
		return
	}

	response.Ok(ctx, dto.LoginResponse{
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

// @Tags Auth
// @Summary Logout user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{msg=string}
// @Router /logout [post]
func (h *AuthHandler) Logout(ctx *gin.Context) {
	clamis := request.GetUserClaims(ctx)
	h.authService.Logout(clamis)
	response.Ok(ctx, nil)
}
