package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nglab-dev/nglab/internal/cache"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/model/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(username, password string) (*model.User, error)
	Logout(user *dto.UserClaims) error
	GenerateToken(user *model.User, ip string, ua string) (string, error)
	ValidateToken(token string) (*dto.UserClaims, error)
}

type authServiceImpl struct {
	secretKey   string
	tokenExpiry int
	db          *gorm.DB
	cache       *cache.Cache
	userService UserService
}

func NewAuthService(
	secretKey string,
	tokenExpiry int,
	db *gorm.DB,
	cache *cache.Cache,
	userService UserService,
) AuthService {
	return &authServiceImpl{
		secretKey,
		tokenExpiry,
		db,
		cache,
		userService,
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

func (s *authServiceImpl) GenerateToken(user *model.User, ip string, ua string) (string, error) {
	expiresAt := time.Now().Add(time.Duration(s.tokenExpiry) * time.Minute)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.UserClaims{
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

	// save sys_session
	session := model.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: expiresAt,
		IP:        ip,
		UserAgent: ua,
	}
	if err := s.db.Create(&session).Error; err != nil {
		return "", err
	}

	return token, nil
}

func (s *authServiceImpl) ValidateToken(tokenString string) (*dto.UserClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*dto.UserClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	var session model.Session
	err = s.db.Where("token = ?", tokenString).First(&session).Error
	if err != nil || session.ID == 0 {
		return nil, errors.New("invalid token")
	}

	// Check if token is expired
	if time.Unix(claims.ExpiresAt.Unix(), 0).Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}

func (s *authServiceImpl) Logout(user *dto.UserClaims) error {
	err := s.db.Where("user_id = ?", user.UserID).Delete(&model.Session{}).Error
	if err != nil {
		return err
	}
	return nil
}
