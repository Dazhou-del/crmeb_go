// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEbStoreProductRule = "eb_store_product_rule"

// EbStoreProductRule 商品规则值(规格)表
type EbStoreProductRule struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RuleName  string `gorm:"column:rule_name;not null;comment:规格名称" json:"rule_name"`  // 规格名称
	RuleValue string `gorm:"column:rule_value;not null;comment:规格值" json:"rule_value"` // 规格值
}

// TableName EbStoreProductRule's table name
func (*EbStoreProductRule) TableName() string {
	return TableNameEbStoreProductRule
}
