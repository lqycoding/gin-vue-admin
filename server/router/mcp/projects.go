package mcp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectsRouter struct {}

// InitProjectsRouter 初始化 mcp服务 路由信息
func (s *ProjectsRouter) InitProjectsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	projectsRouter := Router.Group("projects").Use(middleware.OperationRecord())
	projectsRouterWithoutRecord := Router.Group("projects")
	projectsRouterWithoutAuth := PublicRouter.Group("projects")
	{
		projectsRouter.POST("createProjects", projectsApi.CreateProjects)   // 新建mcp服务
		projectsRouter.DELETE("deleteProjects", projectsApi.DeleteProjects) // 删除mcp服务
		projectsRouter.DELETE("deleteProjectsByIds", projectsApi.DeleteProjectsByIds) // 批量删除mcp服务
		projectsRouter.PUT("updateProjects", projectsApi.UpdateProjects)    // 更新mcp服务
	}
	{
		projectsRouterWithoutRecord.GET("findProjects", projectsApi.FindProjects)        // 根据ID获取mcp服务
		projectsRouterWithoutRecord.GET("getProjectsList", projectsApi.GetProjectsList)  // 获取mcp服务列表
	}
	{
	    projectsRouterWithoutAuth.GET("getProjectsPublic", projectsApi.GetProjectsPublic)  // mcp服务开放接口
	}
}
