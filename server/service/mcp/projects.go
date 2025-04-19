package mcp

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
	mcpReq "github.com/flipped-aurora/gin-vue-admin/server/model/mcp/request"
)

type ProjectsService struct{}

// CreateProjects 创建mcp服务记录
// Author [yourname](https://github.com/yourname)
func (projectsService *ProjectsService) CreateProjects(ctx context.Context, projects *mcp.Projects) (err error) {
	err = global.GVA_DB.Create(projects).Error
	return err
}

// DeleteProjects 删除mcp服务记录
// Author [yourname](https://github.com/yourname)
func (projectsService *ProjectsService) DeleteProjects(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&mcp.Projects{}, "id = ?", ID).Error
	return err
}

// DeleteProjectsByIds 批量删除mcp服务记录
// Author [yourname](https://github.com/yourname)
func (projectsService *ProjectsService) DeleteProjectsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]mcp.Projects{}, "id in ?", IDs).Error
	return err
}

// UpdateProjects 更新mcp服务记录
// Author [yourname](https://github.com/yourname)
func (projectsService *ProjectsService) UpdateProjects(ctx context.Context, projects mcp.Projects) (err error) {
	err = global.GVA_DB.Model(&mcp.Projects{}).Where("id = ?", projects.ID).Updates(&projects).Error
	return err
}

// GetProjects 根据ID获取mcp服务记录
// Author [yourname](https://github.com/yourname)
func (projectsService *ProjectsService) GetProjects(ctx context.Context, ID string) (projects mcp.Projects, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&projects).Error
	return
}

// GetProjectsInfoList 分页获取mcp服务记录
// Author [yourname](https://github.com/yourname)
func (projectsService *ProjectsService) GetProjectsInfoList(ctx context.Context, info mcpReq.ProjectsSearch) (list []mcp.Projects, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mcp.Projects{})
	var projectss []mcp.Projects
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Category != nil && *info.Category != "" {
		db = db.Where("category = ?", *info.Category)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 添加关联查询tools
	err = db.Preload("ToolsList").Find(&projectss).Error
	return projectss, total, err
}
func (projectsService *ProjectsService) GetProjectsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
