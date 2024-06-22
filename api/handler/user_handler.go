package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService}
}

// @Tags User
// @Summary User Get By ID
// @Produce application/json
// @Param id path int true "user id"
// @Success 200 {object} model.User
// @router /users/{id} [get]
func (h UserHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, _ := strconv.Atoi(id)

	user, err := h.userService.Get(uint(userId))
	if err != nil {
		NewResponse(ctx).Error(err.Error())
	}
	NewResponse(ctx).OK(user)
}
