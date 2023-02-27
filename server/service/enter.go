package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/cbKnowlage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/neoKg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	CbknowlageServiceGroup cbKnowlage.ServiceGroup
	NeoKgServiceGroup      neoKg.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
