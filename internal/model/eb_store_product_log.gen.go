// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEbStoreProductLog = "eb_store_product_log"

// EbStoreProductLog 商品日志表
type EbStoreProductLog struct {
	ID          int32   `gorm:"column:id;primaryKey;autoIncrement:true;comment:统计ID" json:"id"`                 // 统计ID
	Type        string  `gorm:"column:type;not null;comment:类型visit,cart,order,pay,collect,refund" json:"type"` // 类型visit,cart,order,pay,collect,refund
	ProductID   int32   `gorm:"column:product_id;not null;comment:商品ID" json:"product_id"`                      // 商品ID
	UID         int32   `gorm:"column:uid;not null;comment:用户ID" json:"uid"`                                    // 用户ID
	VisitNum    bool    `gorm:"column:visit_num;not null;comment:是否浏览" json:"visit_num"`                        // 是否浏览
	CartNum     int32   `gorm:"column:cart_num;not null;comment:加入购物车数量" json:"cart_num"`                       // 加入购物车数量
	OrderNum    int32   `gorm:"column:order_num;not null;comment:下单数量" json:"order_num"`                        // 下单数量
	PayNum      int32   `gorm:"column:pay_num;not null;comment:支付数量" json:"pay_num"`                            // 支付数量
	PayPrice    float64 `gorm:"column:pay_price;not null;default:0.00;comment:支付金额" json:"pay_price"`           // 支付金额
	CostPrice   float64 `gorm:"column:cost_price;not null;default:0.00;comment:商品成本价" json:"cost_price"`        // 商品成本价
	PayUID      int32   `gorm:"column:pay_uid;not null;comment:支付用户ID" json:"pay_uid"`                          // 支付用户ID
	RefundNum   int32   `gorm:"column:refund_num;not null;comment:退款数量" json:"refund_num"`                      // 退款数量
	RefundPrice float64 `gorm:"column:refund_price;not null;default:0.00;comment:退款金额" json:"refund_price"`     // 退款金额
	CollectNum  bool    `gorm:"column:collect_num;not null;comment:收藏" json:"collect_num"`                      // 收藏
	AddTime     int64   `gorm:"column:add_time;not null;comment:添加时间" json:"add_time"`                          // 添加时间
}

// TableName EbStoreProductLog's table name
func (*EbStoreProductLog) TableName() string {
	return TableNameEbStoreProductLog
}
