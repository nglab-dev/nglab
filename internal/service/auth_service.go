package service

import "github.com/nglab-dev/nglab/internal/repo"

type AuthService struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return AuthService{
		userRepo: userRepo,
	}
}
