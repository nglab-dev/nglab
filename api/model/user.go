package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}
