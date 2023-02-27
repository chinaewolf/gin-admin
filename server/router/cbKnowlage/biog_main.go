package cbKnowlage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BiogMainRouter struct {
}

// InitBiogMainRouter 初始化 BiogMain 路由信息
func (s *BiogMainRouter) InitBiogMainRouter(Router *gin.RouterGroup) {
	biogMainRouter := Router.Group("biogMain").Use(middleware.OperationRecord())
	biogMainRouterWithoutRecord := Router.Group("biogMain")
	var biogMainApi = v1.ApiGroupApp.CbknowlageApiGroup.BiogMainApi
	{
		biogMainRouter.POST("createBiogMain", biogMainApi.CreateBiogMain)             // 新建BiogMain
		biogMainRouter.DELETE("deleteBiogMain", biogMainApi.DeleteBiogMain)           // 删除BiogMain
		biogMainRouter.DELETE("deleteBiogMainByIds", biogMainApi.DeleteBiogMainByIds) // 批量删除BiogMain
		biogMainRouter.PUT("updateBiogMain", biogMainApi.UpdateBiogMain)              // 更新BiogMain
	}
	{
		biogMainRouterWithoutRecord.GET("findBiogMain", biogMainApi.FindBiogMain)       // 根据ID获取BiogMain
		biogMainRouterWithoutRecord.GET("getBiogMainList", biogMainApi.GetBiogMainList) // 获取BiogMain列表
	}
}
