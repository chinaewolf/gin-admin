package cbKnowlage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage"
	cbKnowlageReq "github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BiogMainService struct {
}

// CreateBiogMain 创建BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *BiogMainService) CreateBiogMain(biogMain cbKnowlage.BiogMain) (err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Create(&biogMain).Error
	return err
}

// DeleteBiogMain 删除BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *BiogMainService) DeleteBiogMain(biogMain cbKnowlage.BiogMain) (err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Unscoped().Delete(&biogMain).Error
	return err
}

// DeleteBiogMainByIds 批量删除BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *BiogMainService) DeleteBiogMainByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Delete(&[]cbKnowlage.BiogMain{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBiogMain 更新BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *BiogMainService) UpdateBiogMain(biogMain cbKnowlage.BiogMain) (err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Save(&biogMain).Error
	return err
}

// GetBiogMain 根据id获取BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *BiogMainService) GetBiogMain(id uint) (biogMain cbKnowlage.BiogMain, err error) {
	err = global.MustGetGlobalDBByDBName("known_graph").Where("c_id = ?", id).First(&biogMain).Error
	return
}

// GetBiogMainInfoList 分页获取BiogMain记录
// Author [piexlmax](https://github.com/piexlmax)
func (biogMainService *BiogMainService) GetBiogMainInfoList(info cbKnowlageReq.BiogMainSearch) (list []cbKnowlage.BiogMain, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("known_graph").Model(&cbKnowlage.BiogMain{})
	var biogMains []cbKnowlage.BiogMain
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&biogMains).Error
	return biogMains, total, err
}
