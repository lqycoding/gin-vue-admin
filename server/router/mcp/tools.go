package mcp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ToolsRouter struct {}

// InitToolsRouter 初始化 MCP工具 路由信息
func (s *ToolsRouter) InitToolsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	toolsRouter := Router.Group("tools").Use(middleware.OperationRecord())
	toolsRouterWithoutRecord := Router.Group("tools")
	toolsRouterWithoutAuth := PublicRouter.Group("tools")
	{
		toolsRouter.POST("createTools", toolsApi.CreateTools)   // 新建MCP工具
		toolsRouter.DELETE("deleteTools", toolsApi.DeleteTools) // 删除MCP工具
		toolsRouter.DELETE("deleteToolsByIds", toolsApi.DeleteToolsByIds) // 批量删除MCP工具
		toolsRouter.PUT("updateTools", toolsApi.UpdateTools)    // 更新MCP工具
	}
	{
		toolsRouterWithoutRecord.GET("findTools", toolsApi.FindTools)        // 根据ID获取MCP工具
		toolsRouterWithoutRecord.GET("getToolsList", toolsApi.GetToolsList)  // 获取MCP工具列表
	}
	{
	    toolsRouterWithoutAuth.GET("getToolsDataSource", toolsApi.GetToolsDataSource)  // 获取MCP工具数据源
	    toolsRouterWithoutAuth.GET("getToolsPublic", toolsApi.GetToolsPublic)  // MCP工具开放接口
	}
}
