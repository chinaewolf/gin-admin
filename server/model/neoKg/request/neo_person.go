package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/neoKg"
)

type NeoPersonSearch struct {
	neoKg.NeoPerson

	request.PageInfo
}
