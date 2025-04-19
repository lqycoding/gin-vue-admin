package mcp

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
    mcpReq "github.com/flipped-aurora/gin-vue-admin/server/model/mcp/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ToolsApi struct {}



// CreateTools 创建MCP工具
// @Tags Tools
// @Summary 创建MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body mcp.Tools true "创建MCP工具"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tools/createTools [post]
func (toolsApi *ToolsApi) CreateTools(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var tools mcp.Tools
	err := c.ShouldBindJSON(&tools)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = toolsService.CreateTools(ctx,&tools)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTools 删除MCP工具
// @Tags Tools
// @Summary 删除MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body mcp.Tools true "删除MCP工具"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tools/deleteTools [delete]
func (toolsApi *ToolsApi) DeleteTools(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := toolsService.DeleteTools(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteToolsByIds 批量删除MCP工具
// @Tags Tools
// @Summary 批量删除MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tools/deleteToolsByIds [delete]
func (toolsApi *ToolsApi) DeleteToolsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := toolsService.DeleteToolsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTools 更新MCP工具
// @Tags Tools
// @Summary 更新MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body mcp.Tools true "更新MCP工具"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tools/updateTools [put]
func (toolsApi *ToolsApi) UpdateTools(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var tools mcp.Tools
	err := c.ShouldBindJSON(&tools)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = toolsService.UpdateTools(ctx,tools)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTools 用id查询MCP工具
// @Tags Tools
// @Summary 用id查询MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询MCP工具"
// @Success 200 {object} response.Response{data=mcp.Tools,msg=string} "查询成功"
// @Router /tools/findTools [get]
func (toolsApi *ToolsApi) FindTools(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	retools, err := toolsService.GetTools(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(retools, c)
}
// GetToolsList 分页获取MCP工具列表
// @Tags Tools
// @Summary 分页获取MCP工具列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query mcpReq.ToolsSearch true "分页获取MCP工具列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tools/getToolsList [get]
func (toolsApi *ToolsApi) GetToolsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo mcpReq.ToolsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := toolsService.GetToolsInfoList(ctx,pageInfo)
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
// GetToolsDataSource 获取Tools的数据源
// @Tags Tools
// @Summary 获取Tools的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /tools/getToolsDataSource [get]
func (toolsApi *ToolsApi) GetToolsDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := toolsService.GetToolsDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetToolsPublic 不需要鉴权的MCP工具接口
// @Tags Tools
// @Summary 不需要鉴权的MCP工具接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tools/getToolsPublic [get]
func (toolsApi *ToolsApi) GetToolsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    toolsService.GetToolsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的MCP工具接口信息",
    }, "获取成功", c)
}
