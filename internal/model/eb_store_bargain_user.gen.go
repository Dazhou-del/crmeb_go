// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEbStoreBargainUser = "eb_store_bargain_user"

// EbStoreBargainUser 用户参与砍价表
type EbStoreBargainUser struct {
	ID              int32   `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户参与砍价表ID" json:"id"`       // 用户参与砍价表ID
	UID             int32   `gorm:"column:uid;comment:用户ID" json:"uid"`                                        // 用户ID
	BargainID       int32   `gorm:"column:bargain_id;comment:砍价商品id" json:"bargain_id"`                        // 砍价商品id
	BargainPriceMin float64 `gorm:"column:bargain_price_min;comment:砍价的最低价" json:"bargain_price_min"`          // 砍价的最低价
	BargainPrice    float64 `gorm:"column:bargain_price;comment:砍价金额" json:"bargain_price"`                    // 砍价金额
	Price           float64 `gorm:"column:price;comment:砍掉的价格" json:"price"`                                   // 砍掉的价格
	Status          int32   `gorm:"column:status;not null;comment:状态 1参与中 2 活动结束参与失败 3活动结束参与成功" json:"status"` // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	AddTime         int64   `gorm:"column:add_time;comment:参与时间" json:"add_time"`                              // 参与时间
	IsDel           bool    `gorm:"column:is_del;not null;comment:是否取消" json:"is_del"`                         // 是否取消
}

// TableName EbStoreBargainUser's table name
func (*EbStoreBargainUser) TableName() string {
	return TableNameEbStoreBargainUser
}
