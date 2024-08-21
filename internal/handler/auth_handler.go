package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/request"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"github.com/nglab-dev/nglab/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(authService service.AuthService, userService service.UserService) *AuthHandler {
	return &AuthHandler{
		authService,
		userService,
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
