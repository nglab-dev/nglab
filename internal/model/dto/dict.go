package dto

type DictPaginationParam struct {
	PaginationParam
	Type int `json:"type" form:"type" binding:"required"`
}
