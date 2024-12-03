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

func newEbUserSign(db *gorm.DB, opts ...gen.DOOption) ebUserSign {
	_ebUserSign := ebUserSign{}

	_ebUserSign.ebUserSignDo.UseDB(db, opts...)
	_ebUserSign.ebUserSignDo.UseModel(&model.EbUserSign{})

	tableName := _ebUserSign.ebUserSignDo.TableName()
	_ebUserSign.ALL = field.NewAsterisk(tableName)
	_ebUserSign.ID = field.NewInt32(tableName, "id")
	_ebUserSign.UID = field.NewInt32(tableName, "uid")
	_ebUserSign.Title = field.NewString(tableName, "title")
	_ebUserSign.Number = field.NewInt32(tableName, "number")
	_ebUserSign.Balance = field.NewInt32(tableName, "balance")
	_ebUserSign.Type = field.NewBool(tableName, "type")
	_ebUserSign.CreateDay = field.NewTime(tableName, "create_day")
	_ebUserSign.CreateTime = field.NewTime(tableName, "create_time")

	_ebUserSign.fillFieldMap()

	return _ebUserSign
}

// ebUserSign 签到记录表
type ebUserSign struct {
	ebUserSignDo ebUserSignDo

	ALL        field.Asterisk
	ID         field.Int32
	UID        field.Int32  // 用户uid
	Title      field.String // 签到说明
	Number     field.Int32  // 获得
	Balance    field.Int32  // 剩余
	Type       field.Bool   // 类型，1积分，2经验
	CreateDay  field.Time   // 签到日期
	CreateTime field.Time   // 添加时间

	fieldMap map[string]field.Expr
}

func (e ebUserSign) Table(newTableName string) *ebUserSign {
	e.ebUserSignDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebUserSign) As(alias string) *ebUserSign {
	e.ebUserSignDo.DO = *(e.ebUserSignDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebUserSign) updateTableName(table string) *ebUserSign {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.UID = field.NewInt32(table, "uid")
	e.Title = field.NewString(table, "title")
	e.Number = field.NewInt32(table, "number")
	e.Balance = field.NewInt32(table, "balance")
	e.Type = field.NewBool(table, "type")
	e.CreateDay = field.NewTime(table, "create_day")
	e.CreateTime = field.NewTime(table, "create_time")

	e.fillFieldMap()

	return e
}

func (e *ebUserSign) WithContext(ctx context.Context) IEbUserSignDo {
	return e.ebUserSignDo.WithContext(ctx)
}

func (e ebUserSign) TableName() string { return e.ebUserSignDo.TableName() }

func (e ebUserSign) Alias() string { return e.ebUserSignDo.Alias() }

func (e ebUserSign) Columns(cols ...field.Expr) gen.Columns { return e.ebUserSignDo.Columns(cols...) }

func (e *ebUserSign) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebUserSign) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 8)
	e.fieldMap["id"] = e.ID
	e.fieldMap["uid"] = e.UID
	e.fieldMap["title"] = e.Title
	e.fieldMap["number"] = e.Number
	e.fieldMap["balance"] = e.Balance
	e.fieldMap["type"] = e.Type
	e.fieldMap["create_day"] = e.CreateDay
	e.fieldMap["create_time"] = e.CreateTime
}

func (e ebUserSign) clone(db *gorm.DB) ebUserSign {
	e.ebUserSignDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebUserSign) replaceDB(db *gorm.DB) ebUserSign {
	e.ebUserSignDo.ReplaceDB(db)
	return e
}

type ebUserSignDo struct{ gen.DO }

type IEbUserSignDo interface {
	gen.SubQuery
	Debug() IEbUserSignDo
	WithContext(ctx context.Context) IEbUserSignDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbUserSignDo
	WriteDB() IEbUserSignDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbUserSignDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbUserSignDo
	Not(conds ...gen.Condition) IEbUserSignDo
	Or(conds ...gen.Condition) IEbUserSignDo
	Select(conds ...field.Expr) IEbUserSignDo
	Where(conds ...gen.Condition) IEbUserSignDo
	Order(conds ...field.Expr) IEbUserSignDo
	Distinct(cols ...field.Expr) IEbUserSignDo
	Omit(cols ...field.Expr) IEbUserSignDo
	Join(table schema.Tabler, on ...field.Expr) IEbUserSignDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserSignDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbUserSignDo
	Group(cols ...field.Expr) IEbUserSignDo
	Having(conds ...gen.Condition) IEbUserSignDo
	Limit(limit int) IEbUserSignDo
	Offset(offset int) IEbUserSignDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserSignDo
	Unscoped() IEbUserSignDo
	Create(values ...*model.EbUserSign) error
	CreateInBatches(values []*model.EbUserSign, batchSize int) error
	Save(values ...*model.EbUserSign) error
	First() (*model.EbUserSign, error)
	Take() (*model.EbUserSign, error)
	Last() (*model.EbUserSign, error)
	Find() ([]*model.EbUserSign, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserSign, err error)
	FindInBatches(result *[]*model.EbUserSign, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbUserSign) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbUserSignDo
	Assign(attrs ...field.AssignExpr) IEbUserSignDo
	Joins(fields ...field.RelationField) IEbUserSignDo
	Preload(fields ...field.RelationField) IEbUserSignDo
	FirstOrInit() (*model.EbUserSign, error)
	FirstOrCreate() (*model.EbUserSign, error)
	FindByPage(offset int, limit int) (result []*model.EbUserSign, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbUserSignDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebUserSignDo) Debug() IEbUserSignDo {
	return e.withDO(e.DO.Debug())
}

func (e ebUserSignDo) WithContext(ctx context.Context) IEbUserSignDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebUserSignDo) ReadDB() IEbUserSignDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebUserSignDo) WriteDB() IEbUserSignDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebUserSignDo) Session(config *gorm.Session) IEbUserSignDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebUserSignDo) Clauses(conds ...clause.Expression) IEbUserSignDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebUserSignDo) Returning(value interface{}, columns ...string) IEbUserSignDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebUserSignDo) Not(conds ...gen.Condition) IEbUserSignDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebUserSignDo) Or(conds ...gen.Condition) IEbUserSignDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebUserSignDo) Select(conds ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebUserSignDo) Where(conds ...gen.Condition) IEbUserSignDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebUserSignDo) Order(conds ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebUserSignDo) Distinct(cols ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebUserSignDo) Omit(cols ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebUserSignDo) Join(table schema.Tabler, on ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebUserSignDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebUserSignDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebUserSignDo) Group(cols ...field.Expr) IEbUserSignDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebUserSignDo) Having(conds ...gen.Condition) IEbUserSignDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebUserSignDo) Limit(limit int) IEbUserSignDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebUserSignDo) Offset(offset int) IEbUserSignDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebUserSignDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserSignDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebUserSignDo) Unscoped() IEbUserSignDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebUserSignDo) Create(values ...*model.EbUserSign) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebUserSignDo) CreateInBatches(values []*model.EbUserSign, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebUserSignDo) Save(values ...*model.EbUserSign) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebUserSignDo) First() (*model.EbUserSign, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserSign), nil
	}
}

func (e ebUserSignDo) Take() (*model.EbUserSign, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserSign), nil
	}
}

func (e ebUserSignDo) Last() (*model.EbUserSign, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserSign), nil
	}
}

func (e ebUserSignDo) Find() ([]*model.EbUserSign, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbUserSign), err
}

func (e ebUserSignDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserSign, err error) {
	buf := make([]*model.EbUserSign, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebUserSignDo) FindInBatches(result *[]*model.EbUserSign, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebUserSignDo) Attrs(attrs ...field.AssignExpr) IEbUserSignDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebUserSignDo) Assign(attrs ...field.AssignExpr) IEbUserSignDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebUserSignDo) Joins(fields ...field.RelationField) IEbUserSignDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebUserSignDo) Preload(fields ...field.RelationField) IEbUserSignDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebUserSignDo) FirstOrInit() (*model.EbUserSign, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserSign), nil
	}
}

func (e ebUserSignDo) FirstOrCreate() (*model.EbUserSign, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserSign), nil
	}
}

func (e ebUserSignDo) FindByPage(offset int, limit int) (result []*model.EbUserSign, count int64, err error) {
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

func (e ebUserSignDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebUserSignDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebUserSignDo) Delete(models ...*model.EbUserSign) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebUserSignDo) withDO(do gen.Dao) *ebUserSignDo {
	e.DO = *do.(*gen.DO)
	return e
}