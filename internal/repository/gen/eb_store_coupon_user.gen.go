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

func newEbStoreCouponUser(db *gorm.DB, opts ...gen.DOOption) ebStoreCouponUser {
	_ebStoreCouponUser := ebStoreCouponUser{}

	_ebStoreCouponUser.ebStoreCouponUserDo.UseDB(db, opts...)
	_ebStoreCouponUser.ebStoreCouponUserDo.UseModel(&model.EbStoreCouponUser{})

	tableName := _ebStoreCouponUser.ebStoreCouponUserDo.TableName()
	_ebStoreCouponUser.ALL = field.NewAsterisk(tableName)
	_ebStoreCouponUser.ID = field.NewInt32(tableName, "id")
	_ebStoreCouponUser.CouponID = field.NewInt32(tableName, "coupon_id")
	_ebStoreCouponUser.Cid = field.NewInt32(tableName, "cid")
	_ebStoreCouponUser.UID = field.NewInt32(tableName, "uid")
	_ebStoreCouponUser.Name = field.NewString(tableName, "name")
	_ebStoreCouponUser.Money = field.NewFloat64(tableName, "money")
	_ebStoreCouponUser.MinPrice = field.NewFloat64(tableName, "min_price")
	_ebStoreCouponUser.Type = field.NewString(tableName, "type")
	_ebStoreCouponUser.Status = field.NewBool(tableName, "status")
	_ebStoreCouponUser.CreateTime = field.NewTime(tableName, "create_time")
	_ebStoreCouponUser.UpdateTime = field.NewTime(tableName, "update_time")
	_ebStoreCouponUser.StartTime = field.NewTime(tableName, "start_time")
	_ebStoreCouponUser.EndTime = field.NewTime(tableName, "end_time")
	_ebStoreCouponUser.UseTime = field.NewTime(tableName, "use_time")
	_ebStoreCouponUser.UseType = field.NewBool(tableName, "use_type")
	_ebStoreCouponUser.PrimaryKey = field.NewString(tableName, "primary_key")

	_ebStoreCouponUser.fillFieldMap()

	return _ebStoreCouponUser
}

// ebStoreCouponUser 优惠券记录表
type ebStoreCouponUser struct {
	ebStoreCouponUserDo ebStoreCouponUserDo

	ALL        field.Asterisk
	ID         field.Int32   // id
	CouponID   field.Int32   // 优惠券发布id
	Cid        field.Int32   // 兑换的项目id
	UID        field.Int32   // 领取人id
	Name       field.String  // 优惠券名称
	Money      field.Float64 // 优惠券的面值
	MinPrice   field.Float64 // 最低消费多少金额可用优惠券
	Type       field.String  // 获取方式，send后台发放, 用户领取 get
	Status     field.Bool    // 状态（0：未使用，1：已使用, 2:已失效）
	CreateTime field.Time    // 创建时间
	UpdateTime field.Time    // 更新时间
	StartTime  field.Time    // 开始使用时间
	EndTime    field.Time    // 过期时间
	UseTime    field.Time    // 使用时间
	UseType    field.Bool    // 使用类型 1 全场通用, 2 商品券, 3 品类券
	PrimaryKey field.String  // 所属商品id / 分类id

	fieldMap map[string]field.Expr
}

func (e ebStoreCouponUser) Table(newTableName string) *ebStoreCouponUser {
	e.ebStoreCouponUserDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreCouponUser) As(alias string) *ebStoreCouponUser {
	e.ebStoreCouponUserDo.DO = *(e.ebStoreCouponUserDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreCouponUser) updateTableName(table string) *ebStoreCouponUser {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.CouponID = field.NewInt32(table, "coupon_id")
	e.Cid = field.NewInt32(table, "cid")
	e.UID = field.NewInt32(table, "uid")
	e.Name = field.NewString(table, "name")
	e.Money = field.NewFloat64(table, "money")
	e.MinPrice = field.NewFloat64(table, "min_price")
	e.Type = field.NewString(table, "type")
	e.Status = field.NewBool(table, "status")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")
	e.StartTime = field.NewTime(table, "start_time")
	e.EndTime = field.NewTime(table, "end_time")
	e.UseTime = field.NewTime(table, "use_time")
	e.UseType = field.NewBool(table, "use_type")
	e.PrimaryKey = field.NewString(table, "primary_key")

	e.fillFieldMap()

	return e
}

func (e *ebStoreCouponUser) WithContext(ctx context.Context) IEbStoreCouponUserDo {
	return e.ebStoreCouponUserDo.WithContext(ctx)
}

func (e ebStoreCouponUser) TableName() string { return e.ebStoreCouponUserDo.TableName() }

func (e ebStoreCouponUser) Alias() string { return e.ebStoreCouponUserDo.Alias() }

func (e ebStoreCouponUser) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreCouponUserDo.Columns(cols...)
}

func (e *ebStoreCouponUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreCouponUser) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 16)
	e.fieldMap["id"] = e.ID
	e.fieldMap["coupon_id"] = e.CouponID
	e.fieldMap["cid"] = e.Cid
	e.fieldMap["uid"] = e.UID
	e.fieldMap["name"] = e.Name
	e.fieldMap["money"] = e.Money
	e.fieldMap["min_price"] = e.MinPrice
	e.fieldMap["type"] = e.Type
	e.fieldMap["status"] = e.Status
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
	e.fieldMap["start_time"] = e.StartTime
	e.fieldMap["end_time"] = e.EndTime
	e.fieldMap["use_time"] = e.UseTime
	e.fieldMap["use_type"] = e.UseType
	e.fieldMap["primary_key"] = e.PrimaryKey
}

func (e ebStoreCouponUser) clone(db *gorm.DB) ebStoreCouponUser {
	e.ebStoreCouponUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreCouponUser) replaceDB(db *gorm.DB) ebStoreCouponUser {
	e.ebStoreCouponUserDo.ReplaceDB(db)
	return e
}

type ebStoreCouponUserDo struct{ gen.DO }

type IEbStoreCouponUserDo interface {
	gen.SubQuery
	Debug() IEbStoreCouponUserDo
	WithContext(ctx context.Context) IEbStoreCouponUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreCouponUserDo
	WriteDB() IEbStoreCouponUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreCouponUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreCouponUserDo
	Not(conds ...gen.Condition) IEbStoreCouponUserDo
	Or(conds ...gen.Condition) IEbStoreCouponUserDo
	Select(conds ...field.Expr) IEbStoreCouponUserDo
	Where(conds ...gen.Condition) IEbStoreCouponUserDo
	Order(conds ...field.Expr) IEbStoreCouponUserDo
	Distinct(cols ...field.Expr) IEbStoreCouponUserDo
	Omit(cols ...field.Expr) IEbStoreCouponUserDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreCouponUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreCouponUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreCouponUserDo
	Group(cols ...field.Expr) IEbStoreCouponUserDo
	Having(conds ...gen.Condition) IEbStoreCouponUserDo
	Limit(limit int) IEbStoreCouponUserDo
	Offset(offset int) IEbStoreCouponUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreCouponUserDo
	Unscoped() IEbStoreCouponUserDo
	Create(values ...*model.EbStoreCouponUser) error
	CreateInBatches(values []*model.EbStoreCouponUser, batchSize int) error
	Save(values ...*model.EbStoreCouponUser) error
	First() (*model.EbStoreCouponUser, error)
	Take() (*model.EbStoreCouponUser, error)
	Last() (*model.EbStoreCouponUser, error)
	Find() ([]*model.EbStoreCouponUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreCouponUser, err error)
	FindInBatches(result *[]*model.EbStoreCouponUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreCouponUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreCouponUserDo
	Assign(attrs ...field.AssignExpr) IEbStoreCouponUserDo
	Joins(fields ...field.RelationField) IEbStoreCouponUserDo
	Preload(fields ...field.RelationField) IEbStoreCouponUserDo
	FirstOrInit() (*model.EbStoreCouponUser, error)
	FirstOrCreate() (*model.EbStoreCouponUser, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreCouponUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreCouponUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreCouponUserDo) Debug() IEbStoreCouponUserDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreCouponUserDo) WithContext(ctx context.Context) IEbStoreCouponUserDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreCouponUserDo) ReadDB() IEbStoreCouponUserDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreCouponUserDo) WriteDB() IEbStoreCouponUserDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreCouponUserDo) Session(config *gorm.Session) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreCouponUserDo) Clauses(conds ...clause.Expression) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreCouponUserDo) Returning(value interface{}, columns ...string) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreCouponUserDo) Not(conds ...gen.Condition) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreCouponUserDo) Or(conds ...gen.Condition) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreCouponUserDo) Select(conds ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreCouponUserDo) Where(conds ...gen.Condition) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreCouponUserDo) Order(conds ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreCouponUserDo) Distinct(cols ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreCouponUserDo) Omit(cols ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreCouponUserDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreCouponUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreCouponUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreCouponUserDo) Group(cols ...field.Expr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreCouponUserDo) Having(conds ...gen.Condition) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreCouponUserDo) Limit(limit int) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreCouponUserDo) Offset(offset int) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreCouponUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreCouponUserDo) Unscoped() IEbStoreCouponUserDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreCouponUserDo) Create(values ...*model.EbStoreCouponUser) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreCouponUserDo) CreateInBatches(values []*model.EbStoreCouponUser, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreCouponUserDo) Save(values ...*model.EbStoreCouponUser) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreCouponUserDo) First() (*model.EbStoreCouponUser, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCouponUser), nil
	}
}

func (e ebStoreCouponUserDo) Take() (*model.EbStoreCouponUser, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCouponUser), nil
	}
}

func (e ebStoreCouponUserDo) Last() (*model.EbStoreCouponUser, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCouponUser), nil
	}
}

func (e ebStoreCouponUserDo) Find() ([]*model.EbStoreCouponUser, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreCouponUser), err
}

func (e ebStoreCouponUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreCouponUser, err error) {
	buf := make([]*model.EbStoreCouponUser, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreCouponUserDo) FindInBatches(result *[]*model.EbStoreCouponUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreCouponUserDo) Attrs(attrs ...field.AssignExpr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreCouponUserDo) Assign(attrs ...field.AssignExpr) IEbStoreCouponUserDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreCouponUserDo) Joins(fields ...field.RelationField) IEbStoreCouponUserDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreCouponUserDo) Preload(fields ...field.RelationField) IEbStoreCouponUserDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreCouponUserDo) FirstOrInit() (*model.EbStoreCouponUser, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCouponUser), nil
	}
}

func (e ebStoreCouponUserDo) FirstOrCreate() (*model.EbStoreCouponUser, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCouponUser), nil
	}
}

func (e ebStoreCouponUserDo) FindByPage(offset int, limit int) (result []*model.EbStoreCouponUser, count int64, err error) {
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

func (e ebStoreCouponUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreCouponUserDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreCouponUserDo) Delete(models ...*model.EbStoreCouponUser) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreCouponUserDo) withDO(do gen.Dao) *ebStoreCouponUserDo {
	e.DO = *do.(*gen.DO)
	return e
}