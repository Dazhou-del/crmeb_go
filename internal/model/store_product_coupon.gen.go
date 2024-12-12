// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreProductCoupon = "store_product_coupon"

// StoreProductCoupon 商品优惠券表
type StoreProductCoupon struct {
	ID            int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	ProductID     int64                 `gorm:"column:product_id;type:int;not null;comment:商品id" json:"product_id"`            // 商品id
	IssueCouponID int64                 `gorm:"column:issue_coupon_id;type:int;not null;comment:优惠劵id" json:"issue_coupon_id"` // 优惠劵id
	AddTime       int64                 `gorm:"column:add_time;type:int;not null;comment:添加时间" json:"add_time"`                // 添加时间
	CreatedAt     int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt     int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreProductCoupon's table name
func (*StoreProductCoupon) TableName() string {
	return TableNameStoreProductCoupon
}