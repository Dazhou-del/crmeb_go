// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbUserSign = "eb_user_sign"

// EbUserSign 签到记录表
type EbUserSign struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID        int32     `gorm:"column:uid;not null;comment:用户uid" json:"uid"`                                          // 用户uid
	Title      string    `gorm:"column:title;not null;comment:签到说明" json:"title"`                                       // 签到说明
	Number     int32     `gorm:"column:number;not null;comment:获得" json:"number"`                                       // 获得
	Balance    int32     `gorm:"column:balance;not null;comment:剩余" json:"balance"`                                     // 剩余
	Type       bool      `gorm:"column:type;not null;default:1;comment:类型，1积分，2经验" json:"type"`                         // 类型，1积分，2经验
	CreateDay  time.Time `gorm:"column:create_day;not null;comment:签到日期" json:"create_day"`                             // 签到日期
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:添加时间" json:"create_time"` // 添加时间
}

// TableName EbUserSign's table name
func (*EbUserSign) TableName() string {
	return TableNameEbUserSign
}
