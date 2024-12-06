// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbStoreProductReply = "eb_store_product_reply"

// EbStoreProductReply 评论表
type EbStoreProductReply struct {
	ID                   int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:评论ID" json:"id"`                         // 评论ID
	UID                  int32     `gorm:"column:uid;not null;comment:用户ID" json:"uid"`                                            // 用户ID
	Oid                  int32     `gorm:"column:oid;not null;comment:订单ID" json:"oid"`                                            // 订单ID
	Unique               string    `gorm:"column:unique;not null;comment:商品唯一id" json:"unique"`                                    // 商品唯一id
	ProductID            int32     `gorm:"column:product_id;not null;comment:商品id" json:"product_id"`                              // 商品id
	ReplyType            string    `gorm:"column:reply_type;not null;default:product;comment:某种商品类型(普通商品、秒杀商品）" json:"reply_type"` // 某种商品类型(普通商品、秒杀商品）
	ProductScore         bool      `gorm:"column:product_score;not null;comment:商品分数" json:"product_score"`                        // 商品分数
	ServiceScore         bool      `gorm:"column:service_score;not null;comment:服务分数" json:"service_score"`                        // 服务分数
	Comment              string    `gorm:"column:comment;not null;comment:评论内容" json:"comment"`                                    // 评论内容
	Pics                 string    `gorm:"column:pics;not null;comment:评论图片" json:"pics"`                                          // 评论图片
	MerchantReplyContent string    `gorm:"column:merchant_reply_content;comment:管理员回复内容" json:"merchant_reply_content"`            // 管理员回复内容
	MerchantReplyTime    int32     `gorm:"column:merchant_reply_time;comment:管理员回复时间" json:"merchant_reply_time"`                  // 管理员回复时间
	IsDel                int32     `gorm:"column:is_del;not null;comment:0未删除1已删除" json:"is_del"`                                  // 0未删除1已删除
	IsReply              bool      `gorm:"column:is_reply;not null;comment:0未回复1已回复" json:"is_reply"`                              // 0未回复1已回复
	Nickname             string    `gorm:"column:nickname;not null;comment:用户名称" json:"nickname"`                                  // 用户名称
	Avatar               string    `gorm:"column:avatar;not null;comment:用户头像" json:"avatar"`                                      // 用户头像
	CreateTime           time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`           // 创建时间
	UpdateTime           time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`           // 更新时间
	Sku                  string    `gorm:"column:sku;not null;comment:商品规格属性值,多个,号隔开" json:"sku"`                                  // 商品规格属性值,多个,号隔开
}

// TableName EbStoreProductReply's table name
func (*EbStoreProductReply) TableName() string {
	return TableNameEbStoreProductReply
}
