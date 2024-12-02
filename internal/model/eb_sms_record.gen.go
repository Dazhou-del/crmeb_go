// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEbSmsRecord = "eb_sms_record"

// EbSmsRecord 短信发送记录表
type EbSmsRecord struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:短信发送记录编号" json:"id"`                         // 短信发送记录编号
	UID        string    `gorm:"column:uid;not null;comment:短信平台账号" json:"uid"`                                              // 短信平台账号
	Phone      string    `gorm:"column:phone;not null;comment:接受短信的手机号" json:"phone"`                                        // 接受短信的手机号
	Content    string    `gorm:"column:content;comment:短信内容" json:"content"`                                                 // 短信内容
	AddIP      string    `gorm:"column:add_ip;comment:添加记录ip" json:"add_ip"`                                                 // 添加记录ip
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`               // 创建时间
	Template   string    `gorm:"column:template;comment:短信模板ID" json:"template"`                                             // 短信模板ID
	Resultcode int32     `gorm:"column:resultcode;comment:状态码 100=成功,130=失败,131=空号,132=停机,133=关机,134=无状态" json:"resultcode"` // 状态码 100=成功,130=失败,131=空号,132=停机,133=关机,134=无状态
	RecordID   int32     `gorm:"column:record_id;comment:发送记录id" json:"record_id"`                                           // 发送记录id
	Memo       string    `gorm:"column:memo;comment:短信平台返回信息" json:"memo"`                                                   // 短信平台返回信息
}

// TableName EbSmsRecord's table name
func (*EbSmsRecord) TableName() string {
	return TableNameEbSmsRecord
}
