// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbStoreCart = "eb_store_cart"

// EbStoreCart 购物车表
type EbStoreCart struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:购物车表ID" json:"id"`                      // 购物车表ID
	UID               int32     `gorm:"column:uid;not null;comment:用户ID" json:"uid"`                                           // 用户ID
	Type              string    `gorm:"column:type;not null;comment:类型" json:"type"`                                           // 类型
	ProductID         int32     `gorm:"column:product_id;not null;comment:商品ID" json:"product_id"`                             // 商品ID
	ProductAttrUnique string    `gorm:"column:product_attr_unique;not null;comment:商品属性" json:"product_attr_unique"`           // 商品属性
	CartNum           int32     `gorm:"column:cart_num;not null;comment:商品数量" json:"cart_num"`                                 // 商品数量
	IsNew             bool      `gorm:"column:is_new;not null;comment:是否为立即购买" json:"is_new"`                                  // 是否为立即购买
	CombinationID     int32     `gorm:"column:combination_id;comment:拼团id" json:"combination_id"`                              // 拼团id
	SeckillID         int32     `gorm:"column:seckill_id;not null;comment:秒杀商品ID" json:"seckill_id"`                           // 秒杀商品ID
	BargainID         int32     `gorm:"column:bargain_id;not null;comment:砍价id" json:"bargain_id"`                             // 砍价id
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:添加时间" json:"create_time"` // 添加时间
	UpdateTime        time.Time `gorm:"column:update_time;comment:g" json:"update_time"`                                       // g
	Status            bool      `gorm:"column:status;not null;default:1;comment:购物车状态" json:"status"`                          // 购物车状态
}

// TableName EbStoreCart's table name
func (*EbStoreCart) TableName() string {
	return TableNameEbStoreCart
}