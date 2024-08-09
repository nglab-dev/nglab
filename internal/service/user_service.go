package service

import (
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"gorm.io/gorm"
)

type UserService interface {
	FindByID(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Page(query dto.PageRequest) (dto.PageResponse, error)
}

type userServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		db: db,
	}
}

func (s *userServiceImpl) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := s.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, err
}

func (s *userServiceImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := s.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, err
}

func (s *userServiceImpl) Page(query dto.PageRequest) (dto.PageResponse, error) {
	var users []model.User
	var total int64
	db := s.db.Model(&model.User{})
	if query.Keyword != "" {
		db = db.Where("nickname LIKE ?", "%"+query.Keyword+"%")
	}
	err := db.Count(&total).Error
	if err != nil {
		return dto.PageResponse{}, err
	}
	err = db.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&users).Error
	if err != nil {
		return dto.PageResponse{}, err
	}
	return dto.PageResponse{
		Total: total,
		Data:  users,
	}, nil
}
