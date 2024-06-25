package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nglab-dev/nglab/api/schema"
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
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, schema.JWTClaims{
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

func (a AuthService) ValidateToken(tokenString string) (*schema.JWTClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtConfig.secretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return token.Claims.(*schema.JWTClaims), nil
}
