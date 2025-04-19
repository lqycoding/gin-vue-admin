package config

type Authing struct {
	AppId      string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`                   // Authing 应用 ID
	Secret     string `mapstructure:"secret" json:"secret" yaml:"secret"`                   // Authing 应用密钥
	UserPoolId string `mapstructure:"user-pool-id" json:"user-pool-id" yaml:"user-pool-id"` // Authing 用户池 ID
	Host       string `mapstructure:"host" json:"host" yaml:"host"`                         // Authing 服务地址（可选，用于私有化部署）
}
