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

func newUserExtract(db *gorm.DB, opts ...gen.DOOption) userExtract {
	_userExtract := userExtract{}

	_userExtract.userExtractDo.UseDB(db, opts...)
	_userExtract.userExtractDo.UseModel(&model.UserExtract{})

	tableName := _userExtract.userExtractDo.TableName()
	_userExtract.ALL = field.NewAsterisk(tableName)
	_userExtract.ID = field.NewInt64(tableName, "id")
	_userExtract.UID = field.NewInt64(tableName, "uid")
	_userExtract.RealName = field.NewString(tableName, "real_name")
	_userExtract.ExtractType = field.NewString(tableName, "extract_type")
	_userExtract.BankCode = field.NewString(tableName, "bank_code")
	_userExtract.BankAddress = field.NewString(tableName, "bank_address")
	_userExtract.AlipayCode = field.NewString(tableName, "alipay_code")
	_userExtract.ExtractPrice = field.NewField(tableName, "extract_price")
	_userExtract.Mark = field.NewString(tableName, "mark")
	_userExtract.Balance = field.NewField(tableName, "balance")
	_userExtract.FailMsg = field.NewString(tableName, "fail_msg")
	_userExtract.Status = field.NewInt64(tableName, "status")
	_userExtract.Wechat = field.NewString(tableName, "wechat")
	_userExtract.BankName = field.NewString(tableName, "bank_name")
	_userExtract.QrcodeURL = field.NewString(tableName, "qrcode_url")
	_userExtract.CreatedAt = field.NewInt64(tableName, "created_at")
	_userExtract.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_userExtract.DeletedAt = field.NewField(tableName, "deleted_at")

	_userExtract.fillFieldMap()

	return _userExtract
}

// userExtract 用户提现表
type userExtract struct {
	userExtractDo userExtractDo

	ALL          field.Asterisk
	ID           field.Int64
	UID          field.Int64
	RealName     field.String // 名称
	ExtractType  field.String // bank = 银行卡 alipay = 支付宝 weixin=微信
	BankCode     field.String // 银行卡
	BankAddress  field.String // 开户地址
	AlipayCode   field.String // 支付宝账号
	ExtractPrice field.Field  // 提现金额
	Mark         field.String
	Balance      field.Field
	FailMsg      field.String // 无效原因
	Status       field.Int64  // -1 未通过 0 审核中 1 已提现
	Wechat       field.String // 微信号
	BankName     field.String // 银行名称
	QrcodeURL    field.String // 微信收款二维码
	CreatedAt    field.Int64
	UpdatedAt    field.Int64
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (u userExtract) Table(newTableName string) *userExtract {
	u.userExtractDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userExtract) As(alias string) *userExtract {
	u.userExtractDo.DO = *(u.userExtractDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userExtract) updateTableName(table string) *userExtract {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.RealName = field.NewString(table, "real_name")
	u.ExtractType = field.NewString(table, "extract_type")
	u.BankCode = field.NewString(table, "bank_code")
	u.BankAddress = field.NewString(table, "bank_address")
	u.AlipayCode = field.NewString(table, "alipay_code")
	u.ExtractPrice = field.NewField(table, "extract_price")
	u.Mark = field.NewString(table, "mark")
	u.Balance = field.NewField(table, "balance")
	u.FailMsg = field.NewString(table, "fail_msg")
	u.Status = field.NewInt64(table, "status")
	u.Wechat = field.NewString(table, "wechat")
	u.BankName = field.NewString(table, "bank_name")
	u.QrcodeURL = field.NewString(table, "qrcode_url")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userExtract) WithContext(ctx context.Context) IUserExtractDo {
	return u.userExtractDo.WithContext(ctx)
}

func (u userExtract) TableName() string { return u.userExtractDo.TableName() }

func (u userExtract) Alias() string { return u.userExtractDo.Alias() }

func (u userExtract) Columns(cols ...field.Expr) gen.Columns { return u.userExtractDo.Columns(cols...) }

func (u *userExtract) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userExtract) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 18)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["real_name"] = u.RealName
	u.fieldMap["extract_type"] = u.ExtractType
	u.fieldMap["bank_code"] = u.BankCode
	u.fieldMap["bank_address"] = u.BankAddress
	u.fieldMap["alipay_code"] = u.AlipayCode
	u.fieldMap["extract_price"] = u.ExtractPrice
	u.fieldMap["mark"] = u.Mark
	u.fieldMap["balance"] = u.Balance
	u.fieldMap["fail_msg"] = u.FailMsg
	u.fieldMap["status"] = u.Status
	u.fieldMap["wechat"] = u.Wechat
	u.fieldMap["bank_name"] = u.BankName
	u.fieldMap["qrcode_url"] = u.QrcodeURL
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userExtract) clone(db *gorm.DB) userExtract {
	u.userExtractDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userExtract) replaceDB(db *gorm.DB) userExtract {
	u.userExtractDo.ReplaceDB(db)
	return u
}

type userExtractDo struct{ gen.DO }

type IUserExtractDo interface {
	gen.SubQuery
	Debug() IUserExtractDo
	WithContext(ctx context.Context) IUserExtractDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserExtractDo
	WriteDB() IUserExtractDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserExtractDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserExtractDo
	Not(conds ...gen.Condition) IUserExtractDo
	Or(conds ...gen.Condition) IUserExtractDo
	Select(conds ...field.Expr) IUserExtractDo
	Where(conds ...gen.Condition) IUserExtractDo
	Order(conds ...field.Expr) IUserExtractDo
	Distinct(cols ...field.Expr) IUserExtractDo
	Omit(cols ...field.Expr) IUserExtractDo
	Join(table schema.Tabler, on ...field.Expr) IUserExtractDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserExtractDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserExtractDo
	Group(cols ...field.Expr) IUserExtractDo
	Having(conds ...gen.Condition) IUserExtractDo
	Limit(limit int) IUserExtractDo
	Offset(offset int) IUserExtractDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserExtractDo
	Unscoped() IUserExtractDo
	Create(values ...*model.UserExtract) error
	CreateInBatches(values []*model.UserExtract, batchSize int) error
	Save(values ...*model.UserExtract) error
	First() (*model.UserExtract, error)
	Take() (*model.UserExtract, error)
	Last() (*model.UserExtract, error)
	Find() ([]*model.UserExtract, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserExtract, err error)
	FindInBatches(result *[]*model.UserExtract, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserExtract) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserExtractDo
	Assign(attrs ...field.AssignExpr) IUserExtractDo
	Joins(fields ...field.RelationField) IUserExtractDo
	Preload(fields ...field.RelationField) IUserExtractDo
	FirstOrInit() (*model.UserExtract, error)
	FirstOrCreate() (*model.UserExtract, error)
	FindByPage(offset int, limit int) (result []*model.UserExtract, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserExtractDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userExtractDo) Debug() IUserExtractDo {
	return u.withDO(u.DO.Debug())
}

func (u userExtractDo) WithContext(ctx context.Context) IUserExtractDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userExtractDo) ReadDB() IUserExtractDo {
	return u.Clauses(dbresolver.Read)
}

func (u userExtractDo) WriteDB() IUserExtractDo {
	return u.Clauses(dbresolver.Write)
}

func (u userExtractDo) Session(config *gorm.Session) IUserExtractDo {
	return u.withDO(u.DO.Session(config))
}

func (u userExtractDo) Clauses(conds ...clause.Expression) IUserExtractDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userExtractDo) Returning(value interface{}, columns ...string) IUserExtractDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userExtractDo) Not(conds ...gen.Condition) IUserExtractDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userExtractDo) Or(conds ...gen.Condition) IUserExtractDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userExtractDo) Select(conds ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userExtractDo) Where(conds ...gen.Condition) IUserExtractDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userExtractDo) Order(conds ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userExtractDo) Distinct(cols ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userExtractDo) Omit(cols ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userExtractDo) Join(table schema.Tabler, on ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userExtractDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userExtractDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userExtractDo) Group(cols ...field.Expr) IUserExtractDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userExtractDo) Having(conds ...gen.Condition) IUserExtractDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userExtractDo) Limit(limit int) IUserExtractDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userExtractDo) Offset(offset int) IUserExtractDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userExtractDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserExtractDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userExtractDo) Unscoped() IUserExtractDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userExtractDo) Create(values ...*model.UserExtract) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userExtractDo) CreateInBatches(values []*model.UserExtract, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userExtractDo) Save(values ...*model.UserExtract) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userExtractDo) First() (*model.UserExtract, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExtract), nil
	}
}

func (u userExtractDo) Take() (*model.UserExtract, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExtract), nil
	}
}

func (u userExtractDo) Last() (*model.UserExtract, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExtract), nil
	}
}

func (u userExtractDo) Find() ([]*model.UserExtract, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserExtract), err
}

func (u userExtractDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserExtract, err error) {
	buf := make([]*model.UserExtract, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userExtractDo) FindInBatches(result *[]*model.UserExtract, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userExtractDo) Attrs(attrs ...field.AssignExpr) IUserExtractDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userExtractDo) Assign(attrs ...field.AssignExpr) IUserExtractDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userExtractDo) Joins(fields ...field.RelationField) IUserExtractDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userExtractDo) Preload(fields ...field.RelationField) IUserExtractDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userExtractDo) FirstOrInit() (*model.UserExtract, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExtract), nil
	}
}

func (u userExtractDo) FirstOrCreate() (*model.UserExtract, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExtract), nil
	}
}

func (u userExtractDo) FindByPage(offset int, limit int) (result []*model.UserExtract, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userExtractDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userExtractDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userExtractDo) Delete(models ...*model.UserExtract) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userExtractDo) withDO(do gen.Dao) *userExtractDo {
	u.DO = *do.(*gen.DO)
	return u
}