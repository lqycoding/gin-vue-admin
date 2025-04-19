import service from '@/utils/request'
// @Tags McpServer
// @Summary 创建mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "创建mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mcpServer/createMcpServer [post]
export const createMcpServer = (data) => {
  return service({
    url: '/mcpServer/createMcpServer',
    method: 'post',
    data
  })
}

// @Tags McpServer
// @Summary 删除mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "删除mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcpServer/deleteMcpServer [delete]
export const deleteMcpServer = (params) => {
  return service({
    url: '/mcpServer/deleteMcpServer',
    method: 'delete',
    params
  })
}

// @Tags McpServer
// @Summary 批量删除mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcpServer/deleteMcpServer [delete]
export const deleteMcpServerByIds = (params) => {
  return service({
    url: '/mcpServer/deleteMcpServerByIds',
    method: 'delete',
    params
  })
}

// @Tags McpServer
// @Summary 更新mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.McpServer true "更新mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mcpServer/updateMcpServer [put]
export const updateMcpServer = (data) => {
  return service({
    url: '/mcpServer/updateMcpServer',
    method: 'put',
    data
  })
}

// @Tags McpServer
// @Summary 用id查询mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.McpServer true "用id查询mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mcpServer/findMcpServer [get]
export const findMcpServer = (params) => {
  return service({
    url: '/mcpServer/findMcpServer',
    method: 'get',
    params
  })
}

// @Tags McpServer
// @Summary 分页获取mcp服务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mcp服务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcpServer/getMcpServerList [get]
export const getMcpServerList = (params) => {
  return service({
    url: '/mcpServer/getMcpServerList',
    method: 'get',
    params
  })
}

// @Tags McpServer
// @Summary 不需要鉴权的mcp服务接口
// @Accept application/json
// @Produce application/json
// @Param data query mcpReq.McpServerSearch true "分页获取mcp服务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mcpServer/getMcpServerPublic [get]
export const getMcpServerPublic = () => {
  return service({
    url: '/mcpServer/getMcpServerPublic',
    method: 'get',
  })
}
