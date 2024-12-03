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

func newEbStoreProductCate(db *gorm.DB, opts ...gen.DOOption) ebStoreProductCate {
	_ebStoreProductCate := ebStoreProductCate{}

	_ebStoreProductCate.ebStoreProductCateDo.UseDB(db, opts...)
	_ebStoreProductCate.ebStoreProductCateDo.UseModel(&model.EbStoreProductCate{})

	tableName := _ebStoreProductCate.ebStoreProductCateDo.TableName()
	_ebStoreProductCate.ALL = field.NewAsterisk(tableName)
	_ebStoreProductCate.ID = field.NewInt32(tableName, "id")
	_ebStoreProductCate.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreProductCate.CateID = field.NewInt32(tableName, "cate_id")
	_ebStoreProductCate.AddTime = field.NewInt32(tableName, "add_time")

	_ebStoreProductCate.fillFieldMap()

	return _ebStoreProductCate
}

// ebStoreProductCate 商品分类辅助表
type ebStoreProductCate struct {
	ebStoreProductCateDo ebStoreProductCateDo

	ALL       field.Asterisk
	ID        field.Int32
	ProductID field.Int32 // 商品id
	CateID    field.Int32 // 分类id
	AddTime   field.Int32 // 添加时间

	fieldMap map[string]field.Expr
}

func (e ebStoreProductCate) Table(newTableName string) *ebStoreProductCate {
	e.ebStoreProductCateDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreProductCate) As(alias string) *ebStoreProductCate {
	e.ebStoreProductCateDo.DO = *(e.ebStoreProductCateDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreProductCate) updateTableName(table string) *ebStoreProductCate {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.ProductID = field.NewInt32(table, "product_id")
	e.CateID = field.NewInt32(table, "cate_id")
	e.AddTime = field.NewInt32(table, "add_time")

	e.fillFieldMap()

	return e
}

func (e *ebStoreProductCate) WithContext(ctx context.Context) IEbStoreProductCateDo {
	return e.ebStoreProductCateDo.WithContext(ctx)
}

func (e ebStoreProductCate) TableName() string { return e.ebStoreProductCateDo.TableName() }

func (e ebStoreProductCate) Alias() string { return e.ebStoreProductCateDo.Alias() }

func (e ebStoreProductCate) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreProductCateDo.Columns(cols...)
}

func (e *ebStoreProductCate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreProductCate) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 4)
	e.fieldMap["id"] = e.ID
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["cate_id"] = e.CateID
	e.fieldMap["add_time"] = e.AddTime
}

func (e ebStoreProductCate) clone(db *gorm.DB) ebStoreProductCate {
	e.ebStoreProductCateDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreProductCate) replaceDB(db *gorm.DB) ebStoreProductCate {
	e.ebStoreProductCateDo.ReplaceDB(db)
	return e
}

type ebStoreProductCateDo struct{ gen.DO }

type IEbStoreProductCateDo interface {
	gen.SubQuery
	Debug() IEbStoreProductCateDo
	WithContext(ctx context.Context) IEbStoreProductCateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreProductCateDo
	WriteDB() IEbStoreProductCateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreProductCateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreProductCateDo
	Not(conds ...gen.Condition) IEbStoreProductCateDo
	Or(conds ...gen.Condition) IEbStoreProductCateDo
	Select(conds ...field.Expr) IEbStoreProductCateDo
	Where(conds ...gen.Condition) IEbStoreProductCateDo
	Order(conds ...field.Expr) IEbStoreProductCateDo
	Distinct(cols ...field.Expr) IEbStoreProductCateDo
	Omit(cols ...field.Expr) IEbStoreProductCateDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreProductCateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCateDo
	Group(cols ...field.Expr) IEbStoreProductCateDo
	Having(conds ...gen.Condition) IEbStoreProductCateDo
	Limit(limit int) IEbStoreProductCateDo
	Offset(offset int) IEbStoreProductCateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductCateDo
	Unscoped() IEbStoreProductCateDo
	Create(values ...*model.EbStoreProductCate) error
	CreateInBatches(values []*model.EbStoreProductCate, batchSize int) error
	Save(values ...*model.EbStoreProductCate) error
	First() (*model.EbStoreProductCate, error)
	Take() (*model.EbStoreProductCate, error)
	Last() (*model.EbStoreProductCate, error)
	Find() ([]*model.EbStoreProductCate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductCate, err error)
	FindInBatches(result *[]*model.EbStoreProductCate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreProductCate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreProductCateDo
	Assign(attrs ...field.AssignExpr) IEbStoreProductCateDo
	Joins(fields ...field.RelationField) IEbStoreProductCateDo
	Preload(fields ...field.RelationField) IEbStoreProductCateDo
	FirstOrInit() (*model.EbStoreProductCate, error)
	FirstOrCreate() (*model.EbStoreProductCate, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreProductCate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreProductCateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreProductCateDo) Debug() IEbStoreProductCateDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreProductCateDo) WithContext(ctx context.Context) IEbStoreProductCateDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreProductCateDo) ReadDB() IEbStoreProductCateDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreProductCateDo) WriteDB() IEbStoreProductCateDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreProductCateDo) Session(config *gorm.Session) IEbStoreProductCateDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreProductCateDo) Clauses(conds ...clause.Expression) IEbStoreProductCateDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreProductCateDo) Returning(value interface{}, columns ...string) IEbStoreProductCateDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreProductCateDo) Not(conds ...gen.Condition) IEbStoreProductCateDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreProductCateDo) Or(conds ...gen.Condition) IEbStoreProductCateDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreProductCateDo) Select(conds ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreProductCateDo) Where(conds ...gen.Condition) IEbStoreProductCateDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreProductCateDo) Order(conds ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreProductCateDo) Distinct(cols ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreProductCateDo) Omit(cols ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreProductCateDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreProductCateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreProductCateDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreProductCateDo) Group(cols ...field.Expr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreProductCateDo) Having(conds ...gen.Condition) IEbStoreProductCateDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreProductCateDo) Limit(limit int) IEbStoreProductCateDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreProductCateDo) Offset(offset int) IEbStoreProductCateDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreProductCateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductCateDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreProductCateDo) Unscoped() IEbStoreProductCateDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreProductCateDo) Create(values ...*model.EbStoreProductCate) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreProductCateDo) CreateInBatches(values []*model.EbStoreProductCate, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreProductCateDo) Save(values ...*model.EbStoreProductCate) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreProductCateDo) First() (*model.EbStoreProductCate, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCate), nil
	}
}

func (e ebStoreProductCateDo) Take() (*model.EbStoreProductCate, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCate), nil
	}
}

func (e ebStoreProductCateDo) Last() (*model.EbStoreProductCate, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCate), nil
	}
}

func (e ebStoreProductCateDo) Find() ([]*model.EbStoreProductCate, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreProductCate), err
}

func (e ebStoreProductCateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductCate, err error) {
	buf := make([]*model.EbStoreProductCate, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreProductCateDo) FindInBatches(result *[]*model.EbStoreProductCate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreProductCateDo) Attrs(attrs ...field.AssignExpr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreProductCateDo) Assign(attrs ...field.AssignExpr) IEbStoreProductCateDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreProductCateDo) Joins(fields ...field.RelationField) IEbStoreProductCateDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreProductCateDo) Preload(fields ...field.RelationField) IEbStoreProductCateDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreProductCateDo) FirstOrInit() (*model.EbStoreProductCate, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCate), nil
	}
}

func (e ebStoreProductCateDo) FirstOrCreate() (*model.EbStoreProductCate, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCate), nil
	}
}

func (e ebStoreProductCateDo) FindByPage(offset int, limit int) (result []*model.EbStoreProductCate, count int64, err error) {
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

func (e ebStoreProductCateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreProductCateDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreProductCateDo) Delete(models ...*model.EbStoreProductCate) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreProductCateDo) withDO(do gen.Dao) *ebStoreProductCateDo {
	e.DO = *do.(*gen.DO)
	return e
}
