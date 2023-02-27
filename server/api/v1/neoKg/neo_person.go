package neoKg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/neoKg"
	neoKgReq "github.com/flipped-aurora/gin-vue-admin/server/model/neoKg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NeoPersonApi struct {
}

var neoPersonService = service.ServiceGroupApp.NeoKgServiceGroup.NeoPersonService

// CreateNeoPerson 创建NeoPerson
// @Tags NeoPerson
// @Summary 创建NeoPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body neoKg.NeoPerson true "创建NeoPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /neoPerson/createNeoPerson [post]
func (neoPersonApi *NeoPersonApi) CreateNeoPerson(c *gin.Context) {
	var neoPerson neoKg.NeoPerson
	err := c.ShouldBindJSON(&neoPerson)
	err1 := c.ShouldBind(neoPerson)
	fmt.Println(err, err1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neoPersonService.CreateNeoPerson(neoPerson); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNeoPerson 删除NeoPerson
// @Tags NeoPerson
// @Summary 删除NeoPerson

// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body neoKg.NeoPerson true "删除NeoPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /neoPerson/deleteNeoPerson [delete]
func (neoPersonApi *NeoPersonApi) DeleteNeoPerson(c *gin.Context) {
	var neoPerson neoKg.NeoPerson
	err := c.ShouldBindJSON(&neoPerson)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neoPersonService.DeleteNeoPerson(neoPerson); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNeoPersonByIds 批量删除NeoPerson
// @Tags NeoPerson
// @Summary 批量删除NeoPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NeoPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /neoPerson/deleteNeoPersonByIds [delete]
func (neoPersonApi *NeoPersonApi) DeleteNeoPersonByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neoPersonService.DeleteNeoPersonByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNeoPerson 更新NeoPerson
// @Tags NeoPerson
// @Summary 更新NeoPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body neoKg.NeoPerson true "更新NeoPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /neoPerson/updateNeoPerson [put]
func (neoPersonApi *NeoPersonApi) UpdateNeoPerson(c *gin.Context) {
	var neoPerson neoKg.NeoPerson
	err := c.ShouldBindJSON(&neoPerson)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := neoPersonService.UpdateNeoPerson(neoPerson); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNeoPerson 用id查询NeoPerson
// @Tags NeoPerson
// @Summary 用id查询NeoPerson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query neoKg.NeoPerson true "用id查询NeoPerson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /neoPerson/findNeoPerson [get]
func (neoPersonApi *NeoPersonApi) FindNeoPerson(c *gin.Context) {
	var neoPerson neoKg.NeoPerson
	err := c.ShouldBindQuery(&neoPerson)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reneoPerson, err := neoPersonService.GetNeoPersonById(neoPerson.PersonId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reneoPerson": reneoPerson}, c)
	}
}

// GetNeoPersonList 分页获取NeoPerson列表
// @Tags NeoPerson
// @Summary 分页获取NeoPerson列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query neoKgReq.NeoPersonSearch true "分页获取NeoPerson列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /neoPerson/getNeoPersonList [get]
func (neoPersonApi *NeoPersonApi) GetNeoPersonList(c *gin.Context) {
	var pageInfo neoKgReq.NeoPersonSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := neoPersonService.GetNeoPersonInfoList(pageInfo); err != nil {
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
