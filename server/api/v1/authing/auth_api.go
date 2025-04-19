package authing

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/authing"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthApi struct{}

// 添加 ApiGroup 结构体以组织 Authing 相关的 API
type ApiGroup struct {
	AuthApi
}

// RegisterByEmail
// @Tags Authing
// @Summary 使用邮箱注册 (Authing)
// @Description 使用 Authing 服务通过邮箱和密码注册新用户，同时在本地系统创建对应用户
// @Accept json
// @Produce application/json
// @Param data body authing.RegisterByEmailInput true "用户注册信息"
// @Success 200 {object} response.Response{data=system.SysUser,msg=string} "注册成功"
// @Router /auth/registerByEmail [post]
func (a *AuthApi) RegisterByEmail(c *gin.Context) {
	var input authing.RegisterByEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithMessage("请求参数错误: "+err.Error(), c)
		return
	}

	// // 基本参数验证
	// if input.Username == "" {
	// 	response.FailWithMessage("用户名不能为空", c)
	// 	return
	// }
	// if len(input.Password) < 6 {
	// 	response.FailWithMessage("密码长度不能小于6位", c)
	// 	return
	// }

	// // 验证邮箱格式
	// if err := utils.Verify(input.Email, utils.EmailVerify); err != nil {
	// 	response.FailWithMessage("邮箱格式错误", c)
	// 	return
	// }

	// 检查用户名是否已存在（本地系统）
	if !errors.Is(global.GVA_DB.Where("username = ?", input.Username).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户名已存在", c)
		return
	}

	// 检查邮箱是否已存在（本地系统）
	if !errors.Is(global.GVA_DB.Where("email = ?", input.Email).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage("邮箱已被注册", c)
		return
	}

	// 调用服务层进行注册
	user, err := authing.AuthServiceApp.RegisterByEmail(input)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(user, "注册成功", c)
}

// LoginByEmail
// @Tags Authing
// @Summary 使用邮箱登录 (Authing)
// @Description 使用 Authing 服务通过邮箱和密码登录，返回用户信息和 token
// @Accept json
// @Produce application/json
// @Param data body authing.LoginByEmailInput true "登录信息"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "登录成功"
// @Router /auth/loginByEmail [post]
func (a *AuthApi) LoginByEmail(c *gin.Context) {
	var input authing.LoginByEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithMessage("请求参数错误: "+err.Error(), c)
		return
	}

	// 验证邮箱格式
	// if err := utils.Verify(input.Email, utils.EmailVerify); err != nil {
	// 	response.FailWithMessage("邮箱格式错误", c)
	// 	return
	// }

	// 调用 Authing 服务进行登录
	authingUser, err := authing.AuthServiceApp.LoginByEmail(input)
	if err != nil {
		global.GVA_LOG.Error("Authing 登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	// 查找对应的系统用户
	var sysUser system.SysUser
	if err := global.GVA_DB.Where("authing_id = ?", authingUser.Id).First(&sysUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("用户未在系统中注册", c)
		} else {
			global.GVA_LOG.Error("查找本地用户失败!", zap.Error(err))
			response.FailWithMessage("系统错误", c)
		}
		return
	}

	// 检查用户状态
	if sysUser.Enable != 1 {
		response.FailWithMessage("用户已被禁用", c)
		return
	}

	// 返回成功响应，包含用户信息和 Authing token
	response.OkWithDetailed(gin.H{
		"user": gin.H{
			"id":          sysUser.ID,
			"uuid":        sysUser.UUID,
			"username":    sysUser.Username,
			"nickname":    sysUser.NickName,
			"headerImg":   sysUser.HeaderImg,
			"authorityId": sysUser.AuthorityId,
			"email":       sysUser.Email,
			"phone":       sysUser.Phone,
			"enable":      sysUser.Enable,
		},
		"token": authingUser.Token,
	}, "登录成功", c)
}
