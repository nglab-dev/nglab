package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"github.com/nglab-dev/nglab/internal/service"
)

type DictHandler struct {
	dictService service.DictService
}

func NewDictHandler(dictService service.DictService) *DictHandler {
	return &DictHandler{
		dictService,
	}
}

// @Tags dicts
// @Summary List Dict
// @Security ApiKeyAuth
// @Produce  json
// @Param data query dto.DictPaginationParam true "DictPaginationParam"
// @Success 0 {object} response.Response{data=dto.PaginationResult{Data=model.Dicts}}
// @Router /dicts [get]
func (h *DictHandler) ListDicts(ctx *gin.Context) {
	var req dto.DictPaginationParam
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	page, err := h.dictService.Page(req)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, page)
}

// @Tags dicts
// @Summary List Dict Types
// @Security ApiKeyAuth
// @Produce  json
// @Success 0 {object} response.Response{data=model.DictTypes}
// @Router /dicts/types [get]
func (h *DictHandler) ListDictTypes(ctx *gin.Context) {
	types, err := h.dictService.Types()
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, types)
}

// @Tags dicts
// @Summary Create Dict Type
// @Security ApiKeyAuth
// @Produce  json
// @Param data body dto.DictType true "DictType"
// @Success 0 {object} response.Response{data=model.DictType}
// @Router /dicts/types [post]
func (h *DictHandler) CreateDictType(ctx *gin.Context) {
	var req model.DictType
	if err := ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	err := h.dictService.CreateType(req)
	if err != nil {
		response.ServerError(ctx, err)
		return
	}
	response.Ok(ctx, nil)
}
