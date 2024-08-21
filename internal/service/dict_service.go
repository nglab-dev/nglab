package service

import (
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"gorm.io/gorm"
)

type DictService interface {
	Page(query dto.DictPaginationParam) (dto.PaginationResult, error)
}

type dictServiceImpl struct {
	db *gorm.DB
}

func NewDictService(db *gorm.DB) DictService {
	return &dictServiceImpl{db}
}

func (s *dictServiceImpl) Page(query dto.DictPaginationParam) (dto.PaginationResult, error) {
	var roles model.Dicts
	var total int64
	db := s.db.Model(&model.Dict{}).Where("type =?", query.Type)
	if query.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+query.Keyword+"%")
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
