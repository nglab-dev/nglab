package model

import "github.com/nglab-dev/nglab/internal/model/dto"

type Role struct {
	BaseModel
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Sort   int    `json:"sort"`
	Enable int    `json:"enable"`
}

type Roles []Role

type RoleQueryParam struct {
	dto.PaginationParam
}

func (r *Role) TableName() string {
	return "sys_role"
}

func (r Roles) ToNames() []string {
	names := make([]string, len(r))
	for i, item := range r {
		names[i] = item.Name
	}
	return names
}

func (r Roles) ToIDs() []uint {
	ids := make([]uint, len(r))
	for i, item := range r {
		ids[i] = item.ID
	}
	return ids
}
