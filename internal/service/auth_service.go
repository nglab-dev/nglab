package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nglab-dev/nglab/internal/model"
)

type AuthService interface {
	Login(user *model.User) (string, error)
	ValidateToken(token string) (*model.UserClaims, error)
}

type authServiceImpl struct {
	secretKey   string
	tokenExpiry int
}

func NewAuthService(secretKey string, tokenExpiry int) AuthService {
	return &authServiceImpl{
		secretKey:   secretKey,
		tokenExpiry: tokenExpiry,
	}
}

func (s *authServiceImpl) Login(user *model.User) (string, error) {
	expiresAt := time.Now().Add(time.Duration(s.tokenExpiry) * time.Minute)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.UserClaims{
		UserID: user.ID,
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

	// TODO: save to redis
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
	return claims, nil
}
