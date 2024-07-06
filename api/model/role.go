package model

type Role struct {
	BaseModel
	Name    string `json:"name"`
	Remark  string `json:"remark"`
	Enabled bool   `json:"enabled"`
}
