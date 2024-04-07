package usersAndRolesServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
)

var UserAndRolesService = &userAndRolesService{}

type userAndRolesService struct {
}

func (u *userAndRolesService) FindRoleById(id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse == nil {
		return errors.New("角色不存在")
	}
	return nil
}

func (u *userAndRolesService) GetRoles(params *dto.QueryRolesParams) ([]*dto.SingleRoleResponse, error) {
	return repositories.RoleRepositorysModules.GetRoles(params)
}

func (u *userAndRolesService) UpdateRole(role *dto.UpdateRoleParams, id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse == nil {
		return errors.New("角色不存在")
	}
	return repositories.RoleRepositorysModules.UpdateRole(role, id)
}
