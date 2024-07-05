package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/internal/config"
)

type JWTConfig struct {
	secretKey  string
	expiration time.Duration
}

type AuthService struct {
	cfg       config.Config
	jwtConfig JWTConfig
}

func NewAuthService(cfg config.Config) AuthService {
	jwtConfig := JWTConfig{
		secretKey:  cfg.Auth.JWTSecret,
		expiration: time.Duration(cfg.Auth.JWTExpireTime) * time.Minute,
	}
	return AuthService{
		cfg,
		jwtConfig,
	}
}

func (a AuthService) GenerateToken(userID uint) (string, error) {
	expiresAt := time.Now().Add(a.jwtConfig.expiration)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err := claims.SignedString([]byte(a.jwtConfig.secretKey))
	if err != nil {
		return "", err
	}
	// TODO set to redis
	return token, nil
}

func (a AuthService) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.jwtConfig.secretKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
