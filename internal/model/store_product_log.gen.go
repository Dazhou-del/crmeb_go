// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreProductLog = "store_product_log"

// StoreProductLog 商品日志表
type StoreProductLog struct {
	ID          int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:统计ID" json:"id"`                // 统计ID
	Type        string                `gorm:"column:type;type:varchar(10);not null;comment:类型visit,cart,order,pay,collect,refund" json:"type"` // 类型visit,cart,order,pay,collect,refund
	ProductID   int64                 `gorm:"column:product_id;type:int;not null;comment:商品ID" json:"product_id"`                              // 商品ID
	UID         int64                 `gorm:"column:uid;type:int;not null;comment:用户ID" json:"uid"`                                            // 用户ID
	VisitNum    int64                 `gorm:"column:visit_num;type:tinyint(1);not null;comment:是否浏览" json:"visit_num"`                         // 是否浏览
	CartNum     int64                 `gorm:"column:cart_num;type:int;not null;comment:加入购物车数量" json:"cart_num"`                               // 加入购物车数量
	OrderNum    int64                 `gorm:"column:order_num;type:int;not null;comment:下单数量" json:"order_num"`                                // 下单数量
	PayNum      int64                 `gorm:"column:pay_num;type:int;not null;comment:支付数量" json:"pay_num"`                                    // 支付数量
	PayPrice    decimal.Decimal       `gorm:"column:pay_price;type:decimal(10,2);not null;default:0.00;comment:支付金额" json:"pay_price"`         // 支付金额
	CostPrice   decimal.Decimal       `gorm:"column:cost_price;type:decimal(10,2);not null;default:0.00;comment:商品成本价" json:"cost_price"`      // 商品成本价
	PayUID      int64                 `gorm:"column:pay_uid;type:int;not null;comment:支付用户ID" json:"pay_uid"`                                  // 支付用户ID
	RefundNum   int64                 `gorm:"column:refund_num;type:int;not null;comment:退款数量" json:"refund_num"`                              // 退款数量
	RefundPrice decimal.Decimal       `gorm:"column:refund_price;type:decimal(10,2);not null;default:0.00;comment:退款金额" json:"refund_price"`   // 退款金额
	CollectNum  int64                 `gorm:"column:collect_num;type:tinyint(1);not null;comment:收藏" json:"collect_num"`                       // 收藏
	AddTime     int64                 `gorm:"column:add_time;type:bigint;not null;comment:添加时间" json:"add_time"`                               // 添加时间
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreProductLog's table name
func (*StoreProductLog) TableName() string {
	return TableNameStoreProductLog
}
