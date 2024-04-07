package pagesServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	pagesResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/pages"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
)

var PageService = &pageService{}

type pageService struct {
}

func (p *pageService) CreatePage(params *dto.CreatePageParams) error {
	return repositories.PageRepositorysModules.CreatePage(params)
}

func (p *pageService) DeletePage(params string) error {
	return repositories.PageRepositorysModules.DeletePage(params)
}

func (p *pageService) GetPages() ([]*dto.SinglePageResponse, error) {
	pages, err := repositories.PageRepositorysModules.GetPages()
	if err != nil {
		return nil, err
	}
	buildPages := utils.BuildPages(pages)
	return buildPages, nil
}

func (p *pageService) GetUserMenus(id string) ([]*pagesResponsiesModules.SinglePageResponse, error) {
	// 查找用户角色ID
	rolesID := repositories.UsersAndRolesRepositorys.FindUserRolesID(id)
	// 查找用户页面权限
	var pagesID []string
	for _, roleID := range rolesID {
		page, err := repositories.RolesAndPagesRepository.GetRecordsByRoleID(roleID)
		if err != nil {
			return nil, err
		}
		pagesID = append(pagesID, page...)
	}
	// pagesID 去重
	pagesID = utils.Unique(pagesID)
	// 获取页面详情
	var pagesDetail []*dto.SinglePageResponse
	for _, pageID := range pagesID {
		pageDetail, err := repositories.PageRepositorysModules.FindPageByID(pageID)
		if err != nil {
			return nil, err
		}
		pagesDetail = append(pagesDetail, pageDetail)
	}
	menu := utils.BuildPages(pagesDetail)
	return menu, nil
}
