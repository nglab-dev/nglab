package model

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
