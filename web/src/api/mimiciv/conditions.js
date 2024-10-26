import service from '@/utils/request'
// @Tags InExCondition
// @Summary 创建纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InExCondition true "创建纳入排除条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /InExCdt/createInExCondition [post]
export const createInExCondition = (data) => {
  return service({
    url: '/InExCdt/createInExCondition',
    method: 'post',
    data
  })
}

// @Tags InExCondition
// @Summary 删除纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InExCondition true "删除纳入排除条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /InExCdt/deleteInExCondition [delete]
export const deleteInExCondition = (params) => {
  return service({
    url: '/InExCdt/deleteInExCondition',
    method: 'delete',
    params
  })
}

// @Tags InExCondition
// @Summary 批量删除纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除纳入排除条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /InExCdt/deleteInExCondition [delete]
export const deleteInExConditionByIds = (params) => {
  return service({
    url: '/InExCdt/deleteInExConditionByIds',
    method: 'delete',
    params
  })
}

// @Tags InExCondition
// @Summary 更新纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InExCondition true "更新纳入排除条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /InExCdt/updateInExCondition [put]
export const updateInExCondition = (data) => {
  return service({
    url: '/InExCdt/updateInExCondition',
    method: 'put',
    data
  })
}

// @Tags InExCondition
// @Summary 用id查询纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InExCondition true "用id查询纳入排除条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /InExCdt/findInExCondition [get]
export const findInExCondition = (params) => {
  return service({
    url: '/InExCdt/findInExCondition',
    method: 'get',
    params
  })
}

// @Tags InExCondition
// @Summary 分页获取纳入排除条件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取纳入排除条件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /InExCdt/getInExConditionList [get]
export const getInExConditionList = (params) => {
  return service({
    url: '/InExCdt/getInExConditionList',
    method: 'get',
    params
  })
}

// @Tags InExCondition
// @Summary 不需要鉴权的纳入排除条件接口
// @accept application/json
// @Produce application/json
// @Param data query mimicivReq.InExConditionSearch true "分页获取纳入排除条件列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /InExCdt/getInExConditionPublic [get]
export const getInExConditionPublic = () => {
  return service({
    url: '/InExCdt/getInExConditionPublic',
    method: 'get',
  })
}
