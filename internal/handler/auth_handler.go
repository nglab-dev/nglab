package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/request"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/service"
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
// @Success 200 {object} response.Response{data=model.LoginResponse}
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

func (h *AuthHandler) GetLoginUser(ctx *gin.Context) {
	user := request.GetLoginUser(ctx)
	response.Ok(ctx, user)
}
