package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/authing"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/mcp"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Mcp     mcp.RouterGroup
	Authing authing.RouterGroup
}
