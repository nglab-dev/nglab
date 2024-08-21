package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"github.com/nglab-dev/nglab/internal/service"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService,
	}
}

// @Tags Roles
// @Summary List Roles
// @Security ApiKeyAuth
// @Description List Roles
// @Produce json
// @Param data query dto.PaginationParam true "PaginationParam"
// @Success 0 {object} response.Response{data=dto.PaginationResult{Data=[]model.Role}}
// @Router /roles [get]
func (h *RoleHandler) ListRoles(ctx *gin.Context) {
	var req dto.PaginationParam
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	page, err := h.roleService.Page(req)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, page)
}
