package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nglab-dev/nglab/internal/cache"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (*model.User, error)
	Logout(user *model.UserClaims) error
	GenerateToken(user *model.User) (string, error)
	ValidateToken(token string) (*model.UserClaims, error)
}

type authServiceImpl struct {
	secretKey   string
	tokenExpiry int
	cache       *cache.Cache
	userService UserService
}

func NewAuthService(secretKey string, tokenExpiry int, cache *cache.Cache, userService UserService) AuthService {
	return &authServiceImpl{
		secretKey:   secretKey,
		tokenExpiry: tokenExpiry,
		cache:       cache,
		userService: userService,
	}
}

func (s *authServiceImpl) Login(username, password string) (*model.User, error) {
	user, err := s.userService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	if user.Enabled == constant.StatusDisabled {
		return nil, errors.New("user is disabled")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *authServiceImpl) GenerateToken(user *model.User) (string, error) {
	expiresAt := time.Now().Add(time.Duration(s.tokenExpiry) * time.Minute)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.UserClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err := claims.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	err = s.cache.Redis.Set(
		context.Background(),
		constant.CacheKeyUser+user.Username,
		token,
		time.Duration(s.tokenExpiry)*time.Minute,
	).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authServiceImpl) ValidateToken(tokenString string) (*model.UserClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	// Check if redis cache has token key
	exists, err := s.cache.Redis.Exists(
		context.Background(),
		constant.CacheKeyUser+claims.Username,
	).Result()
	if err != nil {
		return nil, err
	}
	if exists != 1 {
		return nil, errors.New("token is expired")
	}

	// Check if token is expired
	if time.Unix(claims.ExpiresAt.Unix(), 0).Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}

func (s *authServiceImpl) Logout(user *model.UserClaims) error {
	err := s.cache.Redis.Del(
		context.Background(),
		constant.CacheKeyUser+user.Username,
	).Err()
	if err != nil {
		return err
	}
	return nil
}
