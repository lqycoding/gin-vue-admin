package mcp

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
    mcpReq "github.com/flipped-aurora/gin-vue-admin/server/model/mcp/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ProjectsApi struct {}



// CreateProjects 创建mcp服务
// @Tags Projects
// @Summary 创建mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body mcp.Projects true "创建mcp服务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /projects/createProjects [post]
func (projectsApi *ProjectsApi) CreateProjects(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var projects mcp.Projects
	err := c.ShouldBindJSON(&projects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectsService.CreateProjects(ctx,&projects)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteProjects 删除mcp服务
// @Tags Projects
// @Summary 删除mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body mcp.Projects true "删除mcp服务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /projects/deleteProjects [delete]
func (projectsApi *ProjectsApi) DeleteProjects(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := projectsService.DeleteProjects(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProjectsByIds 批量删除mcp服务
// @Tags Projects
// @Summary 批量删除mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /projects/deleteProjectsByIds [delete]
func (projectsApi *ProjectsApi) DeleteProjectsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := projectsService.DeleteProjectsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProjects 更新mcp服务
// @Tags Projects
// @Summary 更新mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body mcp.Projects true "更新mcp服务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /projects/updateProjects [put]
func (projectsApi *ProjectsApi) UpdateProjects(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var projects mcp.Projects
	err := c.ShouldBindJSON(&projects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectsService.UpdateProjects(ctx,projects)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProjects 用id查询mcp服务
// @Tags Projects
// @Summary 用id查询mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询mcp服务"
// @Success 200 {object} response.Response{data=mcp.Projects,msg=string} "查询成功"
// @Router /projects/findProjects [get]
func (projectsApi *ProjectsApi) FindProjects(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reprojects, err := projectsService.GetProjects(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reprojects, c)
}
// GetProjectsList 分页获取mcp服务列表
// @Tags Projects
// @Summary 分页获取mcp服务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query mcpReq.ProjectsSearch true "分页获取mcp服务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /projects/getProjectsList [get]
func (projectsApi *ProjectsApi) GetProjectsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo mcpReq.ProjectsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := projectsService.GetProjectsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetProjectsPublic 不需要鉴权的mcp服务接口
// @Tags Projects
// @Summary 不需要鉴权的mcp服务接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /projects/getProjectsPublic [get]
func (projectsApi *ProjectsApi) GetProjectsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    projectsService.GetProjectsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的mcp服务接口信息",
    }, "获取成功", c)
}
