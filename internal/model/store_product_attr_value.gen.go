// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"google.golang.org/genproto/googleapis/type/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreProductAttrValue = "store_product_attr_value"

// StoreProductAttrValue 商品属性值表
type StoreProductAttrValue struct {
	ID           int64                 `gorm:"column:id;type:mediumint;primaryKey;autoIncrement:true;comment:主键" json:"id"`                     // 主键
	ProductID    int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品ID" json:"product_id"`                     // 商品ID
	Suk          string                `gorm:"column:suk;type:varchar(128);not null;comment:商品属性索引值 (attr_value|attr_value[|....])" json:"suk"` // 商品属性索引值 (attr_value|attr_value[|....])
	Stock        int64                 `gorm:"column:stock;type:int unsigned;not null;comment:属性对应的库存" json:"stock"`                            // 属性对应的库存
	Sales        int64                 `gorm:"column:sales;type:int unsigned;not null;comment:销量" json:"sales"`                                 // 销量
	Price        decimal.Decimal       `gorm:"column:price;type:decimal(8,2) unsigned;not null;comment:属性金额" json:"price"`                      // 属性金额
	Image        string                `gorm:"column:image;type:varchar(1000);comment:图片" json:"image"`                                         // 图片
	Unique       string                `gorm:"column:unique;type:char(8);not null;comment:唯一值" json:"unique"`                                   // 唯一值
	Cost         decimal.Decimal       `gorm:"column:cost;type:decimal(8,2) unsigned;not null;default:0.00;comment:成本价" json:"cost"`            // 成本价
	BarCode      string                `gorm:"column:bar_code;type:varchar(50);not null;comment:商品条码" json:"bar_code"`                          // 商品条码
	OtPrice      decimal.Decimal       `gorm:"column:ot_price;type:decimal(8,2);not null;default:0.00;comment:原价" json:"ot_price"`              // 原价
	Weight       decimal.Decimal       `gorm:"column:weight;type:decimal(8,2);not null;default:0.00;comment:重量" json:"weight"`                  // 重量
	Volume       decimal.Decimal       `gorm:"column:volume;type:decimal(8,2);not null;default:0.00;comment:体积" json:"volume"`                  // 体积
	Brokerage    decimal.Decimal       `gorm:"column:brokerage;type:decimal(8,2);not null;default:0.00;comment:一级返佣" json:"brokerage"`          // 一级返佣
	BrokerageTwo decimal.Decimal       `gorm:"column:brokerage_two;type:decimal(8,2);not null;default:0.00;comment:二级返佣" json:"brokerage_two"`  // 二级返佣
	Type         int64                 `gorm:"column:type;type:tinyint(1);comment:活动类型 0=商品，1=秒杀，2=砍价，3=拼团" json:"type"`                        // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
	Quota        int64                 `gorm:"column:quota;type:int;comment:活动限购数量" json:"quota"`                                               // 活动限购数量
	QuotaShow    int64                 `gorm:"column:quota_show;type:int;comment:活动限购数量显示" json:"quota_show"`                                   // 活动限购数量显示
	AttrValue    string                `gorm:"column:attr_value;type:text;comment:attr_values 创建更新时的属性对应" json:"attr_value"`                    // attr_values 创建更新时的属性对应
	IsDel        int64                 `gorm:"column:is_del;type:tinyint(1);not null;comment:是否删除,0-否，1-是" json:"is_del"`                       // 是否删除,0-否，1-是
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreProductAttrValue's table name
func (*StoreProductAttrValue) TableName() string {
	return TableNameStoreProductAttrValue
}
