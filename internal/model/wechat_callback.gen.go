// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameWechatCallback = "wechat_callback"

// WechatCallback 微信回调表
type WechatCallback struct {
	ID           int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                 // 主键ID
	ToUserName   string                `gorm:"column:to_user_name;type:varchar(255);comment:商家小程序名称" json:"to_user_name"`               // 商家小程序名称
	FromUserName string                `gorm:"column:from_user_name;type:varchar(255);comment:微信团队的 OpenID(固定值)" json:"from_user_name"` // 微信团队的 OpenID(固定值)
	MsgType      string                `gorm:"column:msg_type;type:varchar(255);comment:消息类型" json:"msg_type"`                          // 消息类型
	Event        string                `gorm:"column:event;type:varchar(255);comment:事件类型" json:"event"`                                // 事件类型
	Content      string                `gorm:"column:content;type:text;comment:内容" json:"content"`                                      // 内容
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName WechatCallback's table name
func (*WechatCallback) TableName() string {
	return TableNameWechatCallback
}