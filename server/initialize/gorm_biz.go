package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(mcp.Projects{}, mcp.Tools{})
	if err != nil {
		return err
	}
	return nil
}
