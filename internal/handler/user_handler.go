package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/request"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"github.com/nglab-dev/nglab/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

// @Tags Users
// @Summary Get login user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 0 {object} response.Response{data=model.User}
// @Router /user [get]
func (h *UserHandler) GetLoginUser(ctx *gin.Context) {
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

	response.Ok(ctx, user)
}

// @Tags Users
// @Summary Update login user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body dto.LoginUser true "User"
// @Success 0 {object} response.Response{data=model.User}
// @Router /user [patch]
func (h *UserHandler) UpdateLoginUser(ctx *gin.Context) {
	claims := request.GetUserClaims(ctx)
	if claims.UserID == 0 {
		response.Unauthorized(ctx, errors.New("invalid token"))
	}
	var req dto.UpdateLoginUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	user := model.User{
		Nickname: req.Nickname,
		Email:    req.Email,
	}

	err := h.userService.Update(&user)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, user)
}

// @Tags Users
// @Summary List users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data query dto.PaginationParam true "PaginationParam"
// @Success 0 {object} response.Response{data=dto.PaginationResult{Data=[]model.User}}
// @Router /users [get]
func (h *UserHandler) ListUsers(ctx *gin.Context) {
	var req dto.PaginationParam
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	page, err := h.userService.Page(req)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, page)
}

// @Tags Users
// @Summary Create user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body model.User true "User"
// @Success 0 {object} response.Response{data=model.User}
// @Router /users [post]
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.User
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	err := h.userService.Create(&req)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, req)
}

// @Tags Users
// @Summary Get user by id
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 0 {object} response.Response{data=model.User}
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.BadRequest(ctx, errors.New("invalid id"))
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(ctx, errors.New("invalid id"))
		return
	}
	user, err := h.userService.FindByID(uint(i))
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	if user == nil {
		response.ServerError(ctx, errors.New("user not found"))
		return
	}
	response.Ok(ctx, user)
}

// @Tags Users
// @Summary Update user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param data body model.User true "User"
// @Success 0 {object} response.Response{data=model.User}
// @Router /users/{id} [patch]
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.BadRequest(ctx, errors.New("invalid id"))
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(ctx, errors.New("invalid id"))
		return
	}
	var req model.User
	req.ID = uint(i)
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	err = h.userService.Update(&req)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, req)
}
