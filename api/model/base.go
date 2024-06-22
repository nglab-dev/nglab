package model

type BaseModel struct {
	ID        uint     `gorm:"primary_key" json:"id"`
	CreatedAt Datetime `json:"created_at"`
	UpdatedAt Datetime `json:"updated_at"`
}
