// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameCategory = "category"

// Category 分类表
type Category struct {
	ID        int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Pid       int64                 `gorm:"column:pid;type:int unsigned;not null;comment:父级ID" json:"pid"`                                                  // 父级ID
	Path      string                `gorm:"column:path;type:varchar(255);not null;default:/0/;comment:路径" json:"path"`                                      // 路径
	Name      string                `gorm:"column:name;type:varchar(50);not null;comment:分类名称" json:"name"`                                                 // 分类名称
	Type      int64                 `gorm:"column:type;type:smallint;default:1;comment:类型，1 产品分类，2 附件分类，3 文章分类， 4 设置分类， 5 菜单分类，6 配置分类， 7 秒杀配置" json:"type"` // 类型，1 产品分类，2 附件分类，3 文章分类， 4 设置分类， 5 菜单分类，6 配置分类， 7 秒杀配置
	URL       string                `gorm:"column:url;type:varchar(255);comment:地址" json:"url"`                                                             // 地址
	Extra     string                `gorm:"column:extra;type:text;comment:扩展字段 Jsos格式" json:"extra"`                                                        // 扩展字段 Jsos格式
	Status    int64                 `gorm:"column:status;type:tinyint(1);not null;default:1;comment:状态, 1正常，0失效" json:"status"`                             // 状态, 1正常，0失效
	Sort      int64                 `gorm:"column:sort;type:int;not null;default:99999;comment:排序" json:"sort"`                                             // 排序
	CreatedAt int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
