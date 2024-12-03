// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbWechatCallback = "eb_wechat_callback"

// EbWechatCallback 微信回调表
type EbWechatCallback struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`         // 主键ID
	ToUserName   string    `gorm:"column:to_user_name;comment:商家小程序名称" json:"to_user_name"`                // 商家小程序名称
	FromUserName string    `gorm:"column:from_user_name;comment:微信团队的 OpenID(固定值)" json:"from_user_name"`  // 微信团队的 OpenID(固定值)
	CreateTime   int64     `gorm:"column:create_time;comment:事件时间,Unix时间戳" json:"create_time"`             // 事件时间,Unix时间戳
	MsgType      string    `gorm:"column:msg_type;comment:消息类型" json:"msg_type"`                           // 消息类型
	Event        string    `gorm:"column:event;comment:事件类型" json:"event"`                                 // 事件类型
	Content      string    `gorm:"column:content;comment:内容" json:"content"`                               // 内容
	AddTime      time.Time `gorm:"column:add_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"add_time"` // 创建时间
}

// TableName EbWechatCallback's table name
func (*EbWechatCallback) TableName() string {
	return TableNameEbWechatCallback
}
