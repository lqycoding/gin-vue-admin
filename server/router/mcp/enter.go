package mcp

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ProjectsRouter
	ToolsRouter
}

var (
	projectsApi = api.ApiGroupApp.McpApiGroup.ProjectsApi
	toolsApi    = api.ApiGroupApp.McpApiGroup.ToolsApi
)
