// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreCart = "store_cart"

// StoreCart 购物车表
type StoreCart struct {
	ID                int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:购物车表ID" json:"id"`        // 购物车表ID
	UID               int64                 `gorm:"column:uid;type:int unsigned;not null;comment:用户ID" json:"uid"`                                // 用户ID
	Type              string                `gorm:"column:type;type:varchar(32);not null;comment:类型" json:"type"`                                 // 类型
	ProductID         int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品ID" json:"product_id"`                  // 商品ID
	ProductAttrUnique string                `gorm:"column:product_attr_unique;type:varchar(16);not null;comment:商品属性" json:"product_attr_unique"` // 商品属性
	CartNum           int64                 `gorm:"column:cart_num;type:smallint unsigned;not null;comment:商品数量" json:"cart_num"`                 // 商品数量
	IsNew             int64                 `gorm:"column:is_new;type:tinyint(1);not null;comment:是否为立即购买" json:"is_new"`                         // 是否为立即购买
	CombinationID     int64                 `gorm:"column:combination_id;type:int unsigned;comment:拼团id" json:"combination_id"`                   // 拼团id
	SeckillID         int64                 `gorm:"column:seckill_id;type:int unsigned;not null;comment:秒杀商品ID" json:"seckill_id"`                // 秒杀商品ID
	BargainID         int64                 `gorm:"column:bargain_id;type:int unsigned;not null;comment:砍价id" json:"bargain_id"`                  // 砍价id
	Status            int64                 `gorm:"column:status;type:tinyint(1);not null;default:1;comment:购物车状态" json:"status"`                 // 购物车状态
	CreatedAt         int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt         int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt         soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreCart's table name
func (*StoreCart) TableName() string {
	return TableNameStoreCart
}