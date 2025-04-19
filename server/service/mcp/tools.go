
package mcp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
    mcpReq "github.com/flipped-aurora/gin-vue-admin/server/model/mcp/request"
)

type ToolsService struct {}
// CreateTools 创建MCP工具记录
// Author [yourname](https://github.com/yourname)
func (toolsService *ToolsService) CreateTools(ctx context.Context, tools *mcp.Tools) (err error) {
	err = global.GVA_DB.Create(tools).Error
	return err
}

// DeleteTools 删除MCP工具记录
// Author [yourname](https://github.com/yourname)
func (toolsService *ToolsService)DeleteTools(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&mcp.Tools{},"id = ?",ID).Error
	return err
}

// DeleteToolsByIds 批量删除MCP工具记录
// Author [yourname](https://github.com/yourname)
func (toolsService *ToolsService)DeleteToolsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]mcp.Tools{},"id in ?",IDs).Error
	return err
}

// UpdateTools 更新MCP工具记录
// Author [yourname](https://github.com/yourname)
func (toolsService *ToolsService)UpdateTools(ctx context.Context, tools mcp.Tools) (err error) {
	err = global.GVA_DB.Model(&mcp.Tools{}).Where("id = ?",tools.ID).Updates(&tools).Error
	return err
}

// GetTools 根据ID获取MCP工具记录
// Author [yourname](https://github.com/yourname)
func (toolsService *ToolsService)GetTools(ctx context.Context, ID string) (tools mcp.Tools, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tools).Error
	return
}
// GetToolsInfoList 分页获取MCP工具记录
// Author [yourname](https://github.com/yourname)
func (toolsService *ToolsService)GetToolsInfoList(ctx context.Context, info mcpReq.ToolsSearch) (list []mcp.Tools, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mcp.Tools{})
    var toolss []mcp.Tools
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&toolss).Error
	return  toolss, total, err
}
func (toolsService *ToolsService)GetToolsDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   serverId := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("projects").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&serverId)
	   res["serverId"] = serverId
	return
}
func (toolsService *ToolsService)GetToolsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
