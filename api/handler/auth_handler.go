package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/model"
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

// LoginHandler handles login requests
// @Summary Login user
// @Description Login user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Login request"
// @Success 200 {object} model.LoginResponse
// @Router /login [post]
func (a *AuthHandler) HandleLogin(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.userService.Verify(req.Username, req.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := a.authService.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(200, model.LoginResponse{
		Token: token,
	})
}

// RegisterHandler handles registration requests
// @Summary Register user
// @Description Register user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "Register request"
// @Success 200 {object} model.RegisterResponse
// @Router /register [post]
func (a *AuthHandler) HandleRegister(ctx *gin.Context) {
	req := &model.RegisterRequest{}
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

	ctx.JSON(200, model.RegisterResponse{
		UserID: user.ID,
	})
}
