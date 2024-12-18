// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreProductDescription = "store_product_description"

// StoreProductDescription 商品描述表
type StoreProductDescription struct {
	ProductID   int64                 `gorm:"column:product_id;type:int;not null;comment:商品ID" json:"product_id"`    // 商品ID
	Description string                `gorm:"column:description;type:text;not null;comment:商品详情" json:"description"` // 商品详情
	Type        int64                 `gorm:"column:type;type:tinyint(1);not null;comment:商品类型" json:"type"`         // 商品类型
	ID          int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint unsigned;not null;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint unsigned;not null;comment:修改时间" json:"updated_at"` // 修改时间
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:tinyint unsigned;not null;comment:是否删除" json:"-"`         // 是否删除
}

// TableName StoreProductDescription's table name
func (*StoreProductDescription) TableName() string {
	return TableNameStoreProductDescription
}
