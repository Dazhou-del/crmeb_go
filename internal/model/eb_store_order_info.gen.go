// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbStoreOrderInfo = "eb_store_order_info"

// EbStoreOrderInfo 订单购物详情表
type EbStoreOrderInfo struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                            // 主键
	OrderID      int32     `gorm:"column:order_id;not null;comment:订单id" json:"order_id"`                                   // 订单id
	ProductID    int32     `gorm:"column:product_id;not null;comment:商品ID" json:"product_id"`                               // 商品ID
	Info         string    `gorm:"column:info;not null;comment:购买东西的详细信息" json:"info"`                                      // 购买东西的详细信息
	Unique       string    `gorm:"column:unique;not null;comment:唯一id" json:"unique"`                                       // 唯一id
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`            // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`   // 更新时间
	OrderNo      string    `gorm:"column:order_no;not null;comment:订单号" json:"order_no"`                                    // 订单号
	ProductName  string    `gorm:"column:product_name;not null;comment:商品名称" json:"product_name"`                           // 商品名称
	AttrValueID  int32     `gorm:"column:attr_value_id;comment:规格属性值id" json:"attr_value_id"`                               // 规格属性值id
	Image        string    `gorm:"column:image;not null;comment:商品图片" json:"image"`                                         // 商品图片
	Sku          string    `gorm:"column:sku;not null;comment:商品sku" json:"sku"`                                            // 商品sku
	Price        float64   `gorm:"column:price;not null;comment:商品价格" json:"price"`                                         // 商品价格
	PayNum       int32     `gorm:"column:pay_num;not null;comment:购买数量" json:"pay_num"`                                     // 购买数量
	Weight       float64   `gorm:"column:weight;not null;comment:重量" json:"weight"`                                         // 重量
	Volume       float64   `gorm:"column:volume;not null;comment:体积" json:"volume"`                                         // 体积
	GiveIntegral int32     `gorm:"column:give_integral;not null;comment:赠送积分" json:"give_integral"`                         // 赠送积分
	IsReply      bool      `gorm:"column:is_reply;not null;comment:是否评价，0-未评价，1-已评价" json:"is_reply"`                       // 是否评价，0-未评价，1-已评价
	IsSub        bool      `gorm:"column:is_sub;not null;comment:是否单独分佣,0-否，1-是" json:"is_sub"`                             // 是否单独分佣,0-否，1-是
	VipPrice     float64   `gorm:"column:vip_price;not null;comment:会员价" json:"vip_price"`                                  // 会员价
	ProductType  int32     `gorm:"column:product_type;not null;comment:商品类型:0-普通，1-秒杀，2-砍价，3-拼团，4-视频号" json:"product_type"` // 商品类型:0-普通，1-秒杀，2-砍价，3-拼团，4-视频号
}

// TableName EbStoreOrderInfo's table name
func (*EbStoreOrderInfo) TableName() string {
	return TableNameEbStoreOrderInfo
}
