package model

type Menu struct {
	BaseModel
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     int    `json:"type"`
	Icon     string `json:"icon"`
	Sort     int    `json:"sort"`
	Enabled  int    `json:"enabled"`
}

type Menus []Menu

func (m Menus) TableName() string {
	return "sys_menu"
}
