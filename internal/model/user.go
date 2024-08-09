package model

import "github.com/nglab-dev/nglab/internal/model/dto"

type User struct {
	BaseModel
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Enabled   int    `json:"enabled"`
	Roles     Roles  `gorm:"-" json:"roles"`
}

type Users []*User

type UserInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Realname string `json:"realname"`
	Roles    Roles  `json:"roles"`
}

type UserQueryParam struct {
	dto.PaginationParam
}

func (u *User) TableName() string {
	return "sys_user"
}
