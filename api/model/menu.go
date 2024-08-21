package model

type Menu struct {
	BaseModel
	Name     string `json:"name"`
	Type     uint   `json:"type"` // 0:目录 1:菜单 2:按钮
	Icon     string `json:"icon"`
	Path     string `json:"path"`
	ParentID uint   `json:"parent_id"`
	Sort     uint   `json:"sort"`
	Enabled  bool   `json:"enabled"`
	Children []Menu `gorm:"-" json:"children"`
}

type MenuCreateRequest struct {
	Name     string `json:"name"`
	Type     uint   `json:"type"` // 0:目录 1:菜单 2:按钮
	Icon     string `json:"icon"`
	Path     string `json:"path"`
	ParentID uint   `json:"parent_id"` // 0:顶级菜单
	Sort     uint   `json:"sort"`
	Enabled  bool   `json:"enabled"`
}

type MenuUpdateRequest struct {
	Name     string `json:"name"`
	Type     uint   `json:"type"` // 0:目录 1:菜单 2:按钮
	Icon     string `json:"icon"`
	Path     string `json:"path"`
	ParentID uint   `json:"parent_id"` // 0:顶级菜单
	Sort     uint   `json:"sort"`
	Enabled  bool   `json:"enabled"`
}
