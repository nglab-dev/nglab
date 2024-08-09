package dto

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID               uint   `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims `json:"claims"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginUser struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Token       string `json:"token"`
	IP          string `json:"ip"`
	UserAgent   string `json:"user_agent"`
	TokenExpiry int64  `json:"token_expiry"`
}

func (u *LoginUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *LoginUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
