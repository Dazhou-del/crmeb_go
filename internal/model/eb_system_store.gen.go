// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbSystemStore = "eb_system_store"

// EbSystemStore 门店自提
type EbSystemStore struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null;comment:门店名称" json:"name"`                                         // 门店名称
	Introduction    string    `gorm:"column:introduction;not null;comment:简介" json:"introduction"`                           // 简介
	Phone           string    `gorm:"column:phone;not null;comment:手机号码" json:"phone"`                                       // 手机号码
	Address         string    `gorm:"column:address;not null;comment:省市区" json:"address"`                                    // 省市区
	DetailedAddress string    `gorm:"column:detailed_address;not null;comment:详细地址" json:"detailed_address"`                 // 详细地址
	Image           string    `gorm:"column:image;not null;comment:门店logo" json:"image"`                                     // 门店logo
	Latitude        string    `gorm:"column:latitude;not null;comment:纬度" json:"latitude"`                                   // 纬度
	Longitude       string    `gorm:"column:longitude;not null;comment:经度" json:"longitude"`                                 // 经度
	ValidTime       string    `gorm:"column:valid_time;not null;comment:核销有效日期" json:"valid_time"`                           // 核销有效日期
	DayTime         string    `gorm:"column:day_time;not null;comment:每日营业开关时间" json:"day_time"`                             // 每日营业开关时间
	IsShow          bool      `gorm:"column:is_show;not null;default:1;comment:是否显示" json:"is_show"`                         // 是否显示
	IsDel           bool      `gorm:"column:is_del;not null;comment:是否删除" json:"is_del"`                                     // 是否删除
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:添加时间" json:"create_time"` // 添加时间
	UpdateTime      time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName EbSystemStore's table name
func (*EbSystemStore) TableName() string {
	return TableNameEbSystemStore
}
