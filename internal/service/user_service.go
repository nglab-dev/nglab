package service

import (
	"github.com/nglab-dev/nglab/internal/model"
	"gorm.io/gorm"
)

type UserService interface {
	FindByUsername(username string) (*model.User, error)
}

type userServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		db: db,
	}
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
