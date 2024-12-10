// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"crmeb_go/internal/model"
)

func newStoreOrder(db *gorm.DB, opts ...gen.DOOption) storeOrder {
	_storeOrder := storeOrder{}

	_storeOrder.storeOrderDo.UseDB(db, opts...)
	_storeOrder.storeOrderDo.UseModel(&model.StoreOrder{})

	tableName := _storeOrder.storeOrderDo.TableName()
	_storeOrder.ALL = field.NewAsterisk(tableName)
	_storeOrder.ID = field.NewInt64(tableName, "id")
	_storeOrder.OrderID = field.NewString(tableName, "order_id")
	_storeOrder.UID = field.NewInt64(tableName, "uid")
	_storeOrder.RealName = field.NewString(tableName, "real_name")
	_storeOrder.UserPhone = field.NewString(tableName, "user_phone")
	_storeOrder.UserAddress = field.NewString(tableName, "user_address")
	_storeOrder.FreightPrice = field.NewString(tableName, "freight_price")
	_storeOrder.TotalNum = field.NewInt64(tableName, "total_num")
	_storeOrder.TotalPrice = field.NewString(tableName, "total_price")
	_storeOrder.TotalPostage = field.NewString(tableName, "total_postage")
	_storeOrder.PayPrice = field.NewString(tableName, "pay_price")
	_storeOrder.PayPostage = field.NewString(tableName, "pay_postage")
	_storeOrder.DeductionPrice = field.NewString(tableName, "deduction_price")
	_storeOrder.CouponID = field.NewInt64(tableName, "coupon_id")
	_storeOrder.CouponPrice = field.NewString(tableName, "coupon_price")
	_storeOrder.Paid = field.NewInt64(tableName, "paid")
	_storeOrder.PayTime = field.NewInt64(tableName, "pay_time")
	_storeOrder.PayType = field.NewString(tableName, "pay_type")
	_storeOrder.Status = field.NewInt64(tableName, "status")
	_storeOrder.RefundStatus = field.NewInt64(tableName, "refund_status")
	_storeOrder.RefundReasonWapImg = field.NewString(tableName, "refund_reason_wap_img")
	_storeOrder.RefundReasonWapExplain = field.NewString(tableName, "refund_reason_wap_explain")
	_storeOrder.RefundReasonWap = field.NewString(tableName, "refund_reason_wap")
	_storeOrder.RefundReason = field.NewString(tableName, "refund_reason")
	_storeOrder.RefundReasonTime = field.NewInt64(tableName, "refund_reason_time")
	_storeOrder.RefundPrice = field.NewString(tableName, "refund_price")
	_storeOrder.DeliveryName = field.NewString(tableName, "delivery_name")
	_storeOrder.DeliveryType = field.NewString(tableName, "delivery_type")
	_storeOrder.DeliveryID = field.NewString(tableName, "delivery_id")
	_storeOrder.GainIntegral = field.NewInt64(tableName, "gain_integral")
	_storeOrder.UseIntegral = field.NewInt64(tableName, "use_integral")
	_storeOrder.BackIntegral = field.NewInt64(tableName, "back_integral")
	_storeOrder.Mark = field.NewString(tableName, "mark")
	_storeOrder.IsDel = field.NewInt64(tableName, "is_del")
	_storeOrder.Remark = field.NewString(tableName, "remark")
	_storeOrder.MerID = field.NewInt64(tableName, "mer_id")
	_storeOrder.IsMerCheck = field.NewInt64(tableName, "is_mer_check")
	_storeOrder.CombinationID = field.NewInt64(tableName, "combination_id")
	_storeOrder.PinkID = field.NewInt64(tableName, "pink_id")
	_storeOrder.Cost = field.NewString(tableName, "cost")
	_storeOrder.SeckillID = field.NewInt64(tableName, "seckill_id")
	_storeOrder.BargainID = field.NewInt64(tableName, "bargain_id")
	_storeOrder.VerifyCode = field.NewString(tableName, "verify_code")
	_storeOrder.StoreID = field.NewInt64(tableName, "store_id")
	_storeOrder.ShippingType = field.NewInt64(tableName, "shipping_type")
	_storeOrder.ClerkID = field.NewInt64(tableName, "clerk_id")
	_storeOrder.IsChannel = field.NewInt64(tableName, "is_channel")
	_storeOrder.IsRemind = field.NewInt64(tableName, "is_remind")
	_storeOrder.IsSystemDel = field.NewInt64(tableName, "is_system_del")
	_storeOrder.DeliveryCode = field.NewString(tableName, "delivery_code")
	_storeOrder.BargainUserID = field.NewInt64(tableName, "bargain_user_id")
	_storeOrder.Type = field.NewInt64(tableName, "type")
	_storeOrder.ProTotalPrice = field.NewString(tableName, "pro_total_price")
	_storeOrder.BeforePayPrice = field.NewString(tableName, "before_pay_price")
	_storeOrder.IsAlterPrice = field.NewInt64(tableName, "is_alter_price")
	_storeOrder.OutTradeNo = field.NewString(tableName, "out_trade_no")
	_storeOrder.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeOrder.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeOrder.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeOrder.fillFieldMap()

	return _storeOrder
}

// storeOrder 订单表
type storeOrder struct {
	storeOrderDo storeOrderDo

	ALL                    field.Asterisk
	ID                     field.Int64  // 订单ID
	OrderID                field.String // 订单号
	UID                    field.Int64  // 用户id
	RealName               field.String // 用户姓名
	UserPhone              field.String // 用户电话
	UserAddress            field.String // 详细地址
	FreightPrice           field.String // 运费金额
	TotalNum               field.Int64  // 订单商品总数
	TotalPrice             field.String // 订单总价
	TotalPostage           field.String // 邮费
	PayPrice               field.String // 实际支付金额
	PayPostage             field.String // 支付邮费
	DeductionPrice         field.String // 抵扣金额
	CouponID               field.Int64  // 优惠券id
	CouponPrice            field.String // 优惠券金额
	Paid                   field.Int64  // 支付状态
	PayTime                field.Int64
	PayType                field.String // 支付方式
	Status                 field.Int64  // 订单状态（0：待发货；1：待收货；2：已收货，待评价；3：已完成；）
	RefundStatus           field.Int64  // 0 未退款 1 申请中 2 已退款 3 退款中
	RefundReasonWapImg     field.String // 退款图片
	RefundReasonWapExplain field.String // 退款用户说明
	RefundReasonWap        field.String // 前台退款原因
	RefundReason           field.String // 不退款的理由
	RefundReasonTime       field.Int64
	RefundPrice            field.String // 退款金额
	DeliveryName           field.String // 快递名称/送货人姓名
	DeliveryType           field.String // 发货类型
	DeliveryID             field.String // 快递单号/手机号
	GainIntegral           field.Int64  // 消费赚取积分
	UseIntegral            field.Int64  // 使用积分
	BackIntegral           field.Int64  // 给用户退了多少积分
	Mark                   field.String // 备注
	IsDel                  field.Int64  // 是否删除
	Remark                 field.String // 管理员备注
	MerID                  field.Int64  // 商户ID
	IsMerCheck             field.Int64
	CombinationID          field.Int64  // 拼团商品id0一般商品
	PinkID                 field.Int64  // 拼团id 0没有拼团
	Cost                   field.String // 成本价
	SeckillID              field.Int64  // 秒杀商品ID
	BargainID              field.Int64  // 砍价id
	VerifyCode             field.String // 核销码
	StoreID                field.Int64  // 门店id
	ShippingType           field.Int64  // 配送方式 1=快递 ，2=门店自提
	ClerkID                field.Int64  // 店员id/核销员id
	IsChannel              field.Int64  // 支付渠道(0微信公众号1微信小程序2余额)
	IsRemind               field.Int64  // 消息提醒
	IsSystemDel            field.Int64  // 后台是否删除
	DeliveryCode           field.String // 快递公司简称
	BargainUserID          field.Int64  // 用户拼团活动id 0没有
	Type                   field.Int64  // 订单类型:0-普通订单，1-视频号订单
	ProTotalPrice          field.String // 商品总价
	BeforePayPrice         field.String // 改价前支付金额
	IsAlterPrice           field.Int64  // 是否改价,0-否，1-是
	OutTradeNo             field.String // 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
	CreatedAt              field.Int64
	UpdatedAt              field.Int64
	DeletedAt              field.Field

	fieldMap map[string]field.Expr
}

func (s storeOrder) Table(newTableName string) *storeOrder {
	s.storeOrderDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeOrder) As(alias string) *storeOrder {
	s.storeOrderDo.DO = *(s.storeOrderDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeOrder) updateTableName(table string) *storeOrder {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.OrderID = field.NewString(table, "order_id")
	s.UID = field.NewInt64(table, "uid")
	s.RealName = field.NewString(table, "real_name")
	s.UserPhone = field.NewString(table, "user_phone")
	s.UserAddress = field.NewString(table, "user_address")
	s.FreightPrice = field.NewString(table, "freight_price")
	s.TotalNum = field.NewInt64(table, "total_num")
	s.TotalPrice = field.NewString(table, "total_price")
	s.TotalPostage = field.NewString(table, "total_postage")
	s.PayPrice = field.NewString(table, "pay_price")
	s.PayPostage = field.NewString(table, "pay_postage")
	s.DeductionPrice = field.NewString(table, "deduction_price")
	s.CouponID = field.NewInt64(table, "coupon_id")
	s.CouponPrice = field.NewString(table, "coupon_price")
	s.Paid = field.NewInt64(table, "paid")
	s.PayTime = field.NewInt64(table, "pay_time")
	s.PayType = field.NewString(table, "pay_type")
	s.Status = field.NewInt64(table, "status")
	s.RefundStatus = field.NewInt64(table, "refund_status")
	s.RefundReasonWapImg = field.NewString(table, "refund_reason_wap_img")
	s.RefundReasonWapExplain = field.NewString(table, "refund_reason_wap_explain")
	s.RefundReasonWap = field.NewString(table, "refund_reason_wap")
	s.RefundReason = field.NewString(table, "refund_reason")
	s.RefundReasonTime = field.NewInt64(table, "refund_reason_time")
	s.RefundPrice = field.NewString(table, "refund_price")
	s.DeliveryName = field.NewString(table, "delivery_name")
	s.DeliveryType = field.NewString(table, "delivery_type")
	s.DeliveryID = field.NewString(table, "delivery_id")
	s.GainIntegral = field.NewInt64(table, "gain_integral")
	s.UseIntegral = field.NewInt64(table, "use_integral")
	s.BackIntegral = field.NewInt64(table, "back_integral")
	s.Mark = field.NewString(table, "mark")
	s.IsDel = field.NewInt64(table, "is_del")
	s.Remark = field.NewString(table, "remark")
	s.MerID = field.NewInt64(table, "mer_id")
	s.IsMerCheck = field.NewInt64(table, "is_mer_check")
	s.CombinationID = field.NewInt64(table, "combination_id")
	s.PinkID = field.NewInt64(table, "pink_id")
	s.Cost = field.NewString(table, "cost")
	s.SeckillID = field.NewInt64(table, "seckill_id")
	s.BargainID = field.NewInt64(table, "bargain_id")
	s.VerifyCode = field.NewString(table, "verify_code")
	s.StoreID = field.NewInt64(table, "store_id")
	s.ShippingType = field.NewInt64(table, "shipping_type")
	s.ClerkID = field.NewInt64(table, "clerk_id")
	s.IsChannel = field.NewInt64(table, "is_channel")
	s.IsRemind = field.NewInt64(table, "is_remind")
	s.IsSystemDel = field.NewInt64(table, "is_system_del")
	s.DeliveryCode = field.NewString(table, "delivery_code")
	s.BargainUserID = field.NewInt64(table, "bargain_user_id")
	s.Type = field.NewInt64(table, "type")
	s.ProTotalPrice = field.NewString(table, "pro_total_price")
	s.BeforePayPrice = field.NewString(table, "before_pay_price")
	s.IsAlterPrice = field.NewInt64(table, "is_alter_price")
	s.OutTradeNo = field.NewString(table, "out_trade_no")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeOrder) WithContext(ctx context.Context) IStoreOrderDo {
	return s.storeOrderDo.WithContext(ctx)
}

func (s storeOrder) TableName() string { return s.storeOrderDo.TableName() }

func (s storeOrder) Alias() string { return s.storeOrderDo.Alias() }

func (s storeOrder) Columns(cols ...field.Expr) gen.Columns { return s.storeOrderDo.Columns(cols...) }

func (s *storeOrder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeOrder) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 59)
	s.fieldMap["id"] = s.ID
	s.fieldMap["order_id"] = s.OrderID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["real_name"] = s.RealName
	s.fieldMap["user_phone"] = s.UserPhone
	s.fieldMap["user_address"] = s.UserAddress
	s.fieldMap["freight_price"] = s.FreightPrice
	s.fieldMap["total_num"] = s.TotalNum
	s.fieldMap["total_price"] = s.TotalPrice
	s.fieldMap["total_postage"] = s.TotalPostage
	s.fieldMap["pay_price"] = s.PayPrice
	s.fieldMap["pay_postage"] = s.PayPostage
	s.fieldMap["deduction_price"] = s.DeductionPrice
	s.fieldMap["coupon_id"] = s.CouponID
	s.fieldMap["coupon_price"] = s.CouponPrice
	s.fieldMap["paid"] = s.Paid
	s.fieldMap["pay_time"] = s.PayTime
	s.fieldMap["pay_type"] = s.PayType
	s.fieldMap["status"] = s.Status
	s.fieldMap["refund_status"] = s.RefundStatus
	s.fieldMap["refund_reason_wap_img"] = s.RefundReasonWapImg
	s.fieldMap["refund_reason_wap_explain"] = s.RefundReasonWapExplain
	s.fieldMap["refund_reason_wap"] = s.RefundReasonWap
	s.fieldMap["refund_reason"] = s.RefundReason
	s.fieldMap["refund_reason_time"] = s.RefundReasonTime
	s.fieldMap["refund_price"] = s.RefundPrice
	s.fieldMap["delivery_name"] = s.DeliveryName
	s.fieldMap["delivery_type"] = s.DeliveryType
	s.fieldMap["delivery_id"] = s.DeliveryID
	s.fieldMap["gain_integral"] = s.GainIntegral
	s.fieldMap["use_integral"] = s.UseIntegral
	s.fieldMap["back_integral"] = s.BackIntegral
	s.fieldMap["mark"] = s.Mark
	s.fieldMap["is_del"] = s.IsDel
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["mer_id"] = s.MerID
	s.fieldMap["is_mer_check"] = s.IsMerCheck
	s.fieldMap["combination_id"] = s.CombinationID
	s.fieldMap["pink_id"] = s.PinkID
	s.fieldMap["cost"] = s.Cost
	s.fieldMap["seckill_id"] = s.SeckillID
	s.fieldMap["bargain_id"] = s.BargainID
	s.fieldMap["verify_code"] = s.VerifyCode
	s.fieldMap["store_id"] = s.StoreID
	s.fieldMap["shipping_type"] = s.ShippingType
	s.fieldMap["clerk_id"] = s.ClerkID
	s.fieldMap["is_channel"] = s.IsChannel
	s.fieldMap["is_remind"] = s.IsRemind
	s.fieldMap["is_system_del"] = s.IsSystemDel
	s.fieldMap["delivery_code"] = s.DeliveryCode
	s.fieldMap["bargain_user_id"] = s.BargainUserID
	s.fieldMap["type"] = s.Type
	s.fieldMap["pro_total_price"] = s.ProTotalPrice
	s.fieldMap["before_pay_price"] = s.BeforePayPrice
	s.fieldMap["is_alter_price"] = s.IsAlterPrice
	s.fieldMap["out_trade_no"] = s.OutTradeNo
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeOrder) clone(db *gorm.DB) storeOrder {
	s.storeOrderDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeOrder) replaceDB(db *gorm.DB) storeOrder {
	s.storeOrderDo.ReplaceDB(db)
	return s
}

type storeOrderDo struct{ gen.DO }

type IStoreOrderDo interface {
	gen.SubQuery
	Debug() IStoreOrderDo
	WithContext(ctx context.Context) IStoreOrderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreOrderDo
	WriteDB() IStoreOrderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreOrderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreOrderDo
	Not(conds ...gen.Condition) IStoreOrderDo
	Or(conds ...gen.Condition) IStoreOrderDo
	Select(conds ...field.Expr) IStoreOrderDo
	Where(conds ...gen.Condition) IStoreOrderDo
	Order(conds ...field.Expr) IStoreOrderDo
	Distinct(cols ...field.Expr) IStoreOrderDo
	Omit(cols ...field.Expr) IStoreOrderDo
	Join(table schema.Tabler, on ...field.Expr) IStoreOrderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreOrderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreOrderDo
	Group(cols ...field.Expr) IStoreOrderDo
	Having(conds ...gen.Condition) IStoreOrderDo
	Limit(limit int) IStoreOrderDo
	Offset(offset int) IStoreOrderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreOrderDo
	Unscoped() IStoreOrderDo
	Create(values ...*model.StoreOrder) error
	CreateInBatches(values []*model.StoreOrder, batchSize int) error
	Save(values ...*model.StoreOrder) error
	First() (*model.StoreOrder, error)
	Take() (*model.StoreOrder, error)
	Last() (*model.StoreOrder, error)
	Find() ([]*model.StoreOrder, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreOrder, err error)
	FindInBatches(result *[]*model.StoreOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreOrder) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreOrderDo
	Assign(attrs ...field.AssignExpr) IStoreOrderDo
	Joins(fields ...field.RelationField) IStoreOrderDo
	Preload(fields ...field.RelationField) IStoreOrderDo
	FirstOrInit() (*model.StoreOrder, error)
	FirstOrCreate() (*model.StoreOrder, error)
	FindByPage(offset int, limit int) (result []*model.StoreOrder, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreOrderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeOrderDo) Debug() IStoreOrderDo {
	return s.withDO(s.DO.Debug())
}

func (s storeOrderDo) WithContext(ctx context.Context) IStoreOrderDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeOrderDo) ReadDB() IStoreOrderDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeOrderDo) WriteDB() IStoreOrderDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeOrderDo) Session(config *gorm.Session) IStoreOrderDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeOrderDo) Clauses(conds ...clause.Expression) IStoreOrderDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeOrderDo) Returning(value interface{}, columns ...string) IStoreOrderDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeOrderDo) Not(conds ...gen.Condition) IStoreOrderDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeOrderDo) Or(conds ...gen.Condition) IStoreOrderDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeOrderDo) Select(conds ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeOrderDo) Where(conds ...gen.Condition) IStoreOrderDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeOrderDo) Order(conds ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeOrderDo) Distinct(cols ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeOrderDo) Omit(cols ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeOrderDo) Join(table schema.Tabler, on ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeOrderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeOrderDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeOrderDo) Group(cols ...field.Expr) IStoreOrderDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeOrderDo) Having(conds ...gen.Condition) IStoreOrderDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeOrderDo) Limit(limit int) IStoreOrderDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeOrderDo) Offset(offset int) IStoreOrderDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeOrderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreOrderDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeOrderDo) Unscoped() IStoreOrderDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeOrderDo) Create(values ...*model.StoreOrder) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeOrderDo) CreateInBatches(values []*model.StoreOrder, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeOrderDo) Save(values ...*model.StoreOrder) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeOrderDo) First() (*model.StoreOrder, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreOrder), nil
	}
}

func (s storeOrderDo) Take() (*model.StoreOrder, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreOrder), nil
	}
}

func (s storeOrderDo) Last() (*model.StoreOrder, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreOrder), nil
	}
}

func (s storeOrderDo) Find() ([]*model.StoreOrder, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreOrder), err
}

func (s storeOrderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreOrder, err error) {
	buf := make([]*model.StoreOrder, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeOrderDo) FindInBatches(result *[]*model.StoreOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeOrderDo) Attrs(attrs ...field.AssignExpr) IStoreOrderDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeOrderDo) Assign(attrs ...field.AssignExpr) IStoreOrderDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeOrderDo) Joins(fields ...field.RelationField) IStoreOrderDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeOrderDo) Preload(fields ...field.RelationField) IStoreOrderDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeOrderDo) FirstOrInit() (*model.StoreOrder, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreOrder), nil
	}
}

func (s storeOrderDo) FirstOrCreate() (*model.StoreOrder, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreOrder), nil
	}
}

func (s storeOrderDo) FindByPage(offset int, limit int) (result []*model.StoreOrder, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s storeOrderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeOrderDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeOrderDo) Delete(models ...*model.StoreOrder) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeOrderDo) withDO(do gen.Dao) *storeOrderDo {
	s.DO = *do.(*gen.DO)
	return s
}
