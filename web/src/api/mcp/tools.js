import service from '@/utils/request'
// @Tags Tools
// @Summary 创建MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Tools true "创建MCP工具"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tools/createTools [post]
export const createTools = (data) => {
  return service({
    url: '/tools/createTools',
    method: 'post',
    data
  })
}

// @Tags Tools
// @Summary 删除MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Tools true "删除MCP工具"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tools/deleteTools [delete]
export const deleteTools = (params) => {
  return service({
    url: '/tools/deleteTools',
    method: 'delete',
    params
  })
}

// @Tags Tools
// @Summary 批量删除MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MCP工具"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tools/deleteTools [delete]
export const deleteToolsByIds = (params) => {
  return service({
    url: '/tools/deleteToolsByIds',
    method: 'delete',
    params
  })
}

// @Tags Tools
// @Summary 更新MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Tools true "更新MCP工具"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tools/updateTools [put]
export const updateTools = (data) => {
  return service({
    url: '/tools/updateTools',
    method: 'put',
    data
  })
}

// @Tags Tools
// @Summary 用id查询MCP工具
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Tools true "用id查询MCP工具"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tools/findTools [get]
export const findTools = (params) => {
  return service({
    url: '/tools/findTools',
    method: 'get',
    params
  })
}

// @Tags Tools
// @Summary 分页获取MCP工具列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MCP工具列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tools/getToolsList [get]
export const getToolsList = (params) => {
  return service({
    url: '/tools/getToolsList',
    method: 'get',
    params
  })
}
// @Tags Tools
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tools/findToolsDataSource [get]
export const getToolsDataSource = () => {
  return service({
    url: '/tools/getToolsDataSource',
    method: 'get',
  })
}

// @Tags Tools
// @Summary 不需要鉴权的MCP工具接口
// @Accept application/json
// @Produce application/json
// @Param data query mcpReq.ToolsSearch true "分页获取MCP工具列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tools/getToolsPublic [get]
export const getToolsPublic = () => {
  return service({
    url: '/tools/getToolsPublic',
    method: 'get',
  })
}
