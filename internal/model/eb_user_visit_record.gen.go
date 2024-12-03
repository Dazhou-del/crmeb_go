// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEbUserVisitRecord = "eb_user_visit_record"

// EbUserVisitRecord 用户访问记录表
type EbUserVisitRecord struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Date      string `gorm:"column:date;comment:日期" json:"date"`               // 日期
	UID       int32  `gorm:"column:uid;comment:用户uid" json:"uid"`              // 用户uid
	VisitType int32  `gorm:"column:visit_type;comment:访问类型" json:"visit_type"` // 访问类型
}

// TableName EbUserVisitRecord's table name
func (*EbUserVisitRecord) TableName() string {
	return TableNameEbUserVisitRecord
}