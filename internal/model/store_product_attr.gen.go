// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreProductAttr = "store_product_attr"

// StoreProductAttr 商品属性表
type StoreProductAttr struct {
	ID         int64                 `gorm:"column:id;type:mediumint;primaryKey;autoIncrement:true;comment:主键" json:"id"`   // 主键
	ProductID  int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品ID" json:"product_id"`   // 商品ID
	AttrName   string                `gorm:"column:attr_name;type:varchar(32);not null;comment:属性名" json:"attr_name"`       // 属性名
	AttrValues string                `gorm:"column:attr_values;type:varchar(1000);not null;comment:属性值" json:"attr_values"` // 属性值
	Type       int64                 `gorm:"column:type;type:tinyint(1);comment:活动类型 0=商品，1=秒杀，2=砍价，3=拼团" json:"type"`      // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
	IsDel      int64                 `gorm:"column:is_del;type:tinyint(1);not null;comment:是否删除,0-否，1-是" json:"is_del"`     // 是否删除,0-否，1-是
	CreatedAt  int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt  int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt  soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreProductAttr's table name
func (*StoreProductAttr) TableName() string {
	return TableNameStoreProductAttr
}