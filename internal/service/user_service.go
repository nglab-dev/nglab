package service

import (
	"errors"

	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	FindByID(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Page(query dto.PaginationParam) (dto.PaginationResult, error)
}

type userServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		db,
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

func (s *userServiceImpl) Create(user *model.User) error {
	u, err := s.FindByUsername(user.Username)
	if err != nil {
		return err
	}
	if u != nil {
		return errors.New("username already exists")
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPwd)
	return s.db.Create(user).Error
}

func (s *userServiceImpl) Update(user *model.User) error {
	return s.db.Save(user).Error
}

func (s *userServiceImpl) Page(query dto.PaginationParam) (dto.PaginationResult, error) {
	var users []model.User
	var total int64
	db := s.db.Model(&model.User{}).Preload("Roles")
	if query.Keyword != "" {
		db = db.Where("nickname LIKE ?", "%"+query.Keyword+"%")
	}
	err := db.Count(&total).Error
	if err != nil {
		return dto.PaginationResult{}, err
	}
	err = db.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&users).Error
	if err != nil {
		return dto.PaginationResult{}, err
	}
	return dto.PaginationResult{
		Total:    total,
		Items:    users,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}
