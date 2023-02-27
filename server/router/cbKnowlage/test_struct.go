package cbKnowlage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStructRouter struct {
}

// InitTestStructRouter 初始化 TestStruct 路由信息
func (s *TestStructRouter) InitTestStructRouter(Router *gin.RouterGroup) {
	testStructRouter := Router.Group("testStruct").Use(middleware.OperationRecord())
	testStructRouterWithoutRecord := Router.Group("testStruct")
	var testStructApi = v1.ApiGroupApp.CbknowlageApiGroup.TestStructApi
	{
		testStructRouter.POST("createTestStruct", testStructApi.CreateTestStruct)             // 新建TestStruct
		testStructRouter.DELETE("deleteTestStruct", testStructApi.DeleteTestStruct)           // 删除TestStruct
		testStructRouter.DELETE("deleteTestStructByIds", testStructApi.DeleteTestStructByIds) // 批量删除TestStruct
		testStructRouter.PUT("updateTestStruct", testStructApi.UpdateTestStruct)              // 更新TestStruct
	}
	{
		testStructRouterWithoutRecord.GET("findTestStruct", testStructApi.FindTestStruct)       // 根据ID获取TestStruct
		testStructRouterWithoutRecord.GET("getTestStructList", testStructApi.GetTestStructList) // 获取TestStruct列表
	}
}
