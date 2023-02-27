package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cbKnowlage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BiogMainSearch struct {
	cbKnowlage.BiogMain
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
