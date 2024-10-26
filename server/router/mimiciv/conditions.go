package mimiciv

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InExConditionRouter struct{}

// InitInExConditionRouter 初始化 纳入排除条件 路由信息
func (s *InExConditionRouter) InitInExConditionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	InExCdtRouter := Router.Group("InExCdt").Use(middleware.OperationRecord())
	InExCdtRouterWithoutRecord := Router.Group("InExCdt")
	InExCdtRouterWithoutAuth := PublicRouter.Group("InExCdt")
	{
		InExCdtRouter.POST("createInExCondition", InExCdtApi.CreateInExCondition)             // 新建纳入排除条件
		InExCdtRouter.DELETE("deleteInExCondition", InExCdtApi.DeleteInExCondition)           // 删除纳入排除条件
		InExCdtRouter.DELETE("deleteInExConditionByIds", InExCdtApi.DeleteInExConditionByIds) // 批量删除纳入排除条件
		InExCdtRouter.PUT("updateInExCondition", InExCdtApi.UpdateInExCondition)              // 更新纳入排除条件
	}
	{
		InExCdtRouterWithoutRecord.GET("findInExCondition", InExCdtApi.FindInExCondition)       // 根据ID获取纳入排除条件
		InExCdtRouterWithoutRecord.GET("getInExConditionList", InExCdtApi.GetInExConditionList) // 获取纳入排除条件列表
	}
	{
		InExCdtRouterWithoutAuth.GET("getInExConditionPublic", InExCdtApi.GetInExConditionPublic) // 纳入排除条件开放接口
	}
}
