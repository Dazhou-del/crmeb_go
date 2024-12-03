// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbSystemGroupDatum = "eb_system_group_data"

// EbSystemGroupDatum 组合数据详情表
type EbSystemGroupDatum struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:组合数据详情ID" json:"id"`           // 组合数据详情ID
	Gid        int32     `gorm:"column:gid;not null;comment:对应的数据组id" json:"gid"`                              // 对应的数据组id
	Value      string    `gorm:"column:value;not null;comment:数据组对应的数据值（json数据）" json:"value"`                 // 数据组对应的数据值（json数据）
	Sort       int32     `gorm:"column:sort;not null;comment:数据排序" json:"sort"`                                // 数据排序
	Status     bool      `gorm:"column:status;not null;default:1;comment:状态（1：开启；2：关闭；）" json:"status"`        // 状态（1：开启；2：关闭；）
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName EbSystemGroupDatum's table name
func (*EbSystemGroupDatum) TableName() string {
	return TableNameEbSystemGroupDatum
}