package service

import (
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"gorm.io/gorm"
)

type RoleService interface {
	Page(query dto.PaginationParam) (dto.PaginationResult, error)
}

type roleServiceImpl struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleServiceImpl{db}
}

func (s *roleServiceImpl) Page(query dto.PaginationParam) (dto.PaginationResult, error) {
	var roles []model.Role
	var total int64
	db := s.db.Model(&model.Role{})
	if query.Keyword != "" {
		db = db.Where("nickname LIKE ?", "%"+query.Keyword+"%")
	}
	err := db.Count(&total).Error
	if err != nil {
		return dto.PaginationResult{}, err
	}
	err = db.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&roles).Error
	if err != nil {
		return dto.PaginationResult{}, err
	}
	return dto.PaginationResult{
		Total:    total,
		Items:    roles,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}
