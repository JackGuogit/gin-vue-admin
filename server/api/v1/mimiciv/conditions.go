package mimiciv

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mimiciv"
	mimicivReq "github.com/flipped-aurora/gin-vue-admin/server/model/mimiciv/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InExConditionApi struct{}

// CreateInExCondition 创建纳入排除条件
// @Tags InExCondition
// @Summary 创建纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mimiciv.InExCondition true "创建纳入排除条件"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /InExCdt/createInExCondition [post]
func (InExCdtApi *InExConditionApi) CreateInExCondition(c *gin.Context) {
	var InExCdt mimiciv.InExCondition
	err := c.ShouldBindJSON(&InExCdt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = InExCdtService.CreateInExCondition(&InExCdt)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteInExCondition 删除纳入排除条件
// @Tags InExCondition
// @Summary 删除纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mimiciv.InExCondition true "删除纳入排除条件"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /InExCdt/deleteInExCondition [delete]
func (InExCdtApi *InExConditionApi) DeleteInExCondition(c *gin.Context) {
	ID := c.Query("ID")
	err := InExCdtService.DeleteInExCondition(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInExConditionByIds 批量删除纳入排除条件
// @Tags InExCondition
// @Summary 批量删除纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /InExCdt/deleteInExConditionByIds [delete]
func (InExCdtApi *InExConditionApi) DeleteInExConditionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := InExCdtService.DeleteInExConditionByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInExCondition 更新纳入排除条件
// @Tags InExCondition
// @Summary 更新纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mimiciv.InExCondition true "更新纳入排除条件"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /InExCdt/updateInExCondition [put]
func (InExCdtApi *InExConditionApi) UpdateInExCondition(c *gin.Context) {
	var InExCdt mimiciv.InExCondition
	err := c.ShouldBindJSON(&InExCdt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = InExCdtService.UpdateInExCondition(InExCdt)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInExCondition 用id查询纳入排除条件
// @Tags InExCondition
// @Summary 用id查询纳入排除条件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mimiciv.InExCondition true "用id查询纳入排除条件"
// @Success 200 {object} response.Response{data=mimiciv.InExCondition,msg=string} "查询成功"
// @Router /InExCdt/findInExCondition [get]
func (InExCdtApi *InExConditionApi) FindInExCondition(c *gin.Context) {
	ID := c.Query("ID")
	reInExCdt, err := InExCdtService.GetInExCondition(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reInExCdt, c)
}

// GetInExConditionList 分页获取纳入排除条件列表
// @Tags InExCondition
// @Summary 分页获取纳入排除条件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mimicivReq.InExConditionSearch true "分页获取纳入排除条件列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /InExCdt/getInExConditionList [get]
func (InExCdtApi *InExConditionApi) GetInExConditionList(c *gin.Context) {
	var pageInfo mimicivReq.InExConditionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := InExCdtService.GetInExConditionInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetInExConditionPublic 不需要鉴权的纳入排除条件接口
// @Tags InExCondition
// @Summary 不需要鉴权的纳入排除条件接口
// @accept application/json
// @Produce application/json
// @Param data query mimicivReq.InExConditionSearch true "分页获取纳入排除条件列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /InExCdt/getInExConditionPublic [get]
func (InExCdtApi *InExConditionApi) GetInExConditionPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	InExCdtService.GetInExConditionPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的纳入排除条件接口信息",
	}, "获取成功", c)
}
