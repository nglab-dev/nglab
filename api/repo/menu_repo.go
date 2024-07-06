package repo

import (
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/internal/database"
)

type MenuRepo struct {
	db database.Database
}

func NewMenuRepo(db database.Database) MenuRepo {
	return MenuRepo{db}
}

func (r MenuRepo) Create(menu *model.Menu) error {
	return r.db.DB.Create(menu).Error
}

func (r MenuRepo) List() (menus []model.Menu, err error) {
	if err = r.db.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
