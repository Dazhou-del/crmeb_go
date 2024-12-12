// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSystemGroupDatum = "system_group_data"

// SystemGroupDatum 组合数据详情表
type SystemGroupDatum struct {
	ID        int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:组合数据详情ID" json:"id"`           // 组合数据详情ID
	Gid       int64                 `gorm:"column:gid;type:int;not null;comment:对应的数据组id" json:"gid"`                              // 对应的数据组id
	Value     string                `gorm:"column:value;type:text;not null;comment:数据组对应的数据值（json数据）" json:"value"`                // 数据组对应的数据值（json数据）
	Sort      int64                 `gorm:"column:sort;type:int;not null;comment:数据排序" json:"sort"`                                // 数据排序
	Status    int64                 `gorm:"column:status;type:tinyint(1);not null;default:1;comment:状态（1：开启；2：关闭；）" json:"status"` // 状态（1：开启；2：关闭；）
	CreatedAt int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName SystemGroupDatum's table name
func (*SystemGroupDatum) TableName() string {
	return TableNameSystemGroupDatum
}