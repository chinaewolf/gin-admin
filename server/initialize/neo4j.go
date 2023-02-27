package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	//"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.uber.org/zap"
)

func Neo4j() {
	neo4jCfg := global.GVA_CONFIG.Neo4j
	driver, err := neo4j.NewDriver(neo4jCfg.Addr,
		neo4j.BasicAuth(neo4jCfg.User, neo4jCfg.Password, ""))
	if err != nil {
		global.GVA_LOG.Error("neo4j create driver  failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("neo4j create driver Ok:")
		global.GVA_NEO4J = driver
	}

}
