package cbKnowlage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage"
	cbKnowlageReq "github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestStructApi struct {
}

var testStructService = service.ServiceGroupApp.CbknowlageServiceGroup.TestStructService

// CreateTestStruct 创建TestStruct
// @Tags TestStruct
// @Summary 创建TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cbKnowlage.TestStruct true "创建TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct/createTestStruct [post]
func (testStructApi *TestStructApi) CreateTestStruct(c *gin.Context) {
	var testStruct cbKnowlage.TestStruct
	err := c.ShouldBindJSON(&testStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStructService.CreateTestStruct(testStruct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStruct 删除TestStruct
// @Tags TestStruct
// @Summary 删除TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cbKnowlage.TestStruct true "删除TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStruct/deleteTestStruct [delete]
func (testStructApi *TestStructApi) DeleteTestStruct(c *gin.Context) {
	var testStruct cbKnowlage.TestStruct
	err := c.ShouldBindJSON(&testStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStructService.DeleteTestStruct(testStruct); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStructByIds 批量删除TestStruct
// @Tags TestStruct
// @Summary 批量删除TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testStruct/deleteTestStructByIds [delete]
func (testStructApi *TestStructApi) DeleteTestStructByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStructService.DeleteTestStructByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStruct 更新TestStruct
// @Tags TestStruct
// @Summary 更新TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cbKnowlage.TestStruct true "更新TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testStruct/updateTestStruct [put]
func (testStructApi *TestStructApi) UpdateTestStruct(c *gin.Context) {
	var testStruct cbKnowlage.TestStruct
	err := c.ShouldBindJSON(&testStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStructService.UpdateTestStruct(testStruct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStruct 用id查询TestStruct
// @Tags TestStruct
// @Summary 用id查询TestStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cbKnowlage.TestStruct true "用id查询TestStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testStruct/findTestStruct [get]
func (testStructApi *TestStructApi) FindTestStruct(c *gin.Context) {
	var testStruct cbKnowlage.TestStruct
	err := c.ShouldBindQuery(&testStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retestStruct, err := testStructService.GetTestStruct(testStruct.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestStruct": retestStruct}, c)
	}
}

// GetTestStructList 分页获取TestStruct列表
// @Tags TestStruct
// @Summary 分页获取TestStruct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cbKnowlageReq.TestStructSearch true "分页获取TestStruct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct/getTestStructList [get]
func (testStructApi *TestStructApi) GetTestStructList(c *gin.Context) {
	var pageInfo cbKnowlageReq.TestStructSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testStructService.GetTestStructInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
