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

func newEbStoreProductAttrResult(db *gorm.DB, opts ...gen.DOOption) ebStoreProductAttrResult {
	_ebStoreProductAttrResult := ebStoreProductAttrResult{}

	_ebStoreProductAttrResult.ebStoreProductAttrResultDo.UseDB(db, opts...)
	_ebStoreProductAttrResult.ebStoreProductAttrResultDo.UseModel(&model.EbStoreProductAttrResult{})

	tableName := _ebStoreProductAttrResult.ebStoreProductAttrResultDo.TableName()
	_ebStoreProductAttrResult.ALL = field.NewAsterisk(tableName)
	_ebStoreProductAttrResult.ID = field.NewInt32(tableName, "id")
	_ebStoreProductAttrResult.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreProductAttrResult.Result = field.NewString(tableName, "result")
	_ebStoreProductAttrResult.ChangeTime = field.NewInt32(tableName, "change_time")
	_ebStoreProductAttrResult.Type = field.NewBool(tableName, "type")

	_ebStoreProductAttrResult.fillFieldMap()

	return _ebStoreProductAttrResult
}

// ebStoreProductAttrResult 商品属性详情表
type ebStoreProductAttrResult struct {
	ebStoreProductAttrResultDo ebStoreProductAttrResultDo

	ALL        field.Asterisk
	ID         field.Int32  // 主键
	ProductID  field.Int32  // 商品ID
	Result     field.String // 商品属性参数
	ChangeTime field.Int32  // 上次修改时间
	Type       field.Bool   // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团

	fieldMap map[string]field.Expr
}

func (e ebStoreProductAttrResult) Table(newTableName string) *ebStoreProductAttrResult {
	e.ebStoreProductAttrResultDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreProductAttrResult) As(alias string) *ebStoreProductAttrResult {
	e.ebStoreProductAttrResultDo.DO = *(e.ebStoreProductAttrResultDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreProductAttrResult) updateTableName(table string) *ebStoreProductAttrResult {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.ProductID = field.NewInt32(table, "product_id")
	e.Result = field.NewString(table, "result")
	e.ChangeTime = field.NewInt32(table, "change_time")
	e.Type = field.NewBool(table, "type")

	e.fillFieldMap()

	return e
}

func (e *ebStoreProductAttrResult) WithContext(ctx context.Context) IEbStoreProductAttrResultDo {
	return e.ebStoreProductAttrResultDo.WithContext(ctx)
}

func (e ebStoreProductAttrResult) TableName() string { return e.ebStoreProductAttrResultDo.TableName() }

func (e ebStoreProductAttrResult) Alias() string { return e.ebStoreProductAttrResultDo.Alias() }

func (e ebStoreProductAttrResult) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreProductAttrResultDo.Columns(cols...)
}

func (e *ebStoreProductAttrResult) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreProductAttrResult) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 5)
	e.fieldMap["id"] = e.ID
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["result"] = e.Result
	e.fieldMap["change_time"] = e.ChangeTime
	e.fieldMap["type"] = e.Type
}

func (e ebStoreProductAttrResult) clone(db *gorm.DB) ebStoreProductAttrResult {
	e.ebStoreProductAttrResultDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreProductAttrResult) replaceDB(db *gorm.DB) ebStoreProductAttrResult {
	e.ebStoreProductAttrResultDo.ReplaceDB(db)
	return e
}

type ebStoreProductAttrResultDo struct{ gen.DO }

type IEbStoreProductAttrResultDo interface {
	gen.SubQuery
	Debug() IEbStoreProductAttrResultDo
	WithContext(ctx context.Context) IEbStoreProductAttrResultDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreProductAttrResultDo
	WriteDB() IEbStoreProductAttrResultDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreProductAttrResultDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreProductAttrResultDo
	Not(conds ...gen.Condition) IEbStoreProductAttrResultDo
	Or(conds ...gen.Condition) IEbStoreProductAttrResultDo
	Select(conds ...field.Expr) IEbStoreProductAttrResultDo
	Where(conds ...gen.Condition) IEbStoreProductAttrResultDo
	Order(conds ...field.Expr) IEbStoreProductAttrResultDo
	Distinct(cols ...field.Expr) IEbStoreProductAttrResultDo
	Omit(cols ...field.Expr) IEbStoreProductAttrResultDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreProductAttrResultDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductAttrResultDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductAttrResultDo
	Group(cols ...field.Expr) IEbStoreProductAttrResultDo
	Having(conds ...gen.Condition) IEbStoreProductAttrResultDo
	Limit(limit int) IEbStoreProductAttrResultDo
	Offset(offset int) IEbStoreProductAttrResultDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductAttrResultDo
	Unscoped() IEbStoreProductAttrResultDo
	Create(values ...*model.EbStoreProductAttrResult) error
	CreateInBatches(values []*model.EbStoreProductAttrResult, batchSize int) error
	Save(values ...*model.EbStoreProductAttrResult) error
	First() (*model.EbStoreProductAttrResult, error)
	Take() (*model.EbStoreProductAttrResult, error)
	Last() (*model.EbStoreProductAttrResult, error)
	Find() ([]*model.EbStoreProductAttrResult, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductAttrResult, err error)
	FindInBatches(result *[]*model.EbStoreProductAttrResult, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreProductAttrResult) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreProductAttrResultDo
	Assign(attrs ...field.AssignExpr) IEbStoreProductAttrResultDo
	Joins(fields ...field.RelationField) IEbStoreProductAttrResultDo
	Preload(fields ...field.RelationField) IEbStoreProductAttrResultDo
	FirstOrInit() (*model.EbStoreProductAttrResult, error)
	FirstOrCreate() (*model.EbStoreProductAttrResult, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreProductAttrResult, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreProductAttrResultDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreProductAttrResultDo) Debug() IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreProductAttrResultDo) WithContext(ctx context.Context) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreProductAttrResultDo) ReadDB() IEbStoreProductAttrResultDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreProductAttrResultDo) WriteDB() IEbStoreProductAttrResultDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreProductAttrResultDo) Session(config *gorm.Session) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreProductAttrResultDo) Clauses(conds ...clause.Expression) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreProductAttrResultDo) Returning(value interface{}, columns ...string) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreProductAttrResultDo) Not(conds ...gen.Condition) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreProductAttrResultDo) Or(conds ...gen.Condition) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreProductAttrResultDo) Select(conds ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreProductAttrResultDo) Where(conds ...gen.Condition) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreProductAttrResultDo) Order(conds ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreProductAttrResultDo) Distinct(cols ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreProductAttrResultDo) Omit(cols ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreProductAttrResultDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreProductAttrResultDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreProductAttrResultDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreProductAttrResultDo) Group(cols ...field.Expr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreProductAttrResultDo) Having(conds ...gen.Condition) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreProductAttrResultDo) Limit(limit int) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreProductAttrResultDo) Offset(offset int) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreProductAttrResultDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreProductAttrResultDo) Unscoped() IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreProductAttrResultDo) Create(values ...*model.EbStoreProductAttrResult) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreProductAttrResultDo) CreateInBatches(values []*model.EbStoreProductAttrResult, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreProductAttrResultDo) Save(values ...*model.EbStoreProductAttrResult) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreProductAttrResultDo) First() (*model.EbStoreProductAttrResult, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductAttrResult), nil
	}
}

func (e ebStoreProductAttrResultDo) Take() (*model.EbStoreProductAttrResult, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductAttrResult), nil
	}
}

func (e ebStoreProductAttrResultDo) Last() (*model.EbStoreProductAttrResult, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductAttrResult), nil
	}
}

func (e ebStoreProductAttrResultDo) Find() ([]*model.EbStoreProductAttrResult, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreProductAttrResult), err
}

func (e ebStoreProductAttrResultDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductAttrResult, err error) {
	buf := make([]*model.EbStoreProductAttrResult, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreProductAttrResultDo) FindInBatches(result *[]*model.EbStoreProductAttrResult, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreProductAttrResultDo) Attrs(attrs ...field.AssignExpr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreProductAttrResultDo) Assign(attrs ...field.AssignExpr) IEbStoreProductAttrResultDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreProductAttrResultDo) Joins(fields ...field.RelationField) IEbStoreProductAttrResultDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreProductAttrResultDo) Preload(fields ...field.RelationField) IEbStoreProductAttrResultDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreProductAttrResultDo) FirstOrInit() (*model.EbStoreProductAttrResult, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductAttrResult), nil
	}
}

func (e ebStoreProductAttrResultDo) FirstOrCreate() (*model.EbStoreProductAttrResult, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductAttrResult), nil
	}
}

func (e ebStoreProductAttrResultDo) FindByPage(offset int, limit int) (result []*model.EbStoreProductAttrResult, count int64, err error) {
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

func (e ebStoreProductAttrResultDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreProductAttrResultDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreProductAttrResultDo) Delete(models ...*model.EbStoreProductAttrResult) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreProductAttrResultDo) withDO(do gen.Dao) *ebStoreProductAttrResultDo {
	e.DO = *do.(*gen.DO)
	return e
}
