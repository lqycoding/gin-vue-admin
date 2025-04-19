// 自动生成模板Projects
package mcp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// mcp服务 结构体  Projects
type Projects struct {
	global.GVA_MODEL
	Package           *string        `json:"package" form:"package" gorm:"column:package;comment:;"`                                                //包名
	Name              *string        `json:"name" form:"name" gorm:"column:name;comment:;"`                                                         //服务名
	Description       *string        `json:"description" form:"description" gorm:"column:description;comment:;"`                                    //服务器描述
	Category          *string        `json:"category" form:"category" gorm:"column:category;comment:;"`                                             //分类
	Tags              datatypes.JSON `json:"tags" form:"tags" gorm:"column:tags;comment:;type:text;" swaggertype:"array,object"`                    //标签(JSON数组)
	Status            *string        `json:"status" form:"status" gorm:"column:status;comment:;"`                                                   //状态
	CallMethod        datatypes.JSON `json:"callMethod" form:"callMethod" gorm:"column:call_method;comment:;type:text;" swaggertype:"array,object"` //调用方法(JSON数组)
	Owner             *string        `json:"owner" form:"owner" gorm:"column:owner;comment:;"`                                                      //所有者
	OnlineUsageCount  *int           `json:"onlineUsageCount" form:"onlineUsageCount" gorm:"column:online_usage_count;comment:;"`                   //在线使用次数
	OfflineUsageCount *int           `json:"offlineUsageCount" form:"offlineUsageCount" gorm:"column:offline_usage_count;comment:;"`                //离线使用次数
	Likes             *int           `json:"likes" form:"likes" gorm:"column:likes;comment:;"`                                                      //点赞数
	Github            *string        `json:"github" form:"github" gorm:"column:github;comment:;"`                                                   //GitHub链接
	Website           *string        `json:"website" form:"website" gorm:"column:website;comment:;"`                                                //网站链接
	Detail            *string        `json:"detail" form:"detail" gorm:"column:detail;comment:;type:text"`                                          //详细说明
	Logo              *string        `json:"logo" form:"logo" gorm:"column:logo;comment:;"`                                                         //logo图片路径
	Tools             datatypes.JSON `json:"tools" form:"tools" gorm:"column:tools;comment:;type:text;" swaggertype:"array,object"`                 //工具
	Uuid              *string        `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;"`                                                         //uuid
	Title             *string        `json:"title" form:"title" gorm:"column:title;comment:;"`                                                      //服务名称
	AvatarUrl         *string        `json:"avatarUrl" form:"avatarUrl" gorm:"column:avatar_url;comment:;"`                                         //头像地址
	AuthorName        *string        `json:"authorName" form:"authorName" gorm:"column:author_name;comment:;"`                                      //作者
	AuthorAvatarUrl   *string        `json:"authorAvatarUrl" form:"authorAvatarUrl" gorm:"column:author_avatar_url;comment:;"`                      //作者头像地址
	IsFeatured        *bool          `json:"isFeatured" form:"isFeatured" gorm:"column:is_featured;comment:;"`                                      //特色
	Sort              *int           `json:"sort" form:"sort" gorm:"column:sort;comment:;"`                                                         //排序
	Url               *string        `json:"url" form:"url" gorm:"column:url;comment:;"`                                                            //地址
	Type              *string        `json:"type" form:"type" gorm:"column:type;comment:;"`                                                         //类型
	UserUuid          *string        `json:"userUuid" form:"userUuid" gorm:"column:user_uuid;comment:;"`                                            //用户uuid
	SseUrl            *string        `json:"sseUrl" form:"sseUrl" gorm:"column:sse_url;comment:;"`                                                  //sseUrl
	ToolsList         []Tools        `json:"toolsList" gorm:"foreignKey:ServerId;references:id"`                                                    // 关联的工具列表
}

// TableName mcp服务 Projects自定义表名 projects
func (Projects) TableName() string {
	return "projects"
}
