package model

import "time"

type Session struct {
	BaseModel
	UserID    uint      `json:"user_id"`
	Token     string    `json:"token"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *Session) TableName() string {
	return "sys_session"
}
