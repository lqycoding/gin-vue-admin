
// 自动生成模板Tools
package mcp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// MCP工具 结构体  Tools
type Tools struct {
    global.GVA_MODEL
    ServerId  *int `json:"serverId" form:"serverId" gorm:"column:server_id;comment:;"`  //服务id
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:;"`  //工具名称
    Description  *string `json:"description" form:"description" gorm:"column:description;comment:;"`  //工具描述
    Points  *int `json:"points" form:"points" gorm:"column:points;comment:;"`  //积分消耗
    Input_schema  datatypes.JSON `json:"input_schema" form:"input_schema" gorm:"column:input_schema;comment:;type:text;" swaggertype:"object"`  //输入参数模式(JSON)
}


// TableName MCP工具 Tools自定义表名 tools
func (Tools) TableName() string {
    return "tools"
}





