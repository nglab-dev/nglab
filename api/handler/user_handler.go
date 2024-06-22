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

func (h UserHandler) HandleGet(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, _ := strconv.Atoi(id)

	user, err := h.userService.GetByID(uint(userId))
	if err != nil {
		newResponse(ctx).Error(err.Error())
	}
	newResponse(ctx).OK(user)
}
