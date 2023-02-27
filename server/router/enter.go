package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/cbKnowlage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/neoKg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Cbknowlage cbKnowlage.RouterGroup
	NeoKg      neoKg.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
