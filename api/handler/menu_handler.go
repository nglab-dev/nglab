package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/api/service"
)

type MenuHandler struct {
	menuService service.MenuService
}

func NewMenuHandler(menuService service.MenuService) MenuHandler {
	return MenuHandler{menuService}
}

// @Tags Menu
// @Summary
// @Produce json
// @Param request body model.MenuCreateRequest true "Menu request"
// @Success 200 {object} ResponseBody{data=[]model.Menu}
// @Router /menus [post]
func (h *MenuHandler) Create(ctx *gin.Context) {
	req := &model.Menu{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		NewResponse(ctx).BadRequest()
		return
	}
	err := h.menuService.Create(req)
	if err != nil {
		NewResponse(ctx).Error(err.Error())
		return
	}
	NewResponse(ctx).OK(req)
}

// @Tags Menu
// @Summary
// @Produce json
// @Success 200 {object} ResponseBody{data=[]model.MenuTree}
// @Router /menus [get]
func (h *MenuHandler) List(ctx *gin.Context) {
	menus, err := h.menuService.List()
	if err != nil {
		NewResponse(ctx).Error(err.Error())
		return
	}
	NewResponse(ctx).OK(menus)
}
