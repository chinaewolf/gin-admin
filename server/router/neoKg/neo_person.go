package neoKg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NeoKgRouter struct {
}

// 初始化 NeoKg 路由信息
func (s *NeoKgRouter) InitNeoKgRouter(Router *gin.RouterGroup) {
	neoPersonRouter := Router.Group("neo").Use(middleware.OperationRecord())
	neoPersonRouterWithoutRecord := Router.Group("neo")
	var neoPersonApi = v1.ApiGroupApp.NeoKgApiGroup.NeoPersonApi
	{
		neoPersonRouter.POST("createNeoPerson", neoPersonApi.CreateNeoPerson)             // 新建NeoPerson
		neoPersonRouter.DELETE("deleteNeoPerson", neoPersonApi.DeleteNeoPerson)           // 删除NeoPerson
		neoPersonRouter.DELETE("deleteNeoPersonByIds", neoPersonApi.DeleteNeoPersonByIds) // 批量删除NeoPerson
		neoPersonRouter.PUT("updateNeoPerson", neoPersonApi.UpdateNeoPerson)              // 更新NeoPerson
	}
	{
		neoPersonRouterWithoutRecord.GET("findNeoPerson", neoPersonApi.FindNeoPerson)       // 根据ID获取NeoPerson
		neoPersonRouterWithoutRecord.GET("getNeoPersonList", neoPersonApi.GetNeoPersonList) // 获取NeoPerson列表
	}
}
