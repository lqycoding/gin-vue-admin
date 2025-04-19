package authing

import (
	"errors"
	"fmt"

	"github.com/Authing/authing-go-sdk/lib/authentication"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	// 假设 Authing 配置在 config 包中定义
	// "github.com/flipped-aurora/gin-vue-admin/server/config"
)

type AuthService struct{}

var AuthServiceApp = new(AuthService)

// getClient 初始化 Authing 认证客户端
func (authService *AuthService) getClient() *authentication.Client {
	config := global.GVA_CONFIG.Authing

	if config.AppId == "" || config.Secret == "" || config.UserPoolId == "" {
		global.GVA_LOG.Error("Authing configuration is missing. Please check your config file.")
		return nil
	}

	client := authentication.NewClient(config.AppId, config.Secret)
	client.UserPoolId = config.UserPoolId

	// 如果配置了自定义域名（私有化部署），则设置 Host
	if config.Host != "" {
		client.Host = config.Host
	}

	return client
}

// RegisterByEmailInput 定义邮箱注册的输入结构体
type RegisterByEmailInput struct {
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Username    string    `json:"username"`
	NickName    string    `json:"nickName"`
	HeaderImg   string    `json:"headerImg"`
	Phone       string    `json:"phone"`
	AuthorityId uint      `json:"authorityId"`
	Enable      int       `json:"enable"`                           // 用户是否被冻结 1正常 2冻结
	Gender      string    `json:"gender"`                           // 性别，可选值为 M（男）、F（女）、U（未知）
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"` // 用户UUID
}

// RegisterByEmail 使用 Authing SDK 通过邮箱注册用户
func (authService *AuthService) RegisterByEmail(input RegisterByEmailInput) (*system.SysUser, error) {
	client := authService.getClient()
	if client == nil {
		return nil, errors.New("authing client initialization failed, check configuration")
	}

	// 设置默认的 AuthorityId
	authorityId := 666
	var authorities []system.SysAuthority
	authorities = append(authorities, system.SysAuthority{
		AuthorityId: uint(authorityId),
	})

	// 设置默认性别为未知
	if input.Gender == "" {
		input.Gender = "U"
	}
	if input.Enable == 0 {
		input.Enable = 1 // 默认启用用户
	}
	if input.Password == "" {
		input.Password = "123456" // 设置默认密码
	}
	input.UUID = uuid.New()
	// 1. 检查邮箱是否已注册
	exists, err := client.IsUserExists(&model.IsUserExistsRequest{
		Email: &input.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("check email existence failed: %w", err)
	}
	if exists != nil && *exists {
		return nil, errors.New("邮箱已被注册")
	}

	// 2. 在 Authing 中注册用户
	profile := &model.RegisterProfile{
		Gender: &input.Gender,
	}
	if input.Username != "" {
		profile.Username = &input.Username
	}
	if input.NickName != "" {
		profile.Nickname = &input.NickName
	}

	req := &model.RegisterByEmailInput{
		Email:    input.Email,
		Password: input.Password,
		Profile:  profile,
	}

	authingUser, err := client.RegisterByEmail(req)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("Authing RegisterByEmail failed: %v", err))
		return nil, fmt.Errorf("authing registration failed: %w", err)
	}

	// 3. 创建系统用户
	sysUser := &system.SysUser{
		UUID:        input.UUID,
		Username:    input.Email,
		NickName:    input.Email,
		Password:    utils.BcryptHash(input.Password), // 本地密码也进行加密存储
		HeaderImg:   input.HeaderImg,
		AuthorityId: input.AuthorityId, // 使用系统角色ID
		Authorities: authorities,
		Enable:      input.Enable,
		Phone:       input.Phone,
		Email:       input.Email,
		AuthingId:   authingUser.Id, // 存储 Authing 用户 ID
	}

	// 4. 在本地数据库创建用户
	err = global.GVA_DB.Create(sysUser).Error
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("Local user creation failed: %v", err))
		// TODO: 考虑在本地创建失败时是否需要删除 Authing 用户
		return nil, fmt.Errorf("local user creation failed: %w", err)
	}

	return sysUser, nil
}

// LoginByEmailInput 定义邮箱登录的输入结构体
type LoginByEmailInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginByEmail 使用 Authing SDK 通过邮箱登录用户
// 返回 Authing 的 UserInfo 类型，其中包含 token 等信息
func (authService *AuthService) LoginByEmail(input LoginByEmailInput) (*model.User, error) {
	client := authService.getClient()
	if client == nil {
		return nil, errors.New("authing client initialization failed, check configuration")
	}

	loginInput := model.LoginByEmailInput{
		Email:    input.Email,
		Password: input.Password,
	}

	resp, err := client.LoginByEmail(loginInput)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("Authing LoginByEmail failed: %v", err))
		// 可以尝试解析 Authing 返回的更具体的错误信息
		return nil, fmt.Errorf("authing login failed: %w", err)
	}

	// 登录成功，resp 中包含 token 等信息
	return resp, nil
}
