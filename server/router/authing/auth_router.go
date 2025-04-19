package authing

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

// 添加 RouterGroup 结构体以组织 Authing 相关的路由
type RouterGroup struct {
	AuthRouter
}

func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup, PublicGroup *gin.RouterGroup) {
	authApi := v1.ApiGroupApp.AuthingApiGroup.AuthApi
	{
		authRouter := PublicGroup.Group("auth")                     // 创建 /auth 路由组
		authRouter.POST("registerByEmail", authApi.RegisterByEmail) // POST /auth/registerByEmail
		authRouter.POST("loginByEmail", authApi.LoginByEmail)       // POST /auth/loginByEmail
	}
}
