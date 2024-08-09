package service

type RoleService interface{}

type roleServiceImpl struct{}

func NewRoleService() RoleService {
	return &roleServiceImpl{}
}
