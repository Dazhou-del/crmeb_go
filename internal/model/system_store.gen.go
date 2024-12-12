// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSystemStore = "system_store"

// SystemStore 门店自提
type SystemStore struct {
	ID              int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name            string                `gorm:"column:name;type:varchar(100);not null;comment:门店名称" json:"name"`                         // 门店名称
	Introduction    string                `gorm:"column:introduction;type:varchar(1000);not null;comment:简介" json:"introduction"`          // 简介
	Phone           string                `gorm:"column:phone;type:char(25);not null;comment:手机号码" json:"phone"`                           // 手机号码
	Address         string                `gorm:"column:address;type:varchar(255);not null;comment:省市区" json:"address"`                    // 省市区
	DetailedAddress string                `gorm:"column:detailed_address;type:varchar(255);not null;comment:详细地址" json:"detailed_address"` // 详细地址
	Image           string                `gorm:"column:image;type:varchar(255);not null;comment:门店logo" json:"image"`                     // 门店logo
	Latitude        string                `gorm:"column:latitude;type:char(25);not null;comment:纬度" json:"latitude"`                       // 纬度
	Longitude       string                `gorm:"column:longitude;type:char(25);not null;comment:经度" json:"longitude"`                     // 经度
	ValidTime       string                `gorm:"column:valid_time;type:varchar(100);not null;comment:核销有效日期" json:"valid_time"`           // 核销有效日期
	DayTime         string                `gorm:"column:day_time;type:varchar(100);not null;comment:每日营业开关时间" json:"day_time"`             // 每日营业开关时间
	IsShow          int64                 `gorm:"column:is_show;type:tinyint(1);not null;default:1;comment:是否显示" json:"is_show"`           // 是否显示
	IsDel           int64                 `gorm:"column:is_del;type:tinyint(1);not null;comment:是否删除" json:"is_del"`                       // 是否删除
	CreatedAt       int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt       int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName SystemStore's table name
func (*SystemStore) TableName() string {
	return TableNameSystemStore
}