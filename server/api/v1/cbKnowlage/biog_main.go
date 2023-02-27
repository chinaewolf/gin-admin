package cbKnowlage

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage"
	cbKnowlageReq "github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BiogMainApi struct {
}

var biogMainService = service.ServiceGroupApp.CbknowlageServiceGroup.BiogMainService

// CreateBiogMain 创建BiogMain
// @Tags BiogMain
// @Summary 创建BiogMain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cbKnowlage.BiogMain true "创建BiogMain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /biogMain/createBiogMain [post]
func (biogMainApi *BiogMainApi) CreateBiogMain(c *gin.Context) {
	var biogMain cbKnowlage.BiogMain
	err := c.ShouldBindJSON(&biogMain)
	err1 := c.ShouldBind(biogMain)
	fmt.Println(err, err1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := biogMainService.CreateBiogMain(biogMain); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBiogMain 删除BiogMain
// @Tags BiogMain
// @Summary 删除BiogMain

// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cbKnowlage.BiogMain true "删除BiogMain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /biogMain/deleteBiogMain [delete]
func (biogMainApi *BiogMainApi) DeleteBiogMain(c *gin.Context) {
	var biogMain cbKnowlage.BiogMain
	err := c.ShouldBindJSON(&biogMain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := biogMainService.DeleteBiogMain(biogMain); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBiogMainByIds 批量删除BiogMain
// @Tags BiogMain
// @Summary 批量删除BiogMain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BiogMain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /biogMain/deleteBiogMainByIds [delete]
func (biogMainApi *BiogMainApi) DeleteBiogMainByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := biogMainService.DeleteBiogMainByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBiogMain 更新BiogMain
// @Tags BiogMain
// @Summary 更新BiogMain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cbKnowlage.BiogMain true "更新BiogMain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /biogMain/updateBiogMain [put]
func (biogMainApi *BiogMainApi) UpdateBiogMain(c *gin.Context) {
	var biogMain cbKnowlage.BiogMain
	err := c.ShouldBindJSON(&biogMain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := biogMainService.UpdateBiogMain(biogMain); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBiogMain 用id查询BiogMain
// @Tags BiogMain
// @Summary 用id查询BiogMain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cbKnowlage.BiogMain true "用id查询BiogMain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /biogMain/findBiogMain [get]
func (biogMainApi *BiogMainApi) FindBiogMain(c *gin.Context) {
	var biogMain cbKnowlage.BiogMain
	err := c.ShouldBindQuery(&biogMain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebiogMain, err := biogMainService.GetBiogMain(biogMain.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebiogMain": rebiogMain}, c)
	}
}

// GetBiogMainList 分页获取BiogMain列表
// @Tags BiogMain
// @Summary 分页获取BiogMain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cbKnowlageReq.BiogMainSearch true "分页获取BiogMain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /biogMain/getBiogMainList [get]
func (biogMainApi *BiogMainApi) GetBiogMainList(c *gin.Context) {
	var pageInfo cbKnowlageReq.BiogMainSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := biogMainService.GetBiogMainInfoList(pageInfo); err != nil {
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
