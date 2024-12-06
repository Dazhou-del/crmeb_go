// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbStoreSeckill = "eb_store_seckill"

// EbStoreSeckill 商品秒杀产品表
type EbStoreSeckill struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:商品秒杀产品表id" json:"id"`                   // 商品秒杀产品表id
	ProductID    int32     `gorm:"column:product_id;not null;comment:商品id" json:"product_id"`                             // 商品id
	Image        string    `gorm:"column:image;not null;comment:推荐图" json:"image"`                                        // 推荐图
	Images       string    `gorm:"column:images;not null;comment:轮播图" json:"images"`                                      // 轮播图
	Title        string    `gorm:"column:title;not null;comment:活动标题" json:"title"`                                       // 活动标题
	Info         string    `gorm:"column:info;not null;comment:简介" json:"info"`                                           // 简介
	Price        float64   `gorm:"column:price;not null;comment:价格" json:"price"`                                         // 价格
	Cost         float64   `gorm:"column:cost;not null;default:0.00;comment:成本" json:"cost"`                              // 成本
	OtPrice      float64   `gorm:"column:ot_price;not null;default:0.00;comment:原价" json:"ot_price"`                      // 原价
	GiveIntegral float64   `gorm:"column:give_integral;not null;default:0.00;comment:返多少积分" json:"give_integral"`         // 返多少积分
	Sort         int32     `gorm:"column:sort;not null;comment:排序" json:"sort"`                                           // 排序
	Stock        int32     `gorm:"column:stock;not null;comment:库存" json:"stock"`                                         // 库存
	Sales        int32     `gorm:"column:sales;not null;comment:销量" json:"sales"`                                         // 销量
	UnitName     string    `gorm:"column:unit_name;not null;comment:单位名" json:"unit_name"`                                // 单位名
	Postage      float64   `gorm:"column:postage;not null;default:0.00;comment:邮费" json:"postage"`                        // 邮费
	Description  string    `gorm:"column:description;comment:内容" json:"description"`                                      // 内容
	StartTime    time.Time `gorm:"column:start_time;not null;comment:开始时间" json:"start_time"`                             // 开始时间
	StopTime     time.Time `gorm:"column:stop_time;not null;comment:结束时间" json:"stop_time"`                               // 结束时间
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:添加时间" json:"create_time"` // 添加时间
	Status       int32     `gorm:"column:status;not null;comment:秒杀状态 0=关闭 1=开启" json:"status"`                           // 秒杀状态 0=关闭 1=开启
	IsPostage    int32     `gorm:"column:is_postage;not null;comment:是否包邮" json:"is_postage"`                             // 是否包邮
	IsDel        int32     `gorm:"column:is_del;not null;comment:删除 0未删除1已删除" json:"is_del"`                              // 删除 0未删除1已删除
	Num          int32     `gorm:"column:num;not null;comment:当天参与活动次数" json:"num"`                                       // 当天参与活动次数
	IsShow       int32     `gorm:"column:is_show;not null;default:1;comment:显示" json:"is_show"`                           // 显示
	TimeID       int32     `gorm:"column:time_id;comment:时间段ID" json:"time_id"`                                           // 时间段ID
	TempID       int32     `gorm:"column:temp_id;not null;comment:运费模板ID" json:"temp_id"`                                 // 运费模板ID
	Weight       float64   `gorm:"column:weight;not null;default:0.00;comment:重量" json:"weight"`                          // 重量
	Volume       float64   `gorm:"column:volume;not null;default:0.00;comment:体积" json:"volume"`                          // 体积
	Quota        int32     `gorm:"column:quota;not null;comment:限购总数,随减" json:"quota"`                                    // 限购总数,随减
	QuotaShow    int32     `gorm:"column:quota_show;not null;comment:限购总数显示.不变" json:"quota_show"`                        // 限购总数显示.不变
	SpecType     bool      `gorm:"column:spec_type;not null;comment:规格 0=单 1=多" json:"spec_type"`                         // 规格 0=单 1=多
}

// TableName EbStoreSeckill's table name
func (*EbStoreSeckill) TableName() string {
	return TableNameEbStoreSeckill
}
