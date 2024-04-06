package buildTree

import (
	"github.com/Xi-Yuer/cms/dto"
	"sort"
)

// BuildMenu 构建菜单树
func BuildMenu(pages []*dto.SinglePageResponse) []*dto.SinglePageResponse {
	// 构建完整的树形结构
	pageMap := make(map[string]*dto.SinglePageResponse)
	for _, page := range pages {
		pageMap[page.PageID] = page
		if page.ParentPage != nil {
			parent := pageMap[*page.ParentPage]
			if parent != nil {
				parent.Children = append(parent.Children, page)
			}
		}
	}
	// 找到所有根节点
	var roots []*dto.SinglePageResponse
	for _, page := range pages {
		if page.ParentPage == nil {
			roots = append(roots, page)
		}
	}
	// 按照PageOrder排序子页面
	for _, root := range roots {
		sortSubmenu(root)
	}
	// 按照PageOrder排序根节点
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].PageOrder < roots[j].PageOrder
	})
	return roots
}

// sortSubmenu 对子菜单按照PageOrder排序
func sortSubmenu(page *dto.SinglePageResponse) {
	sort.Slice(page.Children, func(i, j int) bool {
		return page.Children[i].PageOrder < page.Children[j].PageOrder
	})
	for _, child := range page.Children {
		sortSubmenu(child)
	}
}

// BuildDepartment 构建部门树
func BuildDepartment(departments []*dto.DepartmentResponse) []*dto.DepartmentResponse {
	// 构建完整的树形结构
	pageMap := make(map[string]*dto.DepartmentResponse)
	for _, page := range departments {
		pageMap[page.ID] = page
		if page.ParentDepartment != nil {
			parent := pageMap[*page.ParentDepartment]
			if parent != nil {
				parent.Children = append(parent.Children, *page)
			}
		}
	}
	// 找到所有根节点
	var roots []*dto.DepartmentResponse
	for _, page := range departments {
		if page.ParentDepartment == nil {
			roots = append(roots, page)
		}
	}
	// 按照PageOrder排序子页面
	for _, root := range roots {
		sortSubDepartment(root)
	}
	// 按照PageOrder排序根节点
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].DepartmentOrder < roots[j].DepartmentOrder
	})
	return roots
}

// sortSubDepartment 对子菜单按照PageOrder排序
func sortSubDepartment(page *dto.DepartmentResponse) {
	sort.Slice(page.Children, func(i, j int) bool {
		return page.Children[i].DepartmentOrder < page.Children[j].DepartmentOrder
	})
	for _, child := range page.Children {
		sortSubDepartment(&child)
	}
}
