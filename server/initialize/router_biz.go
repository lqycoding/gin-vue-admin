package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		mcpRouter := router.RouterGroupApp.Mcp
		mcpRouter.InitProjectsRouter(privateGroup, publicGroup)
		mcpRouter.InitToolsRouter(privateGroup, publicGroup)
	}
	{
		authingRouter := router.RouterGroupApp.Authing
		authingRouter.InitAuthRouter(privateGroup, publicGroup)
	}
}
