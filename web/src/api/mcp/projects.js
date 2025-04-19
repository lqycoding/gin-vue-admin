import service from '@/utils/request'
// @Tags Projects
// @Summary 创建mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Projects true "创建mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /projects/createProjects [post]
export const createProjects = (data) => {
  return service({
    url: '/projects/createProjects',
    method: 'post',
    data
  })
}

// @Tags Projects
// @Summary 删除mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Projects true "删除mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projects/deleteProjects [delete]
export const deleteProjects = (params) => {
  return service({
    url: '/projects/deleteProjects',
    method: 'delete',
    params
  })
}

// @Tags Projects
// @Summary 批量删除mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projects/deleteProjects [delete]
export const deleteProjectsByIds = (params) => {
  return service({
    url: '/projects/deleteProjectsByIds',
    method: 'delete',
    params
  })
}

// @Tags Projects
// @Summary 更新mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Projects true "更新mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /projects/updateProjects [put]
export const updateProjects = (data) => {
  return service({
    url: '/projects/updateProjects',
    method: 'put',
    data
  })
}

// @Tags Projects
// @Summary 用id查询mcp服务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Projects true "用id查询mcp服务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /projects/findProjects [get]
export const findProjects = (params) => {
  return service({
    url: '/projects/findProjects',
    method: 'get',
    params
  })
}

// @Tags Projects
// @Summary 分页获取mcp服务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取mcp服务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projects/getProjectsList [get]
export const getProjectsList = (params) => {
  return service({
    url: '/projects/getProjectsList',
    method: 'get',
    params
  })
}

// @Tags Projects
// @Summary 不需要鉴权的mcp服务接口
// @Accept application/json
// @Produce application/json
// @Param data query mcpReq.ProjectsSearch true "分页获取mcp服务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /projects/getProjectsPublic [get]
export const getProjectsPublic = () => {
  return service({
    url: '/projects/getProjectsPublic',
    method: 'get',
  })
}
