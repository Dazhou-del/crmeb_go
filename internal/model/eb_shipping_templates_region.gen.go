// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbShippingTemplatesRegion = "eb_shipping_templates_region"

// EbShippingTemplatesRegion 运费模板指定区域费用
type EbShippingTemplatesRegion struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                          // 编号
	TempID       int32     `gorm:"column:temp_id;not null;comment:模板ID" json:"temp_id"`                                   // 模板ID
	CityID       int32     `gorm:"column:city_id;not null;comment:城市ID" json:"city_id"`                                   // 城市ID
	Title        string    `gorm:"column:title;comment:描述" json:"title"`                                                  // 描述
	First        float64   `gorm:"column:first;not null;default:0.00;comment:首件" json:"first"`                            // 首件
	FirstPrice   float64   `gorm:"column:first_price;not null;default:0.00;comment:首件运费" json:"first_price"`              // 首件运费
	Renewal      float64   `gorm:"column:renewal;not null;default:0.00;comment:续件" json:"renewal"`                        // 续件
	RenewalPrice float64   `gorm:"column:renewal_price;not null;default:0.00;comment:续件运费" json:"renewal_price"`          // 续件运费
	Type         bool      `gorm:"column:type;not null;default:1;comment:计费方式 1按件数 2按重量 3按体积" json:"type"`                // 计费方式 1按件数 2按重量 3按体积
	Uniqid       string    `gorm:"column:uniqid;not null;comment:分组唯一值" json:"uniqid"`                                    // 分组唯一值
	Status       bool      `gorm:"column:status;comment:是否无效" json:"status"`                                              // 是否无效
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName EbShippingTemplatesRegion's table name
func (*EbShippingTemplatesRegion) TableName() string {
	return TableNameEbShippingTemplatesRegion
}
