// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEbUserTag = "eb_user_tag"

// EbUserTag 标签管理
type EbUserTag struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;comment:标签名称" json:"name"` // 标签名称
}

// TableName EbUserTag's table name
func (*EbUserTag) TableName() string {
	return TableNameEbUserTag
}
