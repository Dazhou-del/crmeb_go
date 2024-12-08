// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreProductAttrResult = "store_product_attr_result"

// StoreProductAttrResult 商品属性详情表
type StoreProductAttrResult struct {
	ID         int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键" json:"id"`           // 主键
	ProductID  int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品ID" json:"product_id"`     // 商品ID
	Result     string                `gorm:"column:result;type:longtext;not null;comment:商品属性参数" json:"result"`               // 商品属性参数
	ChangeTime int64                 `gorm:"column:change_time;type:int unsigned;not null;comment:上次修改时间" json:"change_time"` // 上次修改时间
	Type       int64                 `gorm:"column:type;type:tinyint(1);comment:活动类型 0=商品，1=秒杀，2=砍价，3=拼团" json:"type"`        // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
	CreatedAt  int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt  int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt  soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreProductAttrResult's table name
func (*StoreProductAttrResult) TableName() string {
	return TableNameStoreProductAttrResult
}
