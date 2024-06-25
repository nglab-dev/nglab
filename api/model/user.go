package model

type User struct {
	BaseModel
	RealName  string `json:"real_name"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Phone     string `json:"phone"`
	Type      int    `json:"type"`
	Enabled   int    `json:"enabled"`
}
