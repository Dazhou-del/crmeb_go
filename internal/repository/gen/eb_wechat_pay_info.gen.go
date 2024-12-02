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

func newEbWechatPayInfo(db *gorm.DB, opts ...gen.DOOption) ebWechatPayInfo {
	_ebWechatPayInfo := ebWechatPayInfo{}

	_ebWechatPayInfo.ebWechatPayInfoDo.UseDB(db, opts...)
	_ebWechatPayInfo.ebWechatPayInfoDo.UseModel(&model.EbWechatPayInfo{})

	tableName := _ebWechatPayInfo.ebWechatPayInfoDo.TableName()
	_ebWechatPayInfo.ALL = field.NewAsterisk(tableName)
	_ebWechatPayInfo.ID = field.NewInt32(tableName, "id")
	_ebWechatPayInfo.AppID = field.NewString(tableName, "app_id")
	_ebWechatPayInfo.MchID = field.NewString(tableName, "mch_id")
	_ebWechatPayInfo.DeviceInfo = field.NewString(tableName, "device_info")
	_ebWechatPayInfo.OpenID = field.NewString(tableName, "open_id")
	_ebWechatPayInfo.NonceStr = field.NewString(tableName, "nonce_str")
	_ebWechatPayInfo.Sign = field.NewString(tableName, "sign")
	_ebWechatPayInfo.SignType = field.NewString(tableName, "sign_type")
	_ebWechatPayInfo.Body = field.NewString(tableName, "body")
	_ebWechatPayInfo.Detail = field.NewString(tableName, "detail")
	_ebWechatPayInfo.Attach = field.NewString(tableName, "attach")
	_ebWechatPayInfo.OutTradeNo = field.NewString(tableName, "out_trade_no")
	_ebWechatPayInfo.FeeType = field.NewString(tableName, "fee_type")
	_ebWechatPayInfo.TotalFee = field.NewInt32(tableName, "total_fee")
	_ebWechatPayInfo.SpbillCreateIP = field.NewString(tableName, "spbill_create_ip")
	_ebWechatPayInfo.TimeStart = field.NewString(tableName, "time_start")
	_ebWechatPayInfo.TimeExpire = field.NewString(tableName, "time_expire")
	_ebWechatPayInfo.NotifyURL = field.NewString(tableName, "notify_url")
	_ebWechatPayInfo.TradeType = field.NewString(tableName, "trade_type")
	_ebWechatPayInfo.ProductID = field.NewString(tableName, "product_id")
	_ebWechatPayInfo.SceneInfo = field.NewString(tableName, "scene_info")
	_ebWechatPayInfo.ErrCode = field.NewString(tableName, "err_code")
	_ebWechatPayInfo.PrepayID = field.NewString(tableName, "prepay_id")
	_ebWechatPayInfo.CodeURL = field.NewString(tableName, "code_url")
	_ebWechatPayInfo.IsSubscribe = field.NewString(tableName, "is_subscribe")
	_ebWechatPayInfo.TradeState = field.NewString(tableName, "trade_state")
	_ebWechatPayInfo.BankType = field.NewString(tableName, "bank_type")
	_ebWechatPayInfo.CashFee = field.NewInt32(tableName, "cash_fee")
	_ebWechatPayInfo.CouponFee = field.NewInt32(tableName, "coupon_fee")
	_ebWechatPayInfo.TransactionID = field.NewString(tableName, "transaction_id")
	_ebWechatPayInfo.TimeEnd = field.NewString(tableName, "time_end")
	_ebWechatPayInfo.TradeStateDesc = field.NewString(tableName, "trade_state_desc")

	_ebWechatPayInfo.fillFieldMap()

	return _ebWechatPayInfo
}

// ebWechatPayInfo 微信订单表
type ebWechatPayInfo struct {
	ebWechatPayInfoDo ebWechatPayInfoDo

	ALL            field.Asterisk
	ID             field.Int32
	AppID          field.String // 公众号唯一标识
	MchID          field.String // 商户号
	DeviceInfo     field.String // 设备号,PC网页或公众号内支付可以传-WEB
	OpenID         field.String // 用户的唯一标识
	NonceStr       field.String // 随机字符串
	Sign           field.String // 签名
	SignType       field.String // 签名类型，默认为MD5，支持HMAC-SHA256和MD5
	Body           field.String // 商品描述
	Detail         field.String // 商品详细描述，对于使用单品优惠的商户，该字段必须按照规范上传
	Attach         field.String // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	OutTradeNo     field.String // 商户订单号,要求32个字符内
	FeeType        field.String // 标价币种：CNY：人民币 GBP：英镑 HKD：港币 USD：美元 JPY：日元 CAD：加拿大元 AUD：澳大利亚元 EUR：欧元 NZD：新西兰元 KRW：韩元 THB：泰铢
	TotalFee       field.Int32  // 标价金额
	SpbillCreateIP field.String // 终端IP
	TimeStart      field.String // 交易起始时间
	TimeExpire     field.String // 交易结束时间
	NotifyURL      field.String // 通知地址
	TradeType      field.String // 交易类型,取值为：JSAPI，NATIVE，APP等
	ProductID      field.String // 商品ID
	SceneInfo      field.String // 场景信息
	ErrCode        field.String // 错误代码
	PrepayID       field.String // 预支付交易会话标识
	CodeURL        field.String // 二维码链接
	IsSubscribe    field.String // 是否关注公众账号
	TradeState     field.String // 交易状态
	BankType       field.String // 付款银行
	CashFee        field.Int32  // 现金支付金额
	CouponFee      field.Int32  // 代金券金额
	TransactionID  field.String // 微信支付订单号
	TimeEnd        field.String // 支付完成时间
	TradeStateDesc field.String // 交易状态描述

	fieldMap map[string]field.Expr
}

func (e ebWechatPayInfo) Table(newTableName string) *ebWechatPayInfo {
	e.ebWechatPayInfoDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebWechatPayInfo) As(alias string) *ebWechatPayInfo {
	e.ebWechatPayInfoDo.DO = *(e.ebWechatPayInfoDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebWechatPayInfo) updateTableName(table string) *ebWechatPayInfo {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.AppID = field.NewString(table, "app_id")
	e.MchID = field.NewString(table, "mch_id")
	e.DeviceInfo = field.NewString(table, "device_info")
	e.OpenID = field.NewString(table, "open_id")
	e.NonceStr = field.NewString(table, "nonce_str")
	e.Sign = field.NewString(table, "sign")
	e.SignType = field.NewString(table, "sign_type")
	e.Body = field.NewString(table, "body")
	e.Detail = field.NewString(table, "detail")
	e.Attach = field.NewString(table, "attach")
	e.OutTradeNo = field.NewString(table, "out_trade_no")
	e.FeeType = field.NewString(table, "fee_type")
	e.TotalFee = field.NewInt32(table, "total_fee")
	e.SpbillCreateIP = field.NewString(table, "spbill_create_ip")
	e.TimeStart = field.NewString(table, "time_start")
	e.TimeExpire = field.NewString(table, "time_expire")
	e.NotifyURL = field.NewString(table, "notify_url")
	e.TradeType = field.NewString(table, "trade_type")
	e.ProductID = field.NewString(table, "product_id")
	e.SceneInfo = field.NewString(table, "scene_info")
	e.ErrCode = field.NewString(table, "err_code")
	e.PrepayID = field.NewString(table, "prepay_id")
	e.CodeURL = field.NewString(table, "code_url")
	e.IsSubscribe = field.NewString(table, "is_subscribe")
	e.TradeState = field.NewString(table, "trade_state")
	e.BankType = field.NewString(table, "bank_type")
	e.CashFee = field.NewInt32(table, "cash_fee")
	e.CouponFee = field.NewInt32(table, "coupon_fee")
	e.TransactionID = field.NewString(table, "transaction_id")
	e.TimeEnd = field.NewString(table, "time_end")
	e.TradeStateDesc = field.NewString(table, "trade_state_desc")

	e.fillFieldMap()

	return e
}

func (e *ebWechatPayInfo) WithContext(ctx context.Context) IEbWechatPayInfoDo {
	return e.ebWechatPayInfoDo.WithContext(ctx)
}

func (e ebWechatPayInfo) TableName() string { return e.ebWechatPayInfoDo.TableName() }

func (e ebWechatPayInfo) Alias() string { return e.ebWechatPayInfoDo.Alias() }

func (e ebWechatPayInfo) Columns(cols ...field.Expr) gen.Columns {
	return e.ebWechatPayInfoDo.Columns(cols...)
}

func (e *ebWechatPayInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebWechatPayInfo) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 32)
	e.fieldMap["id"] = e.ID
	e.fieldMap["app_id"] = e.AppID
	e.fieldMap["mch_id"] = e.MchID
	e.fieldMap["device_info"] = e.DeviceInfo
	e.fieldMap["open_id"] = e.OpenID
	e.fieldMap["nonce_str"] = e.NonceStr
	e.fieldMap["sign"] = e.Sign
	e.fieldMap["sign_type"] = e.SignType
	e.fieldMap["body"] = e.Body
	e.fieldMap["detail"] = e.Detail
	e.fieldMap["attach"] = e.Attach
	e.fieldMap["out_trade_no"] = e.OutTradeNo
	e.fieldMap["fee_type"] = e.FeeType
	e.fieldMap["total_fee"] = e.TotalFee
	e.fieldMap["spbill_create_ip"] = e.SpbillCreateIP
	e.fieldMap["time_start"] = e.TimeStart
	e.fieldMap["time_expire"] = e.TimeExpire
	e.fieldMap["notify_url"] = e.NotifyURL
	e.fieldMap["trade_type"] = e.TradeType
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["scene_info"] = e.SceneInfo
	e.fieldMap["err_code"] = e.ErrCode
	e.fieldMap["prepay_id"] = e.PrepayID
	e.fieldMap["code_url"] = e.CodeURL
	e.fieldMap["is_subscribe"] = e.IsSubscribe
	e.fieldMap["trade_state"] = e.TradeState
	e.fieldMap["bank_type"] = e.BankType
	e.fieldMap["cash_fee"] = e.CashFee
	e.fieldMap["coupon_fee"] = e.CouponFee
	e.fieldMap["transaction_id"] = e.TransactionID
	e.fieldMap["time_end"] = e.TimeEnd
	e.fieldMap["trade_state_desc"] = e.TradeStateDesc
}

func (e ebWechatPayInfo) clone(db *gorm.DB) ebWechatPayInfo {
	e.ebWechatPayInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebWechatPayInfo) replaceDB(db *gorm.DB) ebWechatPayInfo {
	e.ebWechatPayInfoDo.ReplaceDB(db)
	return e
}

type ebWechatPayInfoDo struct{ gen.DO }

type IEbWechatPayInfoDo interface {
	gen.SubQuery
	Debug() IEbWechatPayInfoDo
	WithContext(ctx context.Context) IEbWechatPayInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbWechatPayInfoDo
	WriteDB() IEbWechatPayInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbWechatPayInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbWechatPayInfoDo
	Not(conds ...gen.Condition) IEbWechatPayInfoDo
	Or(conds ...gen.Condition) IEbWechatPayInfoDo
	Select(conds ...field.Expr) IEbWechatPayInfoDo
	Where(conds ...gen.Condition) IEbWechatPayInfoDo
	Order(conds ...field.Expr) IEbWechatPayInfoDo
	Distinct(cols ...field.Expr) IEbWechatPayInfoDo
	Omit(cols ...field.Expr) IEbWechatPayInfoDo
	Join(table schema.Tabler, on ...field.Expr) IEbWechatPayInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbWechatPayInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbWechatPayInfoDo
	Group(cols ...field.Expr) IEbWechatPayInfoDo
	Having(conds ...gen.Condition) IEbWechatPayInfoDo
	Limit(limit int) IEbWechatPayInfoDo
	Offset(offset int) IEbWechatPayInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbWechatPayInfoDo
	Unscoped() IEbWechatPayInfoDo
	Create(values ...*model.EbWechatPayInfo) error
	CreateInBatches(values []*model.EbWechatPayInfo, batchSize int) error
	Save(values ...*model.EbWechatPayInfo) error
	First() (*model.EbWechatPayInfo, error)
	Take() (*model.EbWechatPayInfo, error)
	Last() (*model.EbWechatPayInfo, error)
	Find() ([]*model.EbWechatPayInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbWechatPayInfo, err error)
	FindInBatches(result *[]*model.EbWechatPayInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbWechatPayInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbWechatPayInfoDo
	Assign(attrs ...field.AssignExpr) IEbWechatPayInfoDo
	Joins(fields ...field.RelationField) IEbWechatPayInfoDo
	Preload(fields ...field.RelationField) IEbWechatPayInfoDo
	FirstOrInit() (*model.EbWechatPayInfo, error)
	FirstOrCreate() (*model.EbWechatPayInfo, error)
	FindByPage(offset int, limit int) (result []*model.EbWechatPayInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbWechatPayInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebWechatPayInfoDo) Debug() IEbWechatPayInfoDo {
	return e.withDO(e.DO.Debug())
}

func (e ebWechatPayInfoDo) WithContext(ctx context.Context) IEbWechatPayInfoDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebWechatPayInfoDo) ReadDB() IEbWechatPayInfoDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebWechatPayInfoDo) WriteDB() IEbWechatPayInfoDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebWechatPayInfoDo) Session(config *gorm.Session) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebWechatPayInfoDo) Clauses(conds ...clause.Expression) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebWechatPayInfoDo) Returning(value interface{}, columns ...string) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebWechatPayInfoDo) Not(conds ...gen.Condition) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebWechatPayInfoDo) Or(conds ...gen.Condition) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebWechatPayInfoDo) Select(conds ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebWechatPayInfoDo) Where(conds ...gen.Condition) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebWechatPayInfoDo) Order(conds ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebWechatPayInfoDo) Distinct(cols ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebWechatPayInfoDo) Omit(cols ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebWechatPayInfoDo) Join(table schema.Tabler, on ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebWechatPayInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebWechatPayInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebWechatPayInfoDo) Group(cols ...field.Expr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebWechatPayInfoDo) Having(conds ...gen.Condition) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebWechatPayInfoDo) Limit(limit int) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebWechatPayInfoDo) Offset(offset int) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebWechatPayInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebWechatPayInfoDo) Unscoped() IEbWechatPayInfoDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebWechatPayInfoDo) Create(values ...*model.EbWechatPayInfo) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebWechatPayInfoDo) CreateInBatches(values []*model.EbWechatPayInfo, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebWechatPayInfoDo) Save(values ...*model.EbWechatPayInfo) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebWechatPayInfoDo) First() (*model.EbWechatPayInfo, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatPayInfo), nil
	}
}

func (e ebWechatPayInfoDo) Take() (*model.EbWechatPayInfo, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatPayInfo), nil
	}
}

func (e ebWechatPayInfoDo) Last() (*model.EbWechatPayInfo, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatPayInfo), nil
	}
}

func (e ebWechatPayInfoDo) Find() ([]*model.EbWechatPayInfo, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbWechatPayInfo), err
}

func (e ebWechatPayInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbWechatPayInfo, err error) {
	buf := make([]*model.EbWechatPayInfo, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebWechatPayInfoDo) FindInBatches(result *[]*model.EbWechatPayInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebWechatPayInfoDo) Attrs(attrs ...field.AssignExpr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebWechatPayInfoDo) Assign(attrs ...field.AssignExpr) IEbWechatPayInfoDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebWechatPayInfoDo) Joins(fields ...field.RelationField) IEbWechatPayInfoDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebWechatPayInfoDo) Preload(fields ...field.RelationField) IEbWechatPayInfoDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebWechatPayInfoDo) FirstOrInit() (*model.EbWechatPayInfo, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatPayInfo), nil
	}
}

func (e ebWechatPayInfoDo) FirstOrCreate() (*model.EbWechatPayInfo, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatPayInfo), nil
	}
}

func (e ebWechatPayInfoDo) FindByPage(offset int, limit int) (result []*model.EbWechatPayInfo, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e ebWechatPayInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebWechatPayInfoDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebWechatPayInfoDo) Delete(models ...*model.EbWechatPayInfo) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebWechatPayInfoDo) withDO(do gen.Dao) *ebWechatPayInfoDo {
	e.DO = *do.(*gen.DO)
	return e
}
