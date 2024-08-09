package model

type User struct {
	BaseModel
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Enabled   int    `json:"enabled"`
}

func (u *User) TableName() string {
	return "sys_user"
}
