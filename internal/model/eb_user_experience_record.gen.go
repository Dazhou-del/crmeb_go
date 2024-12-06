// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbUserExperienceRecord = "eb_user_experience_record"

// EbUserExperienceRecord 用户经验记录表
type EbUserExperienceRecord struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:记录id" json:"id"`                           // 记录id
	UID        int32     `gorm:"column:uid;not null;comment:用户uid" json:"uid"`                                             // 用户uid
	LinkID     string    `gorm:"column:link_id;not null;default:0;comment:关联id-orderNo,(sign,system默认为0）" json:"link_id"`  // 关联id-orderNo,(sign,system默认为0）
	LinkType   string    `gorm:"column:link_type;not null;default:order;comment:关联类型（order,sign,system）" json:"link_type"` // 关联类型（order,sign,system）
	Type       int32     `gorm:"column:type;not null;default:1;comment:类型：1-增加，2-扣减" json:"type"`                          // 类型：1-增加，2-扣减
	Title      string    `gorm:"column:title;not null;comment:标题" json:"title"`                                            // 标题
	Experience int32     `gorm:"column:experience;not null;comment:经验" json:"experience"`                                  // 经验
	Balance    int32     `gorm:"column:balance;not null;comment:剩余" json:"balance"`                                        // 剩余
	Mark       string    `gorm:"column:mark;not null;comment:备注" json:"mark"`                                              // 备注
	Status     bool      `gorm:"column:status;not null;default:1;comment:状态：1-成功（保留字段）" json:"status"`                     // 状态：1-成功（保留字段）
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:添加时间" json:"create_time"`             // 添加时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`             // 更新时间
}

// TableName EbUserExperienceRecord's table name
func (*EbUserExperienceRecord) TableName() string {
	return TableNameEbUserExperienceRecord
}
