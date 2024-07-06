package service

import (
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/api/repo"
)

type MenuService struct {
	menuRepo repo.MenuRepo
}

func NewMenuService(menuRepo repo.MenuRepo) MenuService {
	return MenuService{menuRepo}
}

func (s MenuService) Create(menu *model.Menu) error {
	return s.menuRepo.Create(menu)
}

func (s MenuService) List() ([]model.Menu, error) {
	menus, err := s.menuRepo.List()
	if err != nil {
		return nil, err
	}
	tree := buildMenuTree(menus, 0)
	return tree, nil
}

func buildMenuTree(menus []model.Menu, pid uint) []model.Menu {
	tree := make([]model.Menu, 0)
	for _, menu := range menus {
		if menu.ParentID == pid {
			menu.Children = buildMenuTree(menus, menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}
