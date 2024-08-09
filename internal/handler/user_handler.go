package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"github.com/nglab-dev/nglab/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Tags Users
// @Summary List users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data query dto.PaginationRequest true "PageRequest"
// @Success 0 {object} response.Response{data=dto.PageResponse{Data=[]model.User}}
// @Router /users [get]
func (h *UserHandler) ListUsers(ctx *gin.Context) {
	var req dto.PageRequest
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
