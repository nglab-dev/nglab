package models

type BaseModel struct {
	ID int64 `gorm:"primary_key" json:"id"`
}
