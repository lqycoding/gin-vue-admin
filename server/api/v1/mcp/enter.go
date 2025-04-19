package mcp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ProjectsApi
	ToolsApi
}

var (
	projectsService = service.ServiceGroupApp.McpServiceGroup.ProjectsService
	toolsService    = service.ServiceGroupApp.McpServiceGroup.ToolsService
)
