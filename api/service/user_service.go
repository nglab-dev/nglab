package service

import (
	"errors"

	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/api/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return UserService{
		userRepo,
	}
}

func (s UserService) Create(user *model.User) error {
	if u, err := s.userRepo.GetByUsername(user.Username); err == nil && u != nil {
		return errors.New("username already exists")
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPwd)

	s.userRepo.Create(user)
	return nil
}

func (s UserService) Verify(username, password string) (*model.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s UserService) Get(id uint) (*model.User, error) {
	return s.userRepo.Get(id)
}
