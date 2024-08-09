package handler

import "github.com/nglab-dev/nglab/internal/service"

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService,
	}
}
